package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

// Tx represents an already existing blockchain transaction
type Tx struct {
	*tx.Tx
	*sdk.TxResponse
}

// ----------------------------------------------------------------------------------------------------------

// TxResponse represents a valid transaction response
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
