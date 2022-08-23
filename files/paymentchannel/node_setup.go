package paymentchannel

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"restapidemo/mdbal"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	echannel "perun.network/go-perun/backend/ethereum/channel"
	ewallet "perun.network/go-perun/backend/ethereum/wallet"
	pkeystore "perun.network/go-perun/backend/ethereum/wallet/keystore"
	"perun.network/go-perun/channel/persistence/keyvalue"
	"perun.network/go-perun/client"
	plog "perun.network/go-perun/log"
	"perun.network/go-perun/pkg/sortedkv/leveldb"
	appWallet "perun.network/go-perun/wallet"
	wtest "perun.network/go-perun/wallet/test"
	wirenet "perun.network/go-perun/wire/net"
	"perun.network/go-perun/wire/net/simple"
)

//Session contains backend and ethereumbackend for an user alias
type Session struct {
	Alias           string
	Backend         *node
	EthereumBackend *ethclient.Client
	ListenerHost    string
	Listener        *simple.Listener
	Peers           map[string]*netConfigEntry
}

var newRand = rand.New(rand.NewSource(42424242))
var appDefAddr appWallet.Address = wtest.NewRandomAddress(newRand)
var odsApp OdsPayment = OdsPayment{Addr: appDefAddr}

//SetupForValidation prepares a Node and return it data
func SetupForValidation(config *Config) (sess Session, err error) {
	err = nil

	if sess.EthereumBackend, err = ethclient.Dial(config.Chain.URL); err != nil {
		log.Error().Msgf("error dial ethereum : %s", err)
		//plog.WithError(err).Fatalln("Could not connect to ethereum Node.")
		if err != nil {
			return sess, err
		}
	}

	// TODO: replace with setupLightNode (remove listener etc.)
	sess.Backend, err = setupFullNode(sess, config)

	if err != nil {
		if strings.Contains(err.Error(), "SetupForValidation - bind: address already in use") {
			log.Error().Msg("SetupForValidation - Ignoring bind error.")
		} else {
			log.Error().Msgf("SetupForValidation - Error happen: %s", err)
			return sess, err
		}
	}
	return sess, err
}

//SetupAndConnect prepares a Node and return it data
func SetupAndConnect(config *Config) (sess Session, err error) {
	err = nil

	if sess.EthereumBackend, err = ethclient.Dial(config.Chain.URL); err != nil {
		log.Error().Msgf("error dial ethereum : %s", err)
		//plog.WithError(err).Fatalln("Could not connect to ethereum Node.")
		if err != nil {
			return sess, err
		}
	}

	sess.Backend, sess.Listener, err = newNode(sess, config)

	if err != nil {
		if strings.Contains(err.Error(), "SetupAndConnect - bind: address already in use") {
			log.Error().Msg("SetupAndConnect - Ignoring bind error.")
		} else {
			log.Error().Msgf("SetupAndConnect - Error happen: %s", err)
			return sess, err
		}
	}

	host := config.Node.IP + ":" + strconv.Itoa(int(config.Node.Port))
	log.Info().Msg("listening  ...")
	//n.plog.WithField("host", host).Trace("Listening for connections")
	sess.ListenerHost = host
	sess.Peers = config.Peers

	return sess, err
}

func setupFullNode(sess Session, config *Config) (*node, error) {
	wallet, acc, err := importAccount(config.SecretKey)
	if err != nil {
		log.Error().Msgf("error importing secret key in ethereum : %s", err)
		return nil, errors.WithMessage(err, "importing secret key")
	}

	if GlobalNodePeers == nil {
		GlobalNodePeers = make(map[string]*NodePeers)
		log.Info().Msg("global node peers created")
	}

	nodePeers := make(map[string]*peer)

	dialer := simple.NewTCPDialer(config.Node.DialTimeout)
	signer := types.NewEIP155Signer(big.NewInt(2000420101))

	n := &node{
		plog:    plog.Get(),
		onChain: acc,
		wallet:  wallet,
		dialer:  dialer,
		cb:      echannel.NewContractBackend(sess.EthereumBackend, pkeystore.NewTransactor(*wallet, signer)),
		peers:   nodePeers,
	}

	err = n.setup()
	if err != nil {
		log.Info().Msgf("error importing while setup node : %s", err)
		return nil, errors.WithMessage(err, "node setup and contracts validation")
	}

	return n, nil
}

