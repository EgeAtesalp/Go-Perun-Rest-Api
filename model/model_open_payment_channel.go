/*
 * ODS
 *
 * This is a simple REST API to access Block Chain on Ethereum and handling Smart Contracts and Payment Channel as well.
 *
 * API version: 1.0.0
 * Contact: u.kuehn@tu-berlin.de
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

//OpenPaymentChannel does
type OpenPaymentChannel struct {
	Target string `json:"target,omitempty"`

	Contract string `json:"contract,omitempty"`

	OwnBalance uint64 `json:"ownBalance,omitempty"`

	TheirsBalance uint64 `json:"theirsBalance,omitempty"`
}
