package paymentchannel

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	_ "perun.network/go-perun/backend/ethereum" // backend init
	echannel "perun.network/go-perun/backend/ethereum/channel"
	ewallet "perun.network/go-perun/backend/ethereum/wallet"
	pkeystore "perun.network/go-perun/backend/ethereum/wallet/keystore"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	"perun.network/go-perun/wallet"
	"perun.network/go-perun/wire"

	plog "perun.network/go-perun/log"
	"perun.network/go-perun/wire/net"
	"perun.network/go-perun/wire/net/simple"

	"restapidemo/model"
)

type peer struct {
	alias  string
	peerID wire.Address
	ch     *paymentChannel
	plog   plog.Logger
}

//Node is the structure to hold all necessary data for a current node (user session)
type node struct {
	alias string
	plog  plog.Logger

	client *client.Client
	dialer *simple.Dialer
	bus    *net.Bus

	// Account for signing on-chain TX. Currently also the Perun-ID.
	// changed into Address
	onChain *pkeystore.Account
	// Account for signing off-chain TX. Currently one Account for all
	// state channels, later one we want one Account per Channel.
	offChain wallet.Account
	wallet   *pkeystore.Wallet

	adjudicator channel.Adjudicator
	adjAddr     common.Address
	asset       channel.Asset
	assetAddr   common.Address
	funder      channel.Funder
	// Needed to deploy contracts.
	cb echannel.ContractBackend

	// Protects peers
	mtx   sync.Mutex
	peers map[string]*peer
}

//NodePeers contain peers for an node
type NodePeers struct {
	Peers map[string]*peer
}

//GlobalNodePeers to know all connected peers
var GlobalNodePeers map[string]*NodePeers

func (sess *Session) getOnChainBal(ctx context.Context, addrs ...wallet.Address) ([]*big.Int, error) {
	bals := make([]*big.Int, len(addrs))
	var err error
	for idx, addr := range addrs {
		bals[idx], err = sess.EthereumBackend.BalanceAt(ctx, ewallet.AsEthAddr(addr), nil)
		if err != nil {
			return nil, errors.Wrap(err, "querying on-chain balance")
		}
	}
	return bals, nil
}

//Connect a peer
func (n *node) Connect(args []string, otherSess *Session, conf *Config) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	alias := args[0]
	log.Info().Msgf("Connecting ... %s", alias)

	if n.peers[alias] != nil {
		log.Info().Msgf("Peer already connected : %s", alias)
		return errors.New("Peer already connected")
	}

	peerCfg, ok := tsconfig.config.Peers[alias]
	if !ok {
		log.Info().Msgf("alias unknown: %s", alias)
		return errors.Errorf("Alias '%s' unknown. Add it to 'network.yaml'.", alias)
	}

	//make addr from string
	peerCfg.peerID = otherSess.Backend.onChain.Address().String() //n.onChain.Address().String()
	peerAddr, err := StrToWireAddress(peerCfg.peerID)
	conf.Peers[alias].peerID = peerCfg.peerID

	if err != nil {
		log.Info().Msgf("could not get peer address [%s] : %s", peerCfg.peerID, err)
		return errors.WithMessage(err, "could not get to peer address")
	}

	peerHost := peerCfg.Hostname + ":" + strconv.Itoa(int(peerCfg.Port))
	n.dialer.Register(peerAddr, peerHost)

	n.peers[alias] = &peer{
		alias:  alias,
		peerID: peerAddr,
		plog:   plog.WithField("peer", peerCfg.peerID),
	}

	log.Info().Msg("Connected.")
	return nil
}

// peer returns the peer with the address `addr` or nil if not found.
func (n *node) peer(addr wire.Address) *peer {
	for _, peer := range n.peers {
		if addr.String() == peer.peerID.String() {
			return peer
		}
	}
	return nil
}

func (n *node) channelPeer(ch *client.Channel) *peer {
	peerID := ch.Peers()[1-ch.Idx()] // assumes two-party channel
	return n.peer(peerID)
}