func newNode(sess Session, config *Config) (*node, *simple.Listener, error) {
	host := config.Node.IP + ":" + strconv.Itoa(int(config.Node.Port))

	n, err := setupFullNode(sess, config)
	if err != nil {
		return nil, nil, errors.WithMessage(err, "could not create the node")
	}

	GlobalNodePeers[config.Alias] = &NodePeers{
		Peers: n.peers,
	}

	sess.Listener, err = simple.NewTCPListener(host)
	if err != nil {
		return n, sess.Listener, errors.WithMessage(err, "could not start tcp listener")
	}

	n.client.OnNewChannel(n.setupChannel)
	if err := n.setupPersistence(); err != nil {
		log.Info().Msgf("error setting up persitence: %s", err)
		return n, sess.Listener, errors.WithMessage(err, "setting up persistence")
	}

	go n.client.Handle(n, n)
	go n.bus.Listen(sess.Listener)
	plog.WithField("host", host).Trace("Listening for connections")
	n.PrintConfig()

	return n, sess.Listener, nil
}

func (n *node) setupBlockchain() error {
	if err := n.setupContracts(); err != nil {
		log.Info().Msgf("error while setting up contracts : %s", err)
		return errors.WithMessage(err, "setting up contracts")
	}

	n.offChain = n.wallet.NewAccount()
	plog.WithField("off-chain", n.offChain.Address()).Info("Generating account")
	err := n.setupContracts()

	if err != nil {
		log.Info().Msgf("could not setup contract : %s", err)
		return errors.WithMessage(err, "could not setup contract")
	}

	return nil
}

// setup does:
//  - Create a new offChain account.
//  - Create a client with the Node's dialer, funder, adjudicator and wallet.
//  - Setup a TCP listener for incoming connections. (done before)
//  - Load or create the database and setting up persistence with it.
//  - Set the OnNewChannel, Proposal and Update handler.
//  - Print the configuration.
func (n *node) setup() error {
	err := n.setupBlockchain()

	if err != nil {
		log.Info().Msgf("error while setting up blockchain : %s", err)
		return errors.WithMessage(err, "setting up blockchain")
	}

	n.bus = wirenet.NewBus(n.onChain, n.dialer)

	n.client, err = client.New(n.onChain.Address(), n.bus, n.funder, n.adjudicator, n.wallet)

	if err != nil {
		log.Info().Msgf("could not setup client : %s \n", err)
		return errors.WithMessage(err, "could not setup client")
	}

	return nil
}

