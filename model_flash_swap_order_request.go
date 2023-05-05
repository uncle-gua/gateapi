/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Parameters of flash swap order creation
type FlashSwapOrderRequest struct {
	// Preview result ID
	PreviewId string `json:"preview_id"`
	// Currency to sell which can be retrieved from supported currency list API `GET /flash_swap/currencies`
	SellCurrency string `json:"sell_currency"`
	// Amount to sell (based on the preview result)
	SellAmount string `json:"sell_amount"`
	// Currency to buy which can be retrieved from supported currency list API `GET /flash_swap/currencies`
	BuyCurrency string `json:"buy_currency"`
	// Amount to buy (based on the preview result)
	BuyAmount string `json:"buy_amount"`
}