func (n *node) setupChannel(ch *client.Channel) {
	channel.RegisterApp(&odsApp)
	app, err := channel.Resolve(appDefAddr)
	if err != nil {
		log.Info().Msgf("error setting up, app: %s", err)
	}
	ch.Params().App = app

	if len(ch.Peers()) != 2 {
		log.Fatal().Msg("Only channels with two participants are currently supported")
	}

	log.Info().Msgf("channel peers %d", 1-ch.Idx())
	log.Info().Msgf("channel owned %d", ch.Idx())

	peerID := ch.Peers()[1-ch.Idx()] // assumes two-party channel

	p := n.peer(peerID)

	if p == nil {
		log.Info().Msgf("Opened channel to unknown peer %s", peerID)
		return
	} else if p.ch != nil {
		log.Info().Msg("Peer tried to open more than one channel")
		return
	}

	p.ch = newPaymentChannel(ch)

	if p.ch.HasApp() {
		log.Info().Msg("Setup Channel Has App")
	} else {
		log.Info().Msg("SetupChannel No App found")
	}

	// Start watching.
	go func() {
		l := plog.WithField("channel", ch.ID())
		//log.Info().Msgf("Watcher started %d \n", ch.ID())
		err := ch.Watch(n)
		l.WithError(err).Debug("Watcher cancelled")
		if err != nil {
			fmt.Printf("Watcher stopped: %s \n", err)
		}
	}()

	bals := weiToEther(ch.State().Balances[0]...)
	log.Info().Msgf("\n🆕 OnNewChannel with %s. Initial balance: [My: %v Ξ, Peer: %v Ξ] idx:%d\n",
		p.alias, bals[ch.Idx()], bals[1-ch.Idx()], ch.Idx()) // assumes two-party channel
}

func (n *node) HandleAdjudicatorEvent(e channel.AdjudicatorEvent) {
	PrintfAsync("Received event\n")
	if _, ok := e.(*channel.ConcludedEvent); ok {
		PrintfAsync("🎭 Received concluded event\n")
		func() {
			n.mtx.Lock()
			defer n.mtx.Unlock()
			ch := n.channel(e.ID())
			if ch == nil {
				// If we initiated the channel closing, then the channel should
				// already be removed and we return.
				return
			}
			peer := n.channelPeer(ch.Channel)
			if err := n.settle(peer); err != nil {
				PrintfAsync("🎭 error while settling: %v\n", err)
			}
			PrintfAsync("🏁 Settled channel with %s.\n", peer.alias)
		}()
	}
}

// HandleFinal is called when the channel with peer `p` received a final update,
// indicating closure.
func (n *node) handleFinal(p *peer) {
	// Without on-chain watchers we just wait one second before try to settle.
	// Otherwise out settling could collide with the other party's.
	// Needs to be increased for geth-Nodes but works for ganache.
	time.Sleep(time.Second)
	n.mtx.Lock()
	defer n.mtx.Unlock()
	n.settle(p)
}

type balTuple struct {
	My, Other *big.Int
}

//GetBals of a node and it connected peers
func (n *node) GetBals() map[string]balTuple {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	bals := make(map[string]balTuple)
	for alias, peer := range n.peers {
		if peer.ch != nil {
			my, other := peer.ch.GetBalances()
			bals[alias] = balTuple{my, other}
		}
	}
	return bals
}

func findConfig(id wire.Address) (string, *netConfigEntry) {
	log.Info().Msgf("try to find id : %s", id.String())

	for alias, e := range tsconfig.config.Peers {
		log.Info().Msgf("validate : %s", alias)
		log.Info().Msgf("validate : %s", e.peerID)
		log.Info().Msgf("validate : %s", e.peerAddress.String())
		if e.peerID == id.String() {
			return alias, e
		}
	}
	return "", nil
}

//HandleUpdate does
func (n *node) HandleUpdate(update client.ChannelUpdate, resp *client.UpdateResponder) {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	plog := plog.WithField("channel", update.State.ID)
	log.Info().Msgf("Channel %s updated", update.State.ID)

	ch := n.channel(update.State.ID)
	if ch == nil {
		log.Info().Msg("Channel for ID not found")
		plog.Println("Channel for ID not found")
		return
	}
	ch.Handle(update, resp)
}

func (n *node) channel(id channel.ID) *paymentChannel {
	for _, p := range n.peers {
		if p.ch != nil && p.ch.ID() == id {
			return p.ch
		}
	}
	return nil
}

//HandleProposal does
func (n *node) HandleProposal(prop client.ChannelProposal, res *client.ProposalResponder) {
	req, ok := prop.(*client.LedgerChannelProposal)
	if !ok {
		log.Fatal().Msg("Can handle only ledger channel proposals.")
	}

	if len(req.Peers) != 2 {
		log.Fatal().Msg("Only channels with two participants are currently supported")
	}

	n.mtx.Lock()
	defer n.mtx.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), tsconfig.config.Node.HandleTimeout)
	defer cancel()
	id := req.Peers[0]
	log.Info().Msg("Received channel proposal")

	// Find the peer by its perunID and create it if not present
	log.Info().Msgf("try to find alias for peer_id %s", id.String())
	p := n.peer(id)
	alias, cfg := findConfig(id)
	log.Info().Msgf("found alias %s", alias)

	if p == nil {
		if cfg == nil {
			res.Reject(ctx, "Unknown identity")
			return
		}
		p = &peer{
			alias:  alias,
			peerID: id,
			plog:   plog.WithField("peer", id),
		}
		n.peers[alias] = p
		log.Info().Msgf("new peer : channel %s alias %s", id, alias)
		plog.WithField("channel", id).WithField("alias", alias).Debug("New peer")
	} else {

	}

	log.Info().Msg("channel proposal")
	plog.WithField("peer", id).Debug("Channel proposal")
	a := req.Accept(n.offChain.Address(), client.WithRandomNonce())
	if _, err := res.Accept(ctx, a); err != nil {
		log.Error().Msgf("could not accept channel proposal %s", err)
		plog.Error(errors.WithMessage(err, "accepting channel proposal"))
		return
	}
}

