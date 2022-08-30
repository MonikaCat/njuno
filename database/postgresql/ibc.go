package postgresql

import (
	"encoding/json"
	"fmt"

	"github.com/forbole/njuno/types"
)

// SaveIBCTransferParams allows to store the given ibc transfer params inside the database
func (db *Database) SaveIBCTransferParams(params *types.IBCTransferParams) error {
	paramsBz, err := json.Marshal(&params.Params)
	if err != nil {
		return fmt.Errorf("error while marshaling ibc transfer params: %s", err)
	}

	stmt := `
INSERT INTO ibc_transfer_params (params, height) 
VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE 
    SET params = excluded.params,
        height = excluded.height
WHERE ibc_transfer_params.height <= excluded.height`

	_, err = db.Sql.Exec(stmt, string(paramsBz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing ibc transfer params: %s", err)
	}

	return nil
}
