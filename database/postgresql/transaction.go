package postgresql

import (
	"fmt"

	dbtypes "github.com/MonikaCat/njuno/database/types"
	"github.com/MonikaCat/njuno/types"
	"github.com/MonikaCat/njuno/types/config"
	"github.com/lib/pq"
)

// SaveTx implements database.Database
func (db *Database) SaveTx(tx types.TxResponse) error {
	var partitionID int64
	partitionSize := config.Cfg.Database.PartitionSize
	if partitionSize > 0 {
		partitionID = tx.Height / partitionSize
		err := db.createPartitionIfNotExists("transaction", partitionID)
		if err != nil {
			return fmt.Errorf("error while creating partition with id %d : %s", partitionID, err)
		}
	}

	return db.saveTxInsidePartition(tx, partitionID)
}

// -------------------------------------------------------------------------------------------------------------------

// saveTxInsidePartition stores the given transaction inside the partition having the given id
func (db *Database) saveTxInsidePartition(tx types.TxResponse, partitionId int64) error {
	sqlStatement := `
INSERT INTO transaction 
(hash, height, memo, signatures, fee, gas, partition_id) 
VALUES ($1, $2, $3, $4, $5, $6, $7) 
ON CONFLICT (hash, partition_id) DO UPDATE 
	SET height = excluded.height, 
		memo = excluded.memo, 
		signatures = excluded.signatures, 
		fee = excluded.fee,
		gas = excluded.gas`

	if tx.Height != 0 {
		_, err := db.Sql.Exec(sqlStatement,
			tx.Hash, tx.Height, tx.Memo, pq.Array(dbtypes.NewDBSignatures(tx.Signatures)),
			pq.Array(dbtypes.NewDbCoins(tx.Fee.Amount)), tx.Fee.Gas, partitionId)
		if err != nil {
			return fmt.Errorf("error while storing transaction with hash %s : %s", tx.Hash, err)
		}
	}

	return nil
}
