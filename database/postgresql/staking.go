package postgresql

import (
	"fmt"

	dbtypes "github.com/MonikaCat/njuno/database/types"
	"github.com/MonikaCat/njuno/types"
)

// GetValidatorDescription returns validators description from database.
func (db *Database) GetValidatorsDescription() ([]types.ValidatorDescription, error) {
	var result []dbtypes.ValidatorDescriptionRow
	stmt := `SELECT * FROM validator_description`

	err := db.Sqlx.Select(&result, stmt)
	if err != nil {
		return nil, nil
	}

	if len(result) == 0 {
		return nil, nil
	}
	var list []types.ValidatorDescription
	for _, index := range result {
		list = append(list,
			types.NewValidatorDescription(index.ValAddress,
				dbtypes.ToString(index.Details),
				dbtypes.ToString(index.Identity),
				dbtypes.ToString(index.Moniker),
				index.Height))
	}

	return list, nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveCommitSignatures implements database.Database
func (db *Database) SaveCommitSignatures(signatures []*types.CommitSig) error {
	if len(signatures) == 0 {
		return nil
	}

	stmt := `INSERT INTO pre_commit (validator_address, height, timestamp, voting_power, proposer_priority) VALUES `

	var sparams []interface{}
	for i, sig := range signatures {
		si := i * 5

		stmt += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d),", si+1, si+2, si+3, si+4, si+5)
		sparams = append(sparams, sig.ValidatorAddress, sig.Height, sig.Timestamp, sig.VotingPower, sig.ProposerPriority)
	}

	stmt = stmt[:len(stmt)-1]
	stmt += " ON CONFLICT (validator_address, timestamp) DO NOTHING"
	_, err := db.Sql.Exec(stmt, sparams...)
	return err
}

// -------------------------------------------------------------------------------------------------------------------

