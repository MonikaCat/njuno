package types

import "database/sql"

// ValidatorCommissionRow represents a single row of the validator_commission database table
type ValidatorCommissionRow struct {
	OperatorAddress   string         `db:"validator_address"`
	Commission        string         `db:"commission"`
	MinSelfDelegation sql.NullString `db:"min_self_delegation"`
	Height            int64          `db:"height"`
}

// NewValidatorCommissionRow allows to build new ValidatorCommissionRow instance
func NewValidatorCommissionRow(
	operatorAddress string, commission string, minSelfDelegation string, height int64,
) ValidatorCommissionRow {
	return ValidatorCommissionRow{
		OperatorAddress:   operatorAddress,
		Commission:        commission,
		MinSelfDelegation: ToNullString(minSelfDelegation),
		Height:            height,
	}
}

// Equal tells whether v and w represent the same rows
func (v ValidatorCommissionRow) Equal(w ValidatorCommissionRow) bool {
	return v.OperatorAddress == w.OperatorAddress &&
		v.Commission == w.Commission &&
		v.MinSelfDelegation == w.MinSelfDelegation &&
		v.Height == w.Height
}

// _________________________________________________________

// ValidatorDescriptionRow represent a single row in validator_description database table.
type ValidatorDescriptionRow struct {
	ValAddress string         `db:"validator_address"`
	Moniker    sql.NullString `db:"moniker"`
	Identity   sql.NullString `db:"identity"`
	Details    sql.NullString `db:"details"`
	Height     int64          `db:"height"`
}

// NewValidatorDescriptionRow allows to build new ValidatorDescriptionRow instance
func NewValidatorDescriptionRow(
	valAddress, moniker, identity, details string, height int64,
) ValidatorDescriptionRow {
	return ValidatorDescriptionRow{
		ValAddress: valAddress,
		Moniker:    ToNullString(moniker),
		Identity:   ToNullString(identity),
		Details:    ToNullString(details),
		Height:     height,
	}
}

// Equal tells whether v and w represent the same rows
func (v ValidatorDescriptionRow) Equal(w ValidatorDescriptionRow) bool {
	return v.ValAddress == w.ValAddress &&
		v.Moniker == w.Moniker &&
		v.Identity == w.Identity &&
		v.Details == w.Details &&
		v.Height == w.Height
}

// _________________________________________________________

// ValidatorStatusRow represents a single row of the validator_status table
type ValidatorStatusRow struct {
	InActiveSet string `db:"in_active_set"`
	Jailed      string `db:"jailed"`
	ConsAddress string `db:"validator_address"`
	Height      int64  `db:"height"`
}

// NewValidatorStatusRow builds a new ValidatorStatusRow
func NewValidatorStatusRow(inActiveSet, jailed, consAddess string, height int64) ValidatorStatusRow {
	return ValidatorStatusRow{
		InActiveSet: inActiveSet,
		Jailed:      jailed,
		ConsAddress: consAddess,
		Height:      height,
	}
}

// Equal tells whether v and w contain the same data
func (v ValidatorStatusRow) Equal(w ValidatorStatusRow) bool {
	return v.InActiveSet == w.InActiveSet &&
		v.Jailed == w.Jailed &&
		v.ConsAddress == w.ConsAddress &&
		v.Height == w.Height
}
