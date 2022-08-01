package types

// IBCTransactionParams represents the x/ibc transaction parameters
type IBCTransactionParams struct {
	ReceiveEnabled bool `json:"receive_enabled" yaml:"receive_enabled"`
	SendEnabled    bool `json:"send_enabled" yaml:"send_enabled"`
}

// IBCTransferParams represents the x/ibc trasfer parameters
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
