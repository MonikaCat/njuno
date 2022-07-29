package types

// Message represents the data of a single message
type Message struct {
	TxHash    string
	Index     int
	Type      string
	Value     string
	Addresses []string
	Height    int64
}

// NewMessage allows to build a new Message instance
func NewMessage(txHash string, index int, msgType string, value string, addresses []string, height int64) *Message {
	return &Message{
		TxHash:    txHash,
		Index:     index,
		Type:      msgType,
		Value:     value,
		Addresses: addresses,
		Height:    height,
	}
}
