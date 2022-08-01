package remote

import (
	"encoding/json"
	"fmt"

	bdtypes "github.com/MonikaCat/njuno/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// Txs implements node.Node
func (cp *Node) Txs(block *tmctypes.ResultBlock) ([]bdtypes.TxResponse, error) {
	txResponses := make([]bdtypes.TxResponse, len(block.Block.Txs))

	// get tx details from the block
	var transaction bdtypes.TxResponse
	for _, t := range block.Block.Txs {
		err := json.Unmarshal(t, &transaction)
		if err != nil {
			// continue
		}
		txResponses = append(txResponses, bdtypes.NewTxResponse(transaction.Fee, transaction.Memo, transaction.Msg, transaction.Signatures, fmt.Sprintf("%X", t.Hash()), block.Block.Height))
	}

	return txResponses, nil
}
