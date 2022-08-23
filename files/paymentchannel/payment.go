package paymentchannel

import (
	"context"
	"math/big"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	"perun.network/go-perun/log"

	zlog "github.com/rs/zerolog/log"
)

type (
	paymentChannel struct {
		*client.Channel

		log     log.Logger
		handler chan bool
		res     chan handlerRes
		onFinal func()
		// save the last state to circumvent the `channel.StateMtxd` problem
		lastState *channel.State
	}

	// A handlerRes encapsulates the result of a channel handling request
	handlerRes struct {
		up  client.ChannelUpdate
		err error
	}
)

func newPaymentChannel(ch *client.Channel) *paymentChannel {
	state := ch.State()
	state.App = &odsApp

	if ch.HasApp() {
		zlog.Info().Msg("newPaymentChannel Has App")
	} else {
		zlog.Info().Msg("newPaymentChannel No App found")
	}

	return &paymentChannel{
		Channel:   ch,
		log:       log.WithField("channel", ch.ID()),
		handler:   make(chan bool, 1),
		res:       make(chan handlerRes),
		lastState: state,
	}
}

func (ch *paymentChannel) sendMoneyAndData(amount *big.Int, pd *PaymentData) error {
	zlog.Info().Msgf("Sending money: %s Eth?\n", amount)
	zlog.Info().Str("data", spew.Sdump(pd)).Msg("Sending data")

	return ch.sendUpdate(
		func(state *channel.State) error {
			myPd := pd.Clone().(PaymentData)
			myPd.timeSent = time.Now()
			state.Data = myPd
			transferBal(stateBals(state), ch.Idx(), amount)
			return nil
		}, "sendMoney")
}

func (ch *paymentChannel) sendFinal() error {
	zlog.Info().Msg("Sending final")
	zlog.Info().Msg("Sending final state")
	return ch.sendUpdate(func(state *channel.State) error {
		state.Data = PaymentData{message: "sending final."}
		state.IsFinal = true
		return nil
	}, "final")
}

func (ch *paymentChannel) sendUpdate(update func(*channel.State) error, desc string) error {
	if ch.HasApp() {
		zlog.Info().Msg("sendUpdate Has App")
	} else {
		zlog.Info().Msg("sendUpdate No App found")
	}
	log.Printf("Sending update: %s \n", desc)
	ctx, cancel := context.WithTimeout(context.Background(), tsconfig.config.Channel.Timeout)
	defer cancel()

	whoAmI := ch.Peers()[ch.Idx()]
	// whoAmOth := ch.Peers()[1-ch.Idx()]

	newState := ch.State().Clone()
	myData := newState.Data
	zlog.Info().Str("old-data", spew.Sdump(myData)).Msg("send update ... old data")

	//myNewData := PaymentData{message: "test-test Testmessage ... :-)"}
	/*myNewData.receiverAlias = whoAmOth.String()
	myNewData.senderAlias = whoAmI.String()
	myNewData.timeSent = time.Now()
	myNewData.balanceTransfered = newState.Balances[0][0].Int64()*/
	//fmt.Printf("send update ... new data\n")
	//spew.Dump(myNewData)

	//cloneNew := myData.Clone()
	//fmt.Printf("cloned ...\n")
	//spew.Dump(cloneNew)

	//newState.Data = &myNewData
	//ch.lastState = newState

	// it looks like not allowed to change State/Data here :-(
	err := ch.UpdateBy(ctx, update)
	zlog.Info().Msgf("Sent update: %s, err: %v", desc, err)

	state := ch.State()
	balChanged := newState.Balances[0][0].Cmp(state.Balances[0][0]) != 0

	zlog.Info().Msgf("Sent update: %s, err: %v", desc, err)

	//myData = state.Data.Clone()
	//fmt.Printf("after update ... new data")
	//spew.Dump(myData)

	if balChanged {
		bals := weiToEther(state.Allocation.Balances[0]...)
		zlog.Info().Msgf("💰 %s sent payment. New balance: [My: %v Ξ, Peer: %v Ξ]\n", whoAmI.String(), bals[ch.Idx()], bals[1-ch.Idx()])
	}

	if err == nil {
		ch.lastState = state
	}

	return err
}

func transferBal(bals []channel.Bal, ourIdx channel.Index, amount *big.Int) {
	a := new(big.Int).Set(amount) // local copy because we mutate it
	otherIdx := ourIdx ^ 1
	ourBal := bals[ourIdx]
	otherBal := bals[otherIdx]
	otherBal.Add(otherBal, a)
	ourBal.Sub(ourBal, a)
}

func stateBals(state *channel.State) []channel.Bal {
	return state.Balances[0]
}

func (ch *paymentChannel) Handle(update client.ChannelUpdate, res *client.UpdateResponder) {
	oldBal := stateBals(ch.lastState)
	balChanged := oldBal[0].Cmp(update.State.Balances[0][0]) != 0
	ctx, cancel := context.WithTimeout(context.Background(), tsconfig.config.Channel.Timeout)
	defer cancel()
	if err := assertValidTransition(ch.lastState, update.State, update.ActorIdx); err != nil {
		res.Reject(ctx, "invalid transition")
	} else if err := res.Accept(ctx); err != nil {
		ch.log.Error(errors.WithMessage(err, "handling payment update"))
	}

	whoAmI := ch.Peers()[ch.Idx()]
	myData := update.State.Data.Clone().(PaymentData)
	myData.timeReceived = time.Now()
	zlog.Info().Str("data", spew.Sdump(myData)).Msg("do handle ... data")

	if balChanged {
		bals := weiToEther(update.State.Allocation.Balances[0]...)
		zlog.Info().Msgf("\n💰 %s received payment. New balance: [My: %v Ξ, Peer: %v Ξ]\n", whoAmI.String(), bals[ch.Idx()], bals[1-ch.Idx()])
	} else {
		zlog.Info().Msgf("\n💰 %s received message. Balance unchanged. \n", whoAmI.String())
	}

	ch.lastState = update.State.Clone()
}

// assertValidTransition checks that money flows only from the actor to the
// other participants.
func assertValidTransition(from, to *channel.State, actor channel.Index) error {
	// if !channel.IsNoData(to.Data) {
	// 	return errors.New("channel must not have app data")
	// }
	myData := to.Data.Clone()
	zlog.Info().Str("data", spew.Sdump(myData)).Msg("send update ... transited data")

	for i, asset := range from.Balances {
		for j, bal := range asset {
			if int(actor) == j && bal.Cmp(to.Balances[i][j]) == -1 {
				return errors.Errorf("payer[%d] steals asset %d, so %d < %d", j, i, bal, to.Balances[i][j])
			} else if int(actor) != j && bal.Cmp(to.Balances[i][j]) == 1 {
				return errors.Errorf("payer[%d] reduces participant[%d]'s asset %d", actor, j, i)
			}
		}
	}
	return nil
}

func (ch *paymentChannel) GetBalances() (our, other *big.Int) {
	bals := stateBals(ch.State())
	if len(bals) != 2 {
		return new(big.Int), new(big.Int)
	}
	return bals[ch.Idx()], bals[1-ch.Idx()]
}
