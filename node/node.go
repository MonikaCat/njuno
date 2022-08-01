package node

import (
	types "github.com/MonikaCat/njuno/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	constypes "github.com/tendermint/tendermint/consensus/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

type Node interface {

	// AccountBalance queries for the balance of given address.
	// An error is returned if the query fails.
	AccountBalance(address string) (sdk.Coins, error)

	// Block queries for a block by height.
	// An error is returned if the query fails.
	Block(height int64) (*tmctypes.ResultBlock, error)

	// BlockResults queries the results of a block by height.
	// An error is returnes if the query fails
	BlockResults(height int64) (*tmctypes.ResultBlockResults, error)

	// ConsensusState queries for the latest consensus state of the chain.
	// An error is returned if the query fails.
	ConsensusState() (*constypes.RoundStateSimple, error)

	// Genesis returns the genesis state.
	// An error is returned if the query fails.
	Genesis() (*tmctypes.ResultGenesis, error)

	// IBCTransferParams queries the latest ibc parameters.
	// An error is returned if the query fails.
	IBCTransferParams() (types.IBCTransfer, error)

	// Inflation queries the latest inflation value.
	// An error is returned if the query fails.
	Inflation() (string, error)

	// LatestHeight returns the latest block height on the active chain.
	// An error is returned if the query fails.
	LatestHeight() (int64, error)

	// StakingPool queries the latest staking pool value.
	// An error is returned if the query fails.
	StakingPool() (stakingtypes.Pool, error)

	// Stop defers the node stop execution to the client.
	Stop()

	// Supply queries the latest supply value.
	// An error is returned if the query fails.
	Supply() (sdk.Coins, error)

	// Validators returns all the known Tendermint validators for a given block
	// height. An error is returned if the query fails.
	Validators(height int64) (*tmctypes.ResultValidators, error)
}
