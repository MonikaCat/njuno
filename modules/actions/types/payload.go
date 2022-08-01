package types

// Payload contains the payload data that is sent from Hasura
type Payload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            PayloadArgs            `json:"input"`
}
type PayloadArgs struct {
	Address    string `json:"address"`
	Height     int64  `json:"height"`
	Offset     uint64 `json:"offset"`
	Limit      uint64 `json:"limit"`
	CountTotal bool   `json:"count_total"`
}

// GetAddress returns the address associated with this payload, if any
func (p *Payload) GetAddress() string {
	return p.Input.Address
}
