package postgresql

import (
	"encoding/json"
	"fmt"

	"github.com/MonikaCat/njuno/types"
)

// SaveIBCParams allows to store the given params inside the database
func (db *Database) SaveIBCParams(params *types.IBCParams) error {
	paramsBz, err := json.Marshal(&params.Params)
	if err != nil {
		return fmt.Errorf("error while marshaling mint params: %s", err)
	}

	stmt := `
INSERT INTO ibc_params (params, height) 
VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE 
    SET params = excluded.params,
        height = excluded.height
WHERE ibc_params.height <= excluded.height`

	_, err = db.Sql.Exec(stmt, string(paramsBz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing ibc params: %s", err)
	}

	return nil
}
