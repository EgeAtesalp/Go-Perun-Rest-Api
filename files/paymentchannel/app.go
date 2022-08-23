package paymentchannel

import (
	"io"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"

	"github.com/rs/zerolog/log"
	"perun.network/go-perun/channel"
	perunio "perun.network/go-perun/pkg/io"
	"perun.network/go-perun/wallet"
)

// OdsApp is a payment app.
type OdsPayment struct {
	Addr wallet.Address
}

// Def returns the address of this payment app.
func (a *OdsPayment) Def() wallet.Address {
	log.Info().Msg("... initialzes my app ods def")
	return a.Addr
}

// DecodeData does try read returns PaymentData from the reader or returns new NoData.
func (a *OdsPayment) DecodeData(r io.Reader) (channel.Data, error) {
	log.Info().Msg("decode data ...")
	var pd PaymentData = PaymentData{}
	err := perunio.Decode(r, &pd.referenceToVersion,
		&pd.senderAlias, &pd.receiverAlias, &pd.message, &pd.balanceTransfered,
		&pd.timeSent, &pd.timeReceived, &pd.part)

	log.Info().Str("decoded-data", spew.Sdump(pd)).Msg("decoded")
	return pd.Clone(), err
}

// ValidTransition checks that money flows only from the actor to the other
// participants.
func (a *OdsPayment) ValidTransition(_ *channel.Params, from, to *channel.State, actor channel.Index) error {
	log.Info().Msg("reached my app ...")
	assertNoDataOrPaymentData(to)
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

// ValidInit panics if State.Data is not *NoData and returns nil otherwise. Any
// valid allocation forms a valid initial state.
func (a *OdsPayment) ValidInit(_ *channel.Params, s *channel.State) error {
	log.Info().Msg("valid init ...")
	assertNoDataOrPaymentData(s)
	return nil
}

func assertNoDataOrPaymentData(s *channel.State) {
	_, ok := s.Data.(*PaymentData)
	if !ok {
		log.Info().Msgf("payment app must have either no data (new(NoData)) or (new(PaymentData)), but has type %T", s.Data)
	} else {
		log.Info().Msg("PaymentData detected.")
	}
}
