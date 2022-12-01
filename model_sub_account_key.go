/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type SubAccountKey struct {
	// User ID
	UserId string `json:"user_id,omitempty"`
	// API key name
	Name  string         `json:"name,omitempty"`
	Perms []ApiV4KeyPerm `json:"perms,omitempty"`
	// ip white list (list will be removed if no value is passed)
	IpWhitelist []string `json:"ip_whitelist,omitempty"`
	// API Key
	Key string `json:"key,omitempty"`
	// State 1 - normal 2 - locked 3 - frozen
	State int32 `json:"state,omitempty"`
	// Creation time
	CreatedAt string `json:"created_at,omitempty"`
	// Last update time
	UpdatedAt string `json:"updated_at,omitempty"`
}