//Open opens a paymentchannel TODO: modify args
func (n *node) Open(mySess *Session, nc model.OpenPaymentChannel) error {
	balOwn := strconv.FormatUint(nc.OwnBalance, 10)
	balTheirs := strconv.FormatUint(nc.TheirsBalance, 10)

	n.mtx.Lock()
	defer n.mtx.Unlock()
	peer := n.peers[nc.Target]

	if peer == nil {
		log.Info().Msgf("peer not found %s", nc.Target)
		return errors.Errorf("peer not found %s", nc.Target)
	}

	log.Info().Msgf("try connect to %s  with own %s and theiers %s balance\n", nc.Target, balOwn, balTheirs)

	myBalEth, _ := new(big.Float).SetString(balOwn)
	peerBalEth, _ := new(big.Float).SetString(balTheirs)

	initBals := &channel.Allocation{
		Assets:   []channel.Asset{n.asset},
		Balances: [][]*big.Int{etherToWei(myBalEth, peerBalEth)},
	}
	//make addr from string
	peerAddr := peer.peerID
	onChainAddr, err := StrToWireAddress(n.onChain.Address().String())

	if err != nil {
		log.Info().Msgf("could not get to peer address : %s \n", err)
		return errors.WithMessage(err, "could not get to peer address")
	}

	propOpts := client.WithRandomNonce()

	appDefs := client.ProposalOpts{}
	appDefs["App"] = &odsApp

	prop, err := client.NewLedgerChannelProposal(
		tsconfig.config.Channel.ChallengeDurationSec,
		n.offChain.Address(),
		initBals,
		[]wire.Address{onChainAddr, peerAddr},
		propOpts,
		appDefs,
	)

	if err != nil {
		return errors.WithMessage(err, "creating channel proposal")
	}

	ctx, cancel := context.WithTimeout(context.Background(), tsconfig.config.Channel.FundTimeout)
	defer cancel()
	//log.Println("=== Proposal ====")
	//spew.Dump(prop)
	//log.Println("Proposing channel part 2")
	ch, err := n.client.ProposeChannel(ctx, prop)
	if err != nil {
		log.Info().Msgf("proposing channel failed : %s \n", err)
		return errors.WithMessage(err, "proposing channel failed")
	}

	if n.channel(ch.ID()) == nil {
		log.Info().Msg("OnNewChannel handler could not setup channel")
		return errors.New("OnNewChannel handler could not setup channel")
	}
	log.Info().Msgf("Proposed channel ... %d", ch.ID())
	return nil
}

//Send does submit money and data through the existing payment channel
func (n *node) Send(source string, upd model.UsePaymentChannel) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	plog.Traceln("Sending...")
	log.Info().Msg("Sending...")

	peer := n.peers[upd.Target]
	if peer == nil {
		log.Info().Msgf("peer not found %s\n", upd.Target)
		return errors.Errorf("peer not found %s", upd.Target)
	} else if peer.ch == nil {
		log.Info().Msg("connect to peer first")
		return errors.Errorf("connect to peer first")
	}

	amountEth, _ := new(big.Float).SetString(strconv.FormatUint(upd.Balance, 10))
	log.Info().Msgf("channel ", peer.ch.ID())
	pd := PaymentData{
		referenceToVersion: upd.ReferenceToVersion,
		balanceTransfered:  upd.Balance,
		message:            upd.Message,
		senderAlias:        source,
		receiverAlias:      upd.Target,
		timeSent:           time.Now(),
		timeReceived:       time.Now(),
		part:               upd.Part,
	}

	err := peer.ch.sendMoneyAndData(etherToWei(amountEth)[0], &pd)
	//log.Info().Str("payment-data", spew.Sdump(pd)).Msg("sent")

	if err != nil {
		log.Info().Msg("no money send")
	}
	return err
}

//Close does
//TODO: modify args
func (n *node) Close(alias string) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	log.Info().Msg("Closing ...")
	peer := n.peers[alias]
	if peer == nil {
		return errors.Errorf("Unknown peer: %s", alias)
	}

	if peer.ch == nil {
		log.Info().Msg("payment channel is nil ...")
		return nil
	}

	if err := peer.ch.sendFinal(); err != nil {
		return errors.WithMessage(err, "sending final state for state closing")
	}

	err := n.settle(peer)

	//removing references from peers
	for tmpAlias := range GlobalNodePeers {
		if alias == tmpAlias {
			GlobalNodePeers[alias] = nil
		} else {
			if GlobalNodePeers[alias] != nil {
				for tmpAl := range GlobalNodePeers[alias].Peers {
					if tmpAl == alias {
						GlobalNodePeers[alias] = nil
					}
				}
			}
		}
	}
	return err
}

