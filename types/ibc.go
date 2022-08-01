package types

// IBCTransfer represents the x/ibc transfer parameters
type IBCTransfer struct {
	ReceiveEnabled bool `json:"receive_enabled" yaml:"receive_enabled"`
	SendEnabled    bool `json:"send_enabled" yaml:"send_enabled"`
}

// IBCTransferParams represents the x/ibc transfer parameters
type IBCTransferParams struct {
	Params IBCTransfer
	Height int64
}

// NewIBCTransferParams allows to build a new IBCTransferParams instance
func NewIBCTransferParams(params IBCTransfer, height int64) *IBCTransferParams {
	return &IBCTransferParams{
		Params: params,
		Height: height,
	}
}
