package types

import "github.com/MonikaCat/njuno/types"

// NewDBSignatures returns signatures in string array
func NewDBSignatures(signaturesList []types.TxSignatures) []string {
	var signatures []string
	for _, index := range signaturesList {
		signatures = append(signatures, index.Signature)
	}
	return signatures
}