// SaveDoubleSignEvidence saves the given double sign evidence inside the proper tables
func (db *Database) SaveDoubleSignEvidence(evidence types.DoubleSignEvidence) error {
	voteA, err := db.saveDoubleSignVote(evidence.VoteA)
	if err != nil {
		return fmt.Errorf("error while storing double sign vote: %s", err)
	}

	voteB, err := db.saveDoubleSignVote(evidence.VoteB)
	if err != nil {
		return fmt.Errorf("error while storing double sign vote: %s", err)
	}

	stmt := `
INSERT INTO double_sign_evidence (height, vote_a_id, vote_b_id) 
VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`
	_, err = db.Sql.Exec(stmt, evidence.Height, voteA, voteB)
	if err != nil {
		return fmt.Errorf("error while storing double sign evidence: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// saveDoubleSignVote saves the given vote inside the database, returning the row id
func (db *Database) saveDoubleSignVote(vote types.DoubleSignVote) (int64, error) {
	stmt := `
INSERT INTO double_sign_vote 
    (type, height, round, block_id, validator_address, validator_index, signature) 
VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT DO NOTHING RETURNING id`

	var id int64
	err := db.Sql.QueryRow(stmt,
		vote.Type, vote.Height, vote.Round, vote.BlockID, vote.ValidatorAddress, vote.ValidatorIndex, vote.Signature,
	).Scan(&id)
	return id, err
}

// -------------------------------------------------------------------------------------------------------------------

// SaveStakingPool allows to store staking pool values for the given height
func (db *Database) SaveStakingPool(pool *types.StakingPool) error {
	stmt := `
INSERT INTO staking_pool (bonded_tokens, not_bonded_tokens, height) 
VALUES ($1, $2, $3)
ON CONFLICT (one_row_id) DO UPDATE 
    SET bonded_tokens = excluded.bonded_tokens, 
        not_bonded_tokens = excluded.not_bonded_tokens, 
        height = excluded.height
WHERE staking_pool.height <= excluded.height`

	_, err := db.Sql.Exec(stmt, pool.BondedTokens.String(), pool.NotBondedTokens.String(), pool.Height)
	if err != nil {
		return fmt.Errorf("error while storing staking pool: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveValidators implements database.Database
func (db *Database) SaveValidators(validators []types.Validator) error {
	if len(validators) == 0 {
		return nil
	}

	validatorQuery := `INSERT INTO validator (consensus_address, consensus_pubkey) VALUES `

	validatorInfoQuery := `
INSERT INTO validator_info (consensus_address, operator_address, self_delegate_address, height) 
VALUES `
	var validatorInfoParams []interface{}
	var validatorParams []interface{}

	for i, validator := range validators {
		vp := i * 2
		vi := i * 4 // Starting position for validator info params

		validatorQuery += fmt.Sprintf("($%d,$%d),", vp+1, vp+2)
		validatorParams = append(validatorParams,
			validator.ConsensusAddr, validator.ConsPubKey)

		validatorInfoQuery += fmt.Sprintf("($%d,$%d,$%d,$%d),", vi+1, vi+2, vi+3, vi+4)
		validatorInfoParams = append(validatorInfoParams,
			validator.ConsensusAddr, validator.OperatorAddr, validator.SelfDelegateAddress,
			validator.Height,
		)
	}

	validatorQuery = validatorQuery[:len(validatorQuery)-1] // Remove trailing ","
	validatorQuery += " ON CONFLICT DO NOTHING"
	_, err := db.Sql.Exec(validatorQuery, validatorParams...)
	if err != nil {
		return fmt.Errorf("error while storing validators: %s", err)
	}

	validatorInfoQuery = validatorInfoQuery[:len(validatorInfoQuery)-1] // Remove the trailing ","
	validatorInfoQuery += `
ON CONFLICT (consensus_address) DO UPDATE 
	SET consensus_address = excluded.consensus_address,
		operator_address = excluded.operator_address,
		self_delegate_address = excluded.self_delegate_address,
		height = excluded.height
WHERE validator_info.height <= excluded.height`
	_, err = db.Sql.Exec(validatorInfoQuery, validatorInfoParams...)
	if err != nil {
		return fmt.Errorf("error while storing validator infos: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveValidatorCommission saves validators commission in database.
func (db *Database) SaveValidatorCommission(validatorsCommission []types.ValidatorCommission) error {
	stmt := `INSERT INTO validator_commission (validator_address, commission, height) VALUES `

	var commissionList []interface{}
	for i, data := range validatorsCommission {
		si := i * 3
		stmt += fmt.Sprintf("($%d, $%d, $%d),", si+1, si+2, si+3)
		commissionList = append(commissionList,
			dbtypes.ToNullString(data.ValAddress),
			dbtypes.ToNullString(data.Commission),
			data.Height)
	}

	stmt = stmt[:len(stmt)-1]
	stmt += `
ON CONFLICT (validator_address) DO UPDATE 
	SET commission = excluded.commission, 
		height = excluded.height
WHERE validator_commission.height <= excluded.height`
	_, err := db.Sql.Exec(stmt, commissionList...)
	return err
}

// -------------------------------------------------------------------------------------------------------------------

// SaveValidatorDescription save validators description in database.
func (db *Database) SaveValidatorDescription(description []types.ValidatorDescription) error {
	stmt := `INSERT INTO validator_description (validator_address, moniker, identity, details, height) VALUES `

	var descriptionList []interface{}
	for i, desc := range description {
		si := i * 5

		stmt += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d),", si+1, si+2, si+3, si+4, si+5)
		descriptionList = append(descriptionList,
			dbtypes.ToNullString(desc.OperatorAddress),
			dbtypes.ToNullString(desc.Moniker),
			dbtypes.ToNullString(desc.Identity),
			dbtypes.ToNullString(desc.Description),
			desc.Height)
	}

	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT (validator_address) DO UPDATE
    SET moniker = excluded.moniker, 
        details = excluded.details,
		identity = excluded.identity,
        height = excluded.height
WHERE validator_description.height <= excluded.height`
	_, err := db.Sql.Exec(stmt, descriptionList...)
	return err

}

// -------------------------------------------------------------------------------------------------------------------

// SaveValidatorsVotingPower saves the given validator voting powers.
func (db *Database) SaveValidatorsVotingPower(entries []types.ValidatorVotingPower) error {
	if len(entries) == 0 {
		return nil
	}

	stmt := `INSERT INTO validator_voting_power (validator_address, voting_power, height) VALUES `
	var params []interface{}

	for i, entry := range entries {
		pi := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", pi+1, pi+2, pi+3)
		params = append(params, entry.ConsensusAddress, entry.VotingPower, entry.Height)
	}

	stmt = stmt[:len(stmt)-1]
	stmt += `
ON CONFLICT (validator_address) DO UPDATE 
	SET voting_power = excluded.voting_power, 
		height = excluded.height
WHERE validator_voting_power.height <= excluded.height`

	_, err := db.Sql.Exec(stmt, params...)
	if err != nil {
		return fmt.Errorf("error while storing validators voting power: %s", err)
	}

	return nil

}
