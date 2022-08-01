package postgresql

import (
	"fmt"

	dbtypes "github.com/MonikaCat/njuno/database/types"
	"github.com/MonikaCat/njuno/types"
)

// GetTokensPriceID returns the slice of price ids for all tokens stored in db
func (db *Database) GetTokensPriceID() ([]string, error) {
	query := `SELECT * FROM token_unit`

	var dbUnits []dbtypes.TokenUnitRow
	err := db.Sqlx.Select(&dbUnits, query)
	if err != nil {
		return nil, err
	}

	var units []string
	for _, unit := range dbUnits {
		if unit.PriceID.String != "" {
			units = append(units, unit.PriceID.String)
		}
	}

	return units, nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveTokensPrices stores the latest tokens price
func (db *Database) SaveTokensPrice(prices []types.TokenPrice) error {
	if len(prices) == 0 {
		return nil
	}

	query := `INSERT INTO token_price (unit_name, price, market_cap, timestamp) VALUES`
	var param []interface{}

	for i, ticker := range prices {
		vi := i * 4
		query += fmt.Sprintf("($%d,$%d,$%d,$%d),", vi+1, vi+2, vi+3, vi+4)
		param = append(param, ticker.UnitName, ticker.Price, ticker.MarketCap, ticker.Timestamp)
	}

	query = query[:len(query)-1] // Remove trailing ","
	query += `
ON CONFLICT (unit_name) DO UPDATE 
	SET price = excluded.price,
	    market_cap = excluded.market_cap,
	    timestamp = excluded.timestamp
WHERE token_price.timestamp <= excluded.timestamp`

	_, err := db.Sql.Exec(query, param...)
	if err != nil {
		return fmt.Errorf("error while saving tokens prices: %s", err)
	}

	return nil
}
