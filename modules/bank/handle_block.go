package bank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog/log"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// HandleBlock implements modules.BlockModule
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, _ *tmctypes.ResultBlockResults, _ *tmctypes.ResultValidators,
) error {
	err := m.updateSupply(block.Block.Height)
	if err != nil {
		log.Error().Str("module", "bank").Int64("height", block.Block.Height).
			Err(err).Msg("error while updating supply")
	}

	return nil
}

// updateSupply updates the supply for a given height
func (m *Module) updateSupply(height int64) error {
	log.Debug().Str("module", "bank").Int64("height", height).
		Msg("updating supply")

	supply, err := m.source.Supply()
	if err != nil {
		return err
	}

	for _, index := range supply {
		if index.Denom == "unom" {
			// Hard code total supply of 21M NOM if the total supply
			// returned from node equals to 1 or less
			if index.Amount.LTE(sdk.NewInt(1)) {
				totalSupply := sdk.NewCoin(index.Denom, sdk.NewInt(21000000000000))
				return m.db.SaveSupply(sdk.NewCoins(totalSupply), height)

			}
		}
	}

	return m.db.SaveSupply(supply, height)
}
