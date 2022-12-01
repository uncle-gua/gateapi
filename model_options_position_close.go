/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type OptionsPositionClose struct {
	// Position close time
	Time float64 `json:"time,omitempty"`
	// Options contract name
	Contract string `json:"contract,omitempty"`
	// Position side, long or short
	Side string `json:"side,omitempty"`
	// PNL
	Pnl string `json:"pnl,omitempty"`
	// Text of close order
	Text string `json:"text,omitempty"`
	// settlement size
	SettleSize string `json:"settle_size,omitempty"`
}
