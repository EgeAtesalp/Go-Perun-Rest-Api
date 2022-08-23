/*
 * ODS
 *
 * This is a simple REST API to access Block Chain on Ethereum and handling Smart Contracts and Payment Channel as well.
 *
 * API version: 1.0.0
 * Contact: u.kuehn@tu-berlin.de
 */

package model

import "math/big"

//Contract contains smart contract data
type InfoResult struct {
	ShortMessage string `json:"shortMessage,omitempty"`

	LongMessage string `json:"longMessage,omitempty"`

	CurrentEthBalance *big.Int `json:"currentEthBalance,omitempty"`

	CurrentEthBalanceWithPending *big.Int `json:"currentEthBalanceWithPending,omitempty"`

	CurrentEthBalanceInWei *big.Int `json:"currentEthBalanceInWei,omitempty"`

	CurrentEthBalanceWithPendingInWei *big.Int `json:"currentEthBalanceWithPendingInWei,omitempty"`

	UnreadMessages uint `json:"unreadMessages,omitempty"`

	ChannelInfos []ChannelInfo `json:"channelInfos,omitempty"`

	IsConnected bool `json:"isConnected,omitempty"`

	PendingTransactionCount uint `json:"pendingTransactionCount,omitempty"`
}

type ChannelInfo struct {
	Peer string `json:"peer,omitempty"`

	Phase string `json:"phase,omitempty"`

	Version uint64 `json:"version,omitempty"`

	MyBalancePayChan *big.Float `json:"myBalancePayChan,omitempty"`

	PeersBalancePayChan *big.Float `json:"peersBalancePayChan,omitempty"`

	MyBalanceOnChain *big.Float `json:"myBalanceOnChain,omitempty"`

	PeersBalanceOnChain *big.Float `json:"peersBalanceOnChain,omitempty"`

	ChannelID [32]byte `json:"cannelId,omitempty"`
}
