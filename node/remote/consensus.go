package remote

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	constypes "github.com/tendermint/tendermint/consensus/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// Block implements node.Node
func (cp *Node) Block(height int64) (*tmctypes.ResultBlock, error) {
	return cp.client.Block(cp.ctx, &height)
}

// -------------------------------------------------------------------------------------------------------------------

// BlockResults implements node.Node
func (cp *Node) BlockResults(height int64) (*tmctypes.ResultBlockResults, error) {
	return cp.client.BlockResults(cp.ctx, &height)
}

// -------------------------------------------------------------------------------------------------------------------

// ConsensusState implements node.Node
func (cp *Node) ConsensusState() (*constypes.RoundStateSimple, error) {
	state, err := cp.client.ConsensusState(context.Background())
	if err != nil {
		return nil, err
	}

	var data constypes.RoundStateSimple
	err = tmjson.Unmarshal(state.RoundState, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// -------------------------------------------------------------------------------------------------------------------

// Genesis implements node.Node
func (cp *Node) Genesis() (*tmctypes.ResultGenesis, error) {
	res, err := cp.client.Genesis(cp.ctx)
	if err != nil && strings.Contains(err.Error(), "use the genesis_chunked API instead") {
		return cp.getGenesisChunked()
	}
	return res, err
}

// -------------------------------------------------------------------------------------------------------------------

// getGenesisChunked gets the genesis data using the chinked API instead
func (cp *Node) getGenesisChunked() (*tmctypes.ResultGenesis, error) {
	bz, err := cp.getGenesisChunksStartingFrom(0)
	if err != nil {
		return nil, err
	}

	var genDoc *tmtypes.GenesisDoc
	err = tmjson.Unmarshal(bz, &genDoc)
	if err != nil {
		return nil, err
	}

	return &tmctypes.ResultGenesis{Genesis: genDoc}, nil
}

// -------------------------------------------------------------------------------------------------------------------

// getGenesisChunksStartingFrom returns all the genesis chunks data starting from the chunk with the given id
func (cp *Node) getGenesisChunksStartingFrom(id uint) ([]byte, error) {
	res, err := cp.client.GenesisChunked(cp.ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error while getting genesis chunk %d out of %d", id, res.TotalChunks)
	}

	bz, err := base64.StdEncoding.DecodeString(res.Data)
	if err != nil {
		return nil, fmt.Errorf("error while decoding genesis chunk %d out of %d", id, res.TotalChunks)
	}

	if id == uint(res.TotalChunks-1) {
		return bz, nil
	}

	nextChunk, err := cp.getGenesisChunksStartingFrom(id + 1)
	if err != nil {
		return nil, err
	}

	return append(bz, nextChunk...), nil
}

// -------------------------------------------------------------------------------------------------------------------

// LatestHeight implements node.Node
func (cp *Node) LatestHeight() (int64, error) {
	status, err := cp.client.Status(cp.ctx)
	if err != nil {
		return -1, err
	}

	height := status.SyncInfo.LatestBlockHeight
	return height, nil
}
