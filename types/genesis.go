package types

import "time"

// Genesis contains the useful information about the genesis
type Genesis struct {
	ChainID       string
	Time          time.Time
	InitialHeight int64
}

// NewGenesis allows to build a new Genesis instance
func NewGenesis(chainID string, startTime time.Time, initialHeight int64) *Genesis {
	return &Genesis{
		ChainID:       chainID,
		Time:          startTime,
		InitialHeight: initialHeight,
	}
}
