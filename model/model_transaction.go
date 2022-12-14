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

//Transaction does
type Transaction struct {
	Sender string `json:"sender,omitempty"`

	Receiver string `json:"receiver,omitempty"`

	Value int64 `json:"value,omitempty"`
}