func (n *node) setupContracts() error {
	var adjAddr common.Address
	var assAddr common.Address
	var err error

	log.Info().Msg("💭 Validating contracts...")

	switch contractSetup := tsconfig.config.Chain.contractSetup; contractSetup {
	case contractSetupOptionValidate:
		if adjAddr, err = validateAdjudicator(n.cb); err == nil { // validate adjudicator
			assAddr, err = validateAssetHolder(n.cb, adjAddr) // validate asset holder
		}
	case contractSetupOptionDeploy:
		if adjAddr, err = deployAdjudicator(n.cb, n.onChain.Account); err == nil { // deploy adjudicator
			assAddr, err = deployAssetHolder(n.cb, adjAddr, n.onChain.Account) // deploy asset holder
		}
	case contractSetupOptionValidateOrDeploy:
		if adjAddr, err = validateAdjudicator(n.cb); err != nil { // validate adjudicator
			log.Info().Msg("❌ Adjudicator invalid")
			adjAddr, err = deployAdjudicator(n.cb, n.onChain.Account) // deploy adjudicator
		}

		if err == nil {
			/* update adjudicator as default in db */

			if assAddr, err = validateAssetHolder(n.cb, adjAddr); err != nil { // validate asset holder
				log.Info().Msg("❌ Asset holder invalid")
				assAddr, err = deployAssetHolder(n.cb, adjAddr, n.onChain.Account) // deploy asset holder
			}

			if err == nil {
				/* update assetholder as default in db */
			}
		}
	default:
		// unsupported setup method
		var msg = fmt.Sprintf("Unsupported contract setup method '%s'.", contractSetup)
		log.Info().Msg(msg)
		err = errors.New(msg)
	}

	log.Info().Msg("✅ Contracts validated.")

	if err != nil {
		return errors.WithMessage(err, "contract setup failed")
	}

	n.adjAddr = adjAddr
	n.assetAddr = assAddr
	recvAddr := ewallet.AsEthAddr(n.onChain.Address())
	n.adjudicator = echannel.NewAdjudicator(n.cb, n.adjAddr, recvAddr, n.onChain.Account)
	n.asset = (*ewallet.Address)(&n.assetAddr)

	/*
	 check current db settings (might change while setupContracts)
	*/
	mdbal.MariaDbConn.Open()
	changed := false
	contracts, err := mdbal.MariaDbConn.ReadContracts("default")
	if contracts.Name != "default" {
		// no entry exists
		contracts.Name = "default"
		changed = true
	}

	if contracts.Adjudicator != n.adjAddr.String() {
		contracts.Adjudicator = n.adjAddr.String()
		changed = true
	}

	if contracts.AssetHolder != n.assetAddr.String() {
		contracts.AssetHolder = n.assetAddr.String()
		changed = true
	}

	if changed {
		mdbal.MariaDbConn.CreateOrUpdateContracts(contracts)
	}

	contracts, err := mdbal.MariaDbConn.ReadContracts("default")
	//TODO: return data
	mdbal.MariaDbConn.Close()
	// end persitence contracts

	log.Info().Msgf("Set Contracts Adj %s Asset %s", n.adjAddr, n.assetAddr)

	accounts := map[echannel.Asset]accounts.Account{ewallet.Address(n.assetAddr): n.onChain.Account}
	depositors := map[echannel.Asset]echannel.Depositor{ewallet.Address(n.assetAddr): new(echannel.ETHDepositor)}
	n.funder = echannel.NewFunder(n.cb, accounts, depositors)

	return nil
}

func (n *node) setupPersistence() error {
	if tsconfig.config.Node.PersistenceEnabled {
		log.Info().Msg("Starting persistence")
		db, err := leveldb.LoadDatabase(tsconfig.config.Node.PersistencePath)
		if err != nil {
			log.Info().Msgf("error crete/load database : %s", err)
			return errors.WithMessage(err, "creating/loading database")
		}
		persister := keyvalue.NewPersistRestorer(db)
		/*if err != nil {
			return errors.WithMessage(err, "creating PersistRestorer")
		}*/
		n.client.EnablePersistence(persister)
		ctx, cancel := context.WithTimeout(context.Background(), tsconfig.config.Node.ReconnecTimeout)
		defer cancel()
		if err := n.client.Restore(ctx); err != nil {
			plog.WithError(err).Warn("Could not restore client")
			// return the error.
			log.Error().Msg("Could not restore client")
		}

	} else {
		log.Info().Msg("Persistence disabled")
	}
	return nil
}

func validateAdjudicator(cb echannel.ContractBackend) (common.Address, error) {
	log.Info().Msg("🌐 Validate adjudicator")

	ctx, cancel := newTransactionContext()
	defer cancel()

	adjAddr := tsconfig.config.Chain.adjudicator
	return adjAddr, echannel.ValidateAdjudicator(ctx, cb, adjAddr)
}

