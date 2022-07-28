package types

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"github.com/MonikaCat/njuno/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lib/pq"
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

// DbCoin represents the information stored inside the database about a single coin
type DbCoin struct {
	Denom  string
	Amount string
}

// NewDbCoin builds a DbCoin starting from an SDK Coin
func NewDbCoin(coin sdk.Coin) DbCoin {
	return DbCoin{
		Denom:  coin.Denom,
		Amount: coin.Amount.String(),
	}
}

// Value implements driver.Value
func (coin *DbCoin) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%s)", coin.Denom, coin.Amount), nil
}

// _________________________________________________________

// DbCoins represents an array of coins
type DbCoins []*DbCoin

// NewDbCoins build a new DbCoins object starting from an array of coins
func NewDbCoins(coins sdk.Coins) DbCoins {
	dbCoins := make([]*DbCoin, 0)
	for _, coin := range coins {
		dbCoins = append(dbCoins, &DbCoin{Amount: coin.Amount.String(), Denom: coin.Denom})
	}
	return dbCoins
}

// _________________________________________________________

// NewDBSignatures returns signatures in string array
func NewDBSignatures(signaturesList []types.TxSignatures) []string {
	var signatures []string
	for _, index := range signaturesList {
		signatures = append(signatures, index.Signature)
	}
	return signatures
}

// _________________________________________________________

type GenesisRow struct {
	OneRowID      bool      `db:"one_row_id"`
	ChainID       string    `db:"chain_id"`
	Time          time.Time `db:"time"`
	InitialHeight int64     `db:"initial_height"`
}

func NewGenesisRow(chainID string, time time.Time, initialHeight int64) GenesisRow {
	return GenesisRow{
		OneRowID:      true,
		ChainID:       chainID,
		Time:          time,
		InitialHeight: initialHeight,
	}
}

// _________________________________________________________

func ToNullString(value string) sql.NullString {
	value = strings.TrimSpace(value)
	return sql.NullString{
		Valid:  value != "",
		String: value,
	}
}

// _________________________________________________________

// ValidatorDescriptionRow represent a row in validator_description
type ValidatorDescriptionRow struct {
	ValAddress string         `db:"validator_address"`
	Moniker    sql.NullString `db:"moniker"`
	Details    sql.NullString `db:"details"`
	Height     int64          `db:"height"`
}

// NewValidatorDescriptionRow return a row representing data structure in validator_description
func NewValidatorDescriptionRow(
	valAddress, moniker, details string, height int64,
) ValidatorDescriptionRow {
	return ValidatorDescriptionRow{
		ValAddress: valAddress,
		Moniker:    ToNullString(moniker),
		Details:    ToNullString(details),
		Height:     height,
	}
}

// --------------------------------------------------------------------------------------------------------------------

type TokenUnitRow struct {
	TokenName string         `db:"token_name"`
	Denom     string         `db:"denom"`
	Exponent  int            `db:"exponent"`
	Aliases   pq.StringArray `db:"aliases"`
	PriceID   sql.NullString `db:"price_id"`
}

type TokenRow struct {
	Name       string `db:"name"`
	TradedUnit string `db:"traded_unit"`
}

// --------------------------------------------------------------------------------------------------------------------

// TokenPriceRow represent a row of the table token_price in the database
type TokenPriceRow struct {
	ID        string    `db:"id"`
	Name      string    `db:"unit_name"`
	Price     float64   `db:"price"`
	MarketCap int64     `db:"market_cap"`
	Timestamp time.Time `db:"timestamp"`
}

// NewTokenPriceRow allows to easily create a new NewTokenPriceRow
func NewTokenPriceRow(name string, currentPrice float64, marketCap int64, timestamp time.Time) TokenPriceRow {
	return TokenPriceRow{
		Name:      name,
		Price:     currentPrice,
		MarketCap: marketCap,
		Timestamp: timestamp,
	}
}

// Equals return true if u and v represent the same row
func (u TokenPriceRow) Equals(v TokenPriceRow) bool {
	return u.Name == v.Name &&
		u.Price == v.Price &&
		u.MarketCap == v.MarketCap &&
		u.Timestamp.Equal(v.Timestamp)
}
