package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TxResponse struct {
	Fee        TxFee          `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee,omitempty"`
	Memo       string         `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
	Msg        []TxMsg        `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Signatures []TxSignatures `protobuf:"bytes,4,opt,name=signatures,proto3" json:"signatures,omitempty"`
	Hash       string         `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	Height     int64          `protobuf:"bytes,6,opt,name=height,proto3" json:"height,omitempty"`
}

type TxFee struct {
	Amount sdk.Coins `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Gas    string    `protobuf:"bytes,2,opt,name=gas,proto3" json:"gas"`
}

type TxSignatures struct {
	Signature string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature"`
}

type TxMsg struct {
	Type  string     `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value TxMsgValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
}

type TxMsgValue struct {
	Amount           sdk.Coin `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	DelegatorAddress string   `protobuf:"bytes,2,opt,name=delegator_address,proto3" json:"delegator_address"`
	ValidatorAddress string   `protobuf:"bytes,3,opt,name=validator_address,proto3" json:"validator_address"`
}

func NewTxResponse(
	fee TxFee, memo string, msg []TxMsg, sig []TxSignatures, hash string, height int64,
) TxResponse {
	return TxResponse{
		Fee:        fee,
		Memo:       memo,
		Msg:        msg,
		Signatures: sig,
		Hash:       hash,
		Height:     height,
	}
}

type InflationResponse struct {
	Inflation string `json:"inflation" yaml:"inflation"`
}

// StakingPool contains the data of the staking pool at the given height
type StakingPool struct {
	BondedTokens    sdk.Int
	NotBondedTokens sdk.Int
	Height          int64
}

// NewStakingPool allows to build a new StakingPool instance
func NewStakingPool(bondedTokens sdk.Int, notBondedTokens sdk.Int, height int64) *StakingPool {
	return &StakingPool{
		BondedTokens:    bondedTokens,
		NotBondedTokens: notBondedTokens,
		Height:          height,
	}
}

type IBCTransactionParams struct {
	ReceiveEnabled bool `json:"receive_enabled" yaml:"receive_enabled"`
	SendEnabled    bool `json:"send_enabled" yaml:"send_enabled"`
}

type IBCTransferParams struct {
	Params IBCTransactionParams `json:"params" yaml:"params"`
}

// IBCParams represents the x/ibc parameters
type IBCParams struct {
	Params IBCTransactionParams
	Height int64
}

// NewIBCParams allows to build a new IBCParams instance
func NewIBCParams(params IBCTransactionParams, height int64) *IBCParams {
	return &IBCParams{
		Params: params,
		Height: height,
	}
}

// Token represents a valid token inside the chain
type Token struct {
	Name  string      `yaml:"name"`
	Units []TokenUnit `yaml:"units"`
}

func NewToken(name string, units []TokenUnit) Token {
	return Token{
		Name:  name,
		Units: units,
	}
}

// TokenUnit represents a unit of a token
type TokenUnit struct {
	Denom    string   `yaml:"denom"`
	Exponent int      `yaml:"exponent"`
	Aliases  []string `yaml:"aliases,omitempty"`
	PriceID  string   `yaml:"price_id,omitempty"`
}

func NewTokenUnit(denom string, exponent int, aliases []string, priceID string) TokenUnit {
	return TokenUnit{
		Denom:    denom,
		Exponent: exponent,
		Aliases:  aliases,
		PriceID:  priceID,
	}
}

// Genesis contains the useful information about the genesis
type Genesis struct {
	ChainID       string
	Time          time.Time
	InitialHeight int64
}

// NewGenesis allows to build a new Genesis instance
func NewGenesis(chainID string, startTime time.Time, initialHeight int64) *Genesis {
	return &Genesis{
		ChainID:       chainID,
		Time:          startTime,
		InitialHeight: initialHeight,
	}
}

// ValidatorsList represents validators list
type ValidatorsList struct {
	bytes      []byte
	Validators []ValidatorList `yaml:"validators"`
}

type ValidatorList struct {
	Validator ValidatorInfo `yaml:"validator"`
}
type ValidatorInfo struct {
	Hex     string `yaml:"hex"`
	Address string `yaml:"address"`
	Moniker string `yaml:"moniker"`
	Details string `yaml:"details"`
}

// ValidatorDescription contains the description of a validator
// and timestamp do the description get changed
type ValidatorDescription struct {
	OperatorAddress string
	Description     string
	Moniker         string
	Height          int64
}

// NewValidatorDescription returns a new ValidatorDescription object
func NewValidatorDescription(
	opAddr string, description string, moniker string, height int64,
) ValidatorDescription {
	return ValidatorDescription{
		OperatorAddress: opAddr,
		Description:     description,
		Moniker:         moniker,
		Height:          height,
	}
}

type QueryAllBalancesResponse struct {
	Balances   sdk.Coins     `son:"balances"`
	Pagination *PageResponse `json:"pagination,omitempty"`
}

type PageResponse struct {
	NextKey []byte `json:"next_key,omitempty"`
	Total   string `json:"total,omitempty"`
}

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
