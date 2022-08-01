package types

// IBCTransferParams represents the x/ibc transfer parameters
type IBCTransferParams struct {
	ReceiveEnabled bool `json:"receive_enabled" yaml:"receive_enabled"`
	SendEnabled    bool `json:"send_enabled" yaml:"send_enabled"`
}

// IBCParams represents the x/ibc transfer parameters
type IBCParams struct {
	Params IBCTransferParams
	Height int64
}

// NewIBCParams allows to build a new IBCParams instance
func NewIBCParams(params IBCTransferParams, height int64) *IBCParams {
	return &IBCParams{
		Params: params,
		Height: height,
	}
}
