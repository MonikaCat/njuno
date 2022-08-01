package types

// DoubleSignEvidence represent a double sign evidence on each tendermint block
type DoubleSignEvidence struct {
	VoteA  DoubleSignVote
	VoteB  DoubleSignVote
	Height int64
}

// NewDoubleSignEvidence return a new DoubleSignEvidence object
func NewDoubleSignEvidence(height int64, voteA DoubleSignVote, voteB DoubleSignVote) DoubleSignEvidence {
	return DoubleSignEvidence{
		VoteA:  voteA,
		VoteB:  voteB,
		Height: height,
	}
}

// ----------------------------------------------------------------------------------------------------------

// DoubleSignVote represents a double vote which is included inside a DoubleSignEvidence
type DoubleSignVote struct {
	BlockID          string
	ValidatorAddress string
	Signature        string
	Type             int
	Height           int64
	Round            int32
	ValidatorIndex   int32
}

// NewDoubleSignVote allows to create a new DoubleSignVote instance
func NewDoubleSignVote(
	roundType int,
	height int64,
	round int32,
	blockID string,
	validatorAddress string,
	validatorIndex int32,
	signature string,
) DoubleSignVote {
	return DoubleSignVote{
		Type:             roundType,
		Height:           height,
		Round:            round,
		BlockID:          blockID,
		ValidatorAddress: validatorAddress,
		ValidatorIndex:   validatorIndex,
		Signature:        signature,
	}
}

// ----------------------------------------------------------------------------------------------------------

// Validator contains the data of a single validator
type Validator struct {
	ConsensusAddr       string
	ConsPubKey          string
	OperatorAddr        string
	SelfDelegateAddress string
	Height              int64
}

// NewValidator allows to build a new Validator instance
func NewValidator(
	consAddr string, opAddr string, consPubKey string,
	selfDelegateAddress string,
	height int64,
) Validator {
	return Validator{
		ConsensusAddr:       consAddr,
		ConsPubKey:          consPubKey,
		OperatorAddr:        opAddr,
		SelfDelegateAddress: selfDelegateAddress,
		Height:              height,
	}
}

// ----------------------------------------------------------------------------------------------------------

// ValidatorCommission contains the data of a validator commission at a given height
type ValidatorCommission struct {
	ValAddress string
	Commission string
	Height     int64
}

// NewValidatorCommission return a new ValidatorCommission instance
func NewValidatorCommission(
	valAddress string, rate string, height int64,
) ValidatorCommission {
	return ValidatorCommission{
		ValAddress: valAddress,
		Commission: rate,
		Height:     height,
	}
}

// ----------------------------------------------------------------------------------------------------------

// ValidatorDescription contains the description of a validator
// and timestamp do the description get changed
type ValidatorDescription struct {
	OperatorAddress string
	Description     string
	Identity        string
	Moniker         string
	Height          int64
}

// NewValidatorDescription returns a new ValidatorDescription object
func NewValidatorDescription(
	opAddr string, description string, identity string, moniker string, height int64,
) ValidatorDescription {
	return ValidatorDescription{
		OperatorAddress: opAddr,
		Description:     description,
		Identity:        identity,
		Moniker:         moniker,
		Height:          height,
	}
}

// ----------------------------------------------------------------------------------------------------------

// ValidatorsList represents validators list
type ValidatorsList struct {
	Validators []ValidatorList `yaml:"validators"`
}

type ValidatorList struct {
	Validator ValidatorInfo `yaml:"validator"`
}
type ValidatorInfo struct {
	Hex        string `yaml:"hex"`
	Address    string `yaml:"address"`
	Commission string `yaml:"commission"`
	Details    string `yaml:"details"`
	Identity   string `yaml:"identity"`
	Moniker    string `yaml:"moniker"`
}

// ----------------------------------------------------------------------------------------------------------

// ValidatorVotingPower represents the voting power of a validator at a specific block height
type ValidatorVotingPower struct {
	ConsensusAddress string
	VotingPower      int64
	Height           int64
}

// NewValidatorVotingPower creates a new ValidatorVotingPower
func NewValidatorVotingPower(address string, votingPower int64, height int64) ValidatorVotingPower {
	return ValidatorVotingPower{
		ConsensusAddress: address,
		VotingPower:      votingPower,
		Height:           height,
	}
}
