package types

import sdk "github.com/cosmos/cosmos-sdk/types"

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
	SelfDelegateAddress string
	Height              int64
}

// NewValidator allows to build a new Validator instance
func NewValidator(
	consAddr string,
	selfDelegateAddress string,
	height int64,
) Validator {
	return Validator{
		ConsensusAddr:       consAddr,
		SelfDelegateAddress: selfDelegateAddress,
		Height:              height,
	}
}

// ----------------------------------------------------------------------------------------------------------

// ValidatorCommission contains the data of a validator commission at a given height
type ValidatorCommission struct {
	ValAddress          string
	SelfDelegateAddress string
	Commission          string
	MinSelfDelegation   string
	Height              int64
}

// NewValidatorCommission return a new ValidatorCommission instance
func NewValidatorCommission(
	valAddress, selfDelegateAddress, commission, minSelfDelegation string, height int64,
) ValidatorCommission {
	return ValidatorCommission{
		ValAddress:          valAddress,
		SelfDelegateAddress: selfDelegateAddress,
		Commission:          commission,
		MinSelfDelegation:   minSelfDelegation,
		Height:              height,
	}
}

// ----------------------------------------------------------------------------------------------------------

// ValidatorDescription contains the description of a validator
// and timestamp do the description get changed
type ValidatorDescription struct {
	OperatorAddress     string
	SelfDelegateAddress string
	Description         string
	Identity            string
	Moniker             string
	AvatarURL           string
	Height              int64
}

// NewValidatorDescription returns a new ValidatorDescription object
func NewValidatorDescription(
	opAddr, selfDelegateAddress, description, identity, avatarURL, moniker string, height int64,
) ValidatorDescription {
	return ValidatorDescription{
		OperatorAddress:     opAddr,
		SelfDelegateAddress: selfDelegateAddress,
		Description:         description,
		Identity:            identity,
		AvatarURL:           avatarURL,
		Moniker:             moniker,
		Height:              height,
	}
}

// ----------------------------------------------------------------------------------------------------------

// ValidatorsList represents validators list from a file
type ValidatorsList struct {
	Validators []ValidatorList `yaml:"validators"`
}

type ValidatorList struct {
	Validator ValidatorInfo `yaml:"validator"`
}
type ValidatorInfo struct {
	Address           string `yaml:"address"`
	Commission        string `yaml:"commission"`
	Details           string `yaml:"details"`
	Identity          string `yaml:"identity"`
	Jailed            string `yaml:"jailed"`
	MinSelfDelegation string `yaml:"min_self_delegation"`
	Moniker           string `yaml:"moniker"`
	Tombstoned        string `yaml:"tombstoned"`
	InActiveSet       string `yaml:"in_active_set"`
	VotingPower       string `yaml:"voting_power"`
}

// ----------------------------------------------------------------------------------------------------------

// ValidatorVotingPower represents the voting power of a validator at a specific block height
type ValidatorVotingPower struct {
	ConsensusAddress    string
	SelfDelegateAddress string
	VotingPower         string
	Height              int64
}

// NewValidatorVotingPower creates a new ValidatorVotingPower
func NewValidatorVotingPower(address, selfDelegateAddress string, votingPower string, height int64) ValidatorVotingPower {
	return ValidatorVotingPower{
		ConsensusAddress:    address,
		VotingPower:         votingPower,
		SelfDelegateAddress: selfDelegateAddress,
		Height:              height,
	}
}

// ----------------------------------------------------------------------------------------------------------

// ValidatorStatus represents the latest state of a validator
type ValidatorStatus struct {
	ConsensusAddress    string
	SelfDelegateAddress string
	InActiveSet         string
	Jailed              string
	Tombstoned          string
	Height              int64
}

// NewValidatorStatus creates a new ValidatorVotingPower
func NewValidatorStatus(valConsAddr, selfDelegateAddress, inActiveSet, jailed, tombstoned string, height int64) ValidatorStatus {
	return ValidatorStatus{
		ConsensusAddress:    valConsAddr,
		SelfDelegateAddress: selfDelegateAddress,
		InActiveSet:         inActiveSet,
		Jailed:              jailed,
		Tombstoned:          tombstoned,
		Height:              height,
	}
}

// DelegationResponses represents the total delegation value
// of each address
type DelegationResponses struct {
	Balance sdk.Coin `json:"balance"`
}

// QueryAllBalancesResponse contains the account total delegation value
type QueryTotalDelegationsResponse struct {
	DelegationResponses []DelegationResponses `json:"delegation_responses"`
}
