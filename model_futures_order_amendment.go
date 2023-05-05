/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type FuturesOrderAmendment struct {
	// New order size, including filled part.  - If new size is less than or equal to filled size, the order will be cancelled. - Order side must be identical to the original one. - Close order size cannot be changed. - For reduce only orders, increasing size may leads to other reduce only orders being cancelled. - If price is not changed, decreasing size will not change its precedence in order book, while increasing will move it to the last at current price.
	Size int64 `json:"size,omitempty"`
	// New order price.
	Price string `json:"price,omitempty"`
	// Custom info during amending order
	AmendText string `json:"amend_text,omitempty"`
}
