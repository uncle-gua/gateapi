/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

type MyFuturesTradeTimeRange struct {
	// Trade ID
	TradeId string `json:"trade_id,omitempty"`
	// Trading time
	CreateTime float64 `json:"create_time,omitempty"`
	// Futures contract
	Contract string `json:"contract,omitempty"`
	// Order ID related
	OrderId string `json:"order_id,omitempty"`
	// Trading size
	Size int64 `json:"size,omitempty"`
	// Trading price
	Price string `json:"price,omitempty"`
	// Trade role. Available values are `taker` and `maker`
	Role string `json:"role,omitempty"`
	// User defined information
	Text string `json:"text,omitempty"`
	// Fee deducted
	Fee string `json:"fee,omitempty"`
	// Points used to deduct fee
	PointFee string `json:"point_fee,omitempty"`
}
