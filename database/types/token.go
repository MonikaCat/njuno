package types

import (
	"database/sql"

	"github.com/lib/pq"
)

// TokenUnitRow represents a single token unit row stored inside the database
type TokenUnitRow struct {
	TokenName string         `db:"token_name"`
	Denom     string         `db:"denom"`
	Exponent  int            `db:"exponent"`
	Aliases   pq.StringArray `db:"aliases"`
	PriceID   sql.NullString `db:"price_id"`
}

// TokenUnitRow represents a single token row stored inside the database
type TokenRow struct {
	Name       string `db:"name"`
	TradedUnit string `db:"traded_unit"`
}
