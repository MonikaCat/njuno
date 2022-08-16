package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InflationResponse contains the data of the current inflation rate
type InflationResponse struct {
	Inflation string `json:"inflation" yaml:"inflation"`
}

// ----------------------------------------------------------------------------------------------------------

// StakingPool contains the data of the staking pool at the given height
type StakingPool struct {
	BondedTokens    sdk.Int
	NotBondedTokens sdk.Int
	Height          int64
}

// NewStakingPool allows to build a new StakingPool instance
func NewStakingPool(bondedTokens sdk.Int, notBondedTokens sdk.Int, height int64) *StakingPool {
	return &StakingPool{
		BondedTokens:    bondedTokens,
		NotBondedTokens: notBondedTokens,
		Height:          height,
	}
}
