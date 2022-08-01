package postgresql

// GetLastPruned implements database.PruningDb
func (db *Database) GetLastPruned() (int64, error) {
	var lastPrunedHeight int64
	err := db.Sql.QueryRow(`SELECT coalesce(MAX(last_pruned_height),0) FROM pruning LIMIT 1;`).Scan(&lastPrunedHeight)
	return lastPrunedHeight, err
}

// -------------------------------------------------------------------------------------------------------------------

// Prune implements database.PruningDb
func (db *Database) Prune(height int64) error {
	_, err := db.Sql.Exec(`DELETE FROM pre_commit WHERE height = $1`, height)
	if err != nil {
		return err
	}

	return err
}

// -------------------------------------------------------------------------------------------------------------------

// StoreLastPruned implements database.PruningDb
func (db *Database) StoreLastPruned(height int64) error {
	_, err := db.Sql.Exec(`DELETE FROM pruning`)
	if err != nil {
		return err
	}

	_, err = db.Sql.Exec(`INSERT INTO pruning (last_pruned_height) VALUES ($1)`, height)
	return err
}