func (n *node) settle(p *peer) error {
	if p.ch == nil {
		log.Info().Msgf("error closing channel : %s \n", "channel is nil")
	}

	p.ch.log.Debug("Settling")
	log.Info().Msg("Settling started")
	ctx, cancel := context.WithTimeout(context.Background(), tsconfig.config.Channel.SettleTimeout)
	defer cancel()

	// dirty hack, preventing reverting transaction with wrong Params
	p.ch.Params().App = channel.NoApp()
	finalBals := weiToEther(p.ch.GetBalances())

	if err := p.ch.Register(ctx); err != nil {
		return errors.WithMessage(err, "registering")
	}

	if err := p.ch.Settle(ctx, false); err != nil {
		log.Info().Msgf("error settling channel : %s \n", err)
		return errors.WithMessage(err, "settling the channel")
	}

	if err := p.ch.Close(); err != nil {
		log.Info().Msgf("error closing channel : %s \n", err)
		return errors.WithMessage(err, "channel closing")
	}

	log.Info().Msg("Removing channel")
	p.ch = nil

	log.Info().Msgf("\n🏁 Settled channel with %s. Final Balance: [My: %v Ξ, Peer: %v Ξ]\n", p.alias, finalBals[0], finalBals[1])

	delete(n.peers, p.alias)

	log.Info().Msgf("Removed %s from peer list.", p.alias)
	return nil
}

//Info prints the current status of all channels in the current session.
func (sess *Session) Info(sessionAlias string) (ir model.InfoResult, err error) {
	sess.Backend.mtx.Lock()
	defer sess.Backend.mtx.Unlock()
	plog.Traceln("Info...")
	log.Info().Msgf("Info requested for %s", sessionAlias)

	ctx, cancel := context.WithTimeout(context.Background(), tsconfig.config.Chain.TxTimeout)
	defer cancel()
	log.Info().Msgf("Peer\tPhase\tVersion\tMy Ξ\tPeer Ξ\tMy On-Chain Ξ\tPeer On-Chain Ξ\t")

	var infos []model.ChannelInfo
	ir.ChannelInfos = infos

	for alias, peer := range sess.Backend.peers {
		//make addr from string
		peerAddr := peer.peerID

		onChainBals, err := sess.getOnChainBal(ctx, sess.Backend.onChain.Address(), peerAddr)
		if err != nil {
			return ir, err
		}

		onChainBalsEth := weiToEther(onChainBals...)

		if peer.ch == nil {
			log.Info().Msgf("%s\t%s\t \t \t \t%v\t%v\t\n", alias, "Connected", onChainBalsEth[0], onChainBalsEth[1])
			ci := model.ChannelInfo{Peer: alias, Phase: "Connected",
				MyBalanceOnChain: onChainBalsEth[0], PeersBalanceOnChain: onChainBalsEth[1]}
			ir.ChannelInfos = append(ir.ChannelInfos, ci)
		} else {
			bals := weiToEther(peer.ch.GetBalances())
			log.Info().Msgf("%s\t%v\t%d\t%v\t%v\t%v\t%v\t\n",
				alias, peer.ch.Phase(), peer.ch.State().Version, bals[0], bals[1], onChainBalsEth[0], onChainBalsEth[1])

			ci := model.ChannelInfo{Peer: alias, Phase: peer.ch.Phase().String(), Version: peer.ch.State().Version,
				MyBalanceOnChain: onChainBalsEth[0], PeersBalanceOnChain: onChainBalsEth[1],
				MyBalancePayChan: bals[0], PeersBalancePayChan: bals[1], ChannelID: peer.ch.ID()}

			ir.ChannelInfos = append(ir.ChannelInfos, ci)
		}
	}
	return ir, nil
}

//Exit does does close the client node session, finally
func (n *node) Exit() error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	log.Info().Msg("Exiting...")

	return n.client.Close()
}

//ExistsPeer does evaluate if an used alias is available in the peers session object
func (n *node) ExistsPeer(alias string) bool {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	return n.peers[alias] != nil
}

// PrintfAsync prints the given message for an asynchronous event. More
// precisely, the message is prepended with a newline and appended with the
// command prefix.
func PrintfAsync(format string, a ...interface{}) {
	//fmt.Printf(format, a...)
	log.Info().Time("real-log-time", time.Now()).Str("log-mode", "asnycronous").Msgf(format, a...)
}
