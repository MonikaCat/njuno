package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TxResponse represents a valid transaction response
type TxResponse struct {
	Fee        TxFee          `json:"fee" yaml:"fee"`
	Memo       string         `json:"memo" yaml:"memo"`
	Msg        []TxMsg        `json:"msg" yaml:"msg"`
	Signatures []TxSignatures `json:"signatures" yaml:"signatures"`
	Hash       string         `json:"hash" yaml:"hash"`
	Height     int64          `json:"height" yaml:"height"`
}

type TxFee struct {
	Amount sdk.Coins `json:"amount" yaml:"amount"`
	Gas    string    `json:"gas" yaml:"gas"`
}

type TxSignatures struct {
	Signature string `json:"signature" yaml:"signature"`
}

type TxMsg struct {
	Type  string     `json:"type" yaml:"type"`
	Value TxMsgValue `json:"value" yaml:"value"`
}

type TxMsgValue struct {
	Amount           sdk.Coin `json:"amount" yaml:"amount"`
	DelegatorAddress string   `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress string   `json:"validator_address" yaml:"validator_address"`
}

// NewTxResponse allows to build a new TxResponse instance
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
