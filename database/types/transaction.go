package types

import "github.com/forbole/njuno/types"

// NewDBSignatures returns signatures in string array
func NewDBSignatures(signaturesList []types.TxSignatures) []string {
	var signatures []string
	for _, index := range signaturesList {
		signatures = append(signatures, index.Signature)
	}
	return signatures
}
