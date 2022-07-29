package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// QueryAllBalancesResponse contains the account balance data
type QueryAllBalancesResponse struct {
	Balances   sdk.Coins     `son:"balances"`
	Pagination *PageResponse `json:"pagination,omitempty"`
}

type PageResponse struct {
	NextKey []byte `json:"next_key,omitempty"`
	Total   string `json:"total,omitempty"`
}
