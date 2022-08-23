package paymentchannel

import (
	"io"
	"time"

	"github.com/pkg/errors"
	"perun.network/go-perun/channel"
	perunio "perun.network/go-perun/pkg/io"
)

//PaymentDataMethods wraps Data between sender and receiver in the payment channel
type PaymentDataMethods interface {
	GetReferenceToVersion() uint64
	GetSenderAlias() string
	GetReceiverAlias() string
	GetMessage() string
	GetBalanceSend() uint64
	GetTimeSent() time.Time
	GetTimeReceived() time.Time
}

//PaymentData implements Data interface
type PaymentData struct {
	referenceToVersion uint64
	senderAlias        string
	receiverAlias      string
	message            string
	balanceTransfered  uint64
	timeSent           time.Time
	timeReceived       time.Time
	part               string
}

//GetReferenceToVersion returns the reference to the version this payment data does or did answer
func (pd PaymentData) GetReferenceToVersion() uint64 {
	return pd.referenceToVersion
}

//GetSenderAlias returns senders alias
func (pd PaymentData) GetSenderAlias() string {
	return pd.senderAlias
}

//GetReceiverAlias returns receivers alias
func (pd PaymentData) GetReceiverAlias() string {
	return pd.receiverAlias
}

//GetMessage returns message
func (pd PaymentData) GetMessage() string {
	return pd.message
}

//GetBalanceSend returns balance
func (pd PaymentData) GetBalanceSend() uint64 {
	return pd.balanceTransfered
}

//GetTimeSent returns time sent
func (pd PaymentData) GetTimeSent() time.Time {
	return pd.timeSent
}

//GetTimeReceived returns time received
func (pd PaymentData) GetTimeReceived() time.Time {
	return pd.timeReceived
}

//GetPart returns receivers alias
func (pd PaymentData) GetPart() string {
	return pd.part
}

//Clone clones the PaymentData
func (pd PaymentData) Clone() channel.Data {
	clone := PaymentData{
		referenceToVersion: pd.referenceToVersion,
		senderAlias:        pd.senderAlias,
		receiverAlias:      pd.receiverAlias,
		message:            pd.message,
		balanceTransfered:  pd.balanceTransfered,
		timeSent:           pd.timeSent,
		timeReceived:       pd.timeReceived,
		part:               pd.part,
	}
	return clone
}

// Encode encodes a state into an `io.Writer` or returns an `error`
func (pd PaymentData) Encode(w io.Writer) error {
	err := perunio.Encode(w, pd.referenceToVersion,
		pd.senderAlias, pd.receiverAlias, pd.message, pd.balanceTransfered,
		pd.timeSent, pd.timeReceived, pd.part)
	return errors.WithMessage(err, "payment data encode")
}
