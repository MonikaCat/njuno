package postgresql

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dbtypes "github.com/forbole/njuno/database/types"
	"github.com/lib/pq"
)

// SaveSupply allows to store total supply for a given height
func (db *Database) SaveSupply(coins sdk.Coins, height int64) error {
	query := `
INSERT INTO supply (coins, height) 
VALUES ($1, $2) 
ON CONFLICT (one_row_id) DO UPDATE 
    SET coins = excluded.coins,
    	height = excluded.height
WHERE supply.height <= excluded.height`

	_, err := db.Sql.Exec(query, pq.Array(dbtypes.NewDbCoins(coins)), height)
	if err != nil {
		return fmt.Errorf("error while storing supply: %s", err)
	}

	return nil
}