func validateAssetHolder(cb echannel.ContractBackend, adjAddr common.Address) (common.Address, error) {
	log.Info().Msg("🌐 Validate asset holder")

	ctx, cancel := newTransactionContext()
	defer cancel()

	assAddr := tsconfig.config.Chain.assetholder
	return assAddr, echannel.ValidateAssetHolderETH(ctx, cb, assAddr, adjAddr)
}

// deployAdjudicator deploys the Adjudicator to the blockchain and returns its address
// or an error.
func deployAdjudicator(cb echannel.ContractBackend, acc accounts.Account) (common.Address, error) {
	log.Info().Msg("🌐 Deploying adjudicator")
	ctx, cancel := context.WithTimeout(context.Background(), tsconfig.config.Chain.TxTimeout)
	defer cancel()
	adjAddr, err := echannel.DeployAdjudicator(ctx, cb, acc)
	return adjAddr, errors.WithMessage(err, "deploying eth adjudicator")
}

// deployAssetHolder deploys the Assetholder to the blockchain and returns its address
// or an error. Needs an Adjudicator address as second argument.
func deployAssetHolder(cb echannel.ContractBackend, adjudicator common.Address, acc accounts.Account) (common.Address, error) {
	log.Info().Msg("🌐 Deploying asset holder")
	ctx, cancel := context.WithTimeout(context.Background(), tsconfig.config.Chain.TxTimeout)
	defer cancel()
	asset, err := echannel.DeployETHAssetholder(ctx, cb, adjudicator, acc)
	return asset, errors.WithMessage(err, "deploying eth assetholder")
}

// importAccount is a helper method to import secret keys until we have the ethereum wallet done.
func importAccount(secret string) (*pkeystore.Wallet, *pkeystore.Account, error) {
	ks := keystore.NewKeyStore(tsconfig.config.WalletPath, 2, 1)
	sk, err := crypto.HexToECDSA(secret[2:])
	if err != nil {
		log.Info().Msgf("error decoding secret : %s\n", err)
		return nil, nil, errors.WithMessage(err, "decoding secret key")
	}

	var ethAcc accounts.Account
	addr := crypto.PubkeyToAddress(sk.PublicKey)
	if ethAcc, err = ks.Find(accounts.Account{Address: addr}); err != nil {
		ethAcc, err = ks.ImportECDSA(sk, "")
		if err != nil && errors.Cause(err).Error() != "account already exists" {
			log.Info().Msgf("error importing secret : %s\n", err)
			return nil, nil, errors.WithMessage(err, "importing secret key")
		}
	}

	wallet, err := pkeystore.NewWallet(ks, "")
	if err != nil {
		log.Info().Msgf("error creating wallet : %s\n", err)
		return nil, nil, errors.WithMessage(err, "creating wallet")
	}

	wAcc := pkeystore.NewAccountFromEth(wallet, &ethAcc)
	acc, err := wallet.Unlock(wAcc.Address())
	return wallet, acc.(*pkeystore.Account), err
}

func newTransactionContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), tsconfig.config.Chain.TxTimeout)
}

//PrintConfig does
func (n *node) PrintConfig() error {
	log.Printf(
		"Alias: %s\n"+
			"Listening: %s:%d\n"+
			"ETH RPC URL: %s\n"+
			"PeerID: %s\n"+
			"OffChain: %s\n"+
			"ETHAssetHolder: %s\n"+
			"Adjudicator: %s\n"+
			"", tsconfig.config.Alias, tsconfig.config.Node.IP, tsconfig.config.Node.Port,
		tsconfig.config.Chain.URL, n.onChain.Address().String(), n.offChain.Address().String(),
		n.assetAddr.String(), n.adjAddr.String())

	log.Info().Msg("Known peers:")
	for alias, peer := range tsconfig.config.Peers {
		log.Info().Msgf("[%s]\t[%v]\t[%s:%d]\n", alias, peer.peerID, peer.Hostname, peer.Port)
	}
	log.Info().Msg("Startup done successfully")
	return nil
}
