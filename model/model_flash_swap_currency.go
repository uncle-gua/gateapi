/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// Currencies supported in flash swap
type FlashSwapCurrency struct {
	// Currency name
	Currency string `json:"currency,omitempty"`
	// Minimum amount required in flash swap
	MinAmount string `json:"min_amount,omitempty"`
	// Maximum amount allowed in flash swap
	MaxAmount string `json:"max_amount,omitempty"`
	// Currencies which can be swapped to from this currency
	Swappable []string `json:"swappable,omitempty"`
}
