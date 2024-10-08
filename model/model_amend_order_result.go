/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// Batch order modification results
type AmendOrderResult struct {
	// Order ID
	OrderId string `json:"order_id,omitempty"`
	// Trade amount
	Amount string `json:"amount,omitempty"`
	// Order price
	Price string `json:"price,omitempty"`
	// The custom data that the user remarked when amending the order
	AmendText string `json:"amend_text,omitempty"`
	// Update success status
	Succeeded bool `json:"succeeded,omitempty"`
	// Error indicator for failed modifications; empty when successful
	Label string `json:"label,omitempty"`
	// Error description for failed modifications; empty when successful
	Message string `json:"message,omitempty"`
	// Account types， spot - spot account, margin - margin account, unified - unified account, cross_margin - cross margin account.Portfolio margin accounts can only be set to `cross_margin`
	Account string `json:"account,omitempty"`
}
