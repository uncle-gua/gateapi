/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type OptionsMyTrade struct {
	// Trade ID
	Id int64 `json:"id,omitempty"`
	// Trading time
	CreateTime float64 `json:"create_time,omitempty"`
	// Options contract name
	Contract string `json:"contract,omitempty"`
	// Order ID related
	OrderId int32 `json:"order_id,omitempty"`
	// Trading size
	Size int64 `json:"size,omitempty"`
	// Trading price (quote currency)
	Price string `json:"price,omitempty"`
	// Underlying price (quote currency)
	UnderlyingPrice string `json:"underlying_price,omitempty"`
	// Trade role. Available values are `taker` and `maker`
	Role string `json:"role,omitempty"`
}
