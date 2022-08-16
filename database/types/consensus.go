package types

import (
	"database/sql"
	"time"
)

// BlockRow represents a single block row stored inside the database
type BlockRow struct {
	Height          int64          `db:"height"`
	Hash            string         `db:"hash"`
	TxNum           int64          `db:"num_txs"`
	TotalGas        int64          `db:"total_gas"`
	ProposerAddress sql.NullString `db:"proposer_address"`
	PreCommitsNum   int64          `db:"pre_commits"`
	Timestamp       time.Time      `db:"timestamp"`
}

// _________________________________________________________

// GenesisRow represents a single genesis row stored inside the database
type GenesisRow struct {
	OneRowID      bool      `db:"one_row_id"`
	ChainID       string    `db:"chain_id"`
	Time          time.Time `db:"time"`
	InitialHeight int64     `db:"initial_height"`
}

// NewGenesisRow allows to create new GenesisRow instance
func NewGenesisRow(chainID string, time time.Time, initialHeight int64) GenesisRow {
	return GenesisRow{
		OneRowID:      true,
		ChainID:       chainID,
		Time:          time,
		InitialHeight: initialHeight,
	}
}

// Equal tells whether v and w represent the same rows
func (v GenesisRow) Equal(w GenesisRow) bool {
	return v.OneRowID == w.OneRowID &&
		v.ChainID == w.ChainID &&
		v.Time == w.Time &&
		v.InitialHeight == w.InitialHeight
}
