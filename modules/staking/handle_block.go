package staking

import (
	"encoding/hex"

	types "github.com/MonikaCat/njuno/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog/log"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// HandleBlock implements modules.BlockModule
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, _ *tmctypes.ResultBlockResults, vals *tmctypes.ResultValidators,
) error {

	// Update double sign evidences
	go m.updateDoubleSignEvidence(block.Block.Height, block.Block.Evidence.Evidence)

	// Update staking pool
	go m.updateStakingPool(block.Block.Height)

	// Update validators voting power
	go m.updateValidatorsVotingPower(vals.Validators, block.Block.Height)


	return nil
}

// updateDoubleSignEvidence updates the double sign evidence of all validators
func (m *Module) updateDoubleSignEvidence(height int64, evidenceList tmtypes.EvidenceList) {
	log.Debug().Str("module", "staking").Int64("height", height).
		Msg("updating double sign evidence")

	for _, ev := range evidenceList {
		dve, ok := ev.(*tmtypes.DuplicateVoteEvidence)
		if !ok {
			continue
		}

		evidence := types.NewDoubleSignEvidence(
			height,
			types.NewDoubleSignVote(
				int(dve.VoteA.Type),
				dve.VoteA.Height,
				dve.VoteA.Round,
				dve.VoteA.BlockID.String(),
				types.ConvertValidatorAddressToBech32String(dve.VoteA.ValidatorAddress),
				dve.VoteA.ValidatorIndex,
				hex.EncodeToString(dve.VoteA.Signature),
			),
			types.NewDoubleSignVote(
				int(dve.VoteB.Type),
				dve.VoteB.Height,
				dve.VoteB.Round,
				dve.VoteB.BlockID.String(),
				types.ConvertValidatorAddressToBech32String(dve.VoteB.ValidatorAddress),
				dve.VoteB.ValidatorIndex,
				hex.EncodeToString(dve.VoteB.Signature),
			),
		)

		err := m.db.SaveDoubleSignEvidence(evidence)
		if err != nil {
			log.Error().Str("module", "staking").Err(err).Int64("height", height).
				Msg("error while saving double sign evidence")
			return
		}

	}
}

// updateStakingPool reads the current staking pool and stores its value inside the database
func (m *Module) updateStakingPool(height int64) {
	log.Debug().Str("module", "staking").Int64("height", height).
		Msg("updating staking pool")

	stakingPool, err := m.source.StakingPool()
	if err != nil {
		log.Error().Str("module", "staking").Err(err).Int64("height", height).
			Msg("error while getting staking pool")
		return
	}

	err = m.db.SaveStakingPool(types.NewStakingPool(stakingPool.BondedTokens, stakingPool.NotBondedTokens, height))
	if err != nil {
		log.Error().Str("module", "staking").Err(err).Int64("height", height).
			Msg("error while saving staking pool")
		return
	}
}

// updateValidatorsVotingPower stores each validator latest voting power value inside the database
func (m *Module) updateValidatorsVotingPower(vals []*tmtypes.Validator, height int64) {
	log.Debug().Str("module", "staking").Int64("height", height).
		Msg("updating validators voting power")
	var validatorsVP []types.ValidatorVotingPower

	for _, val := range vals {
		consAddr := sdk.ConsAddress(val.Address).String()
		// store voting power
		validatorsVP = append(validatorsVP, types.NewValidatorVotingPower(consAddr, val.VotingPower, height))
	}

	err := m.db.SaveValidatorsVotingPower(validatorsVP)
	if err != nil {
		log.Error().Str("module", "staking").Err(err).Int64("height", height).
			Msg("error while saving validators voting power")
	}
}
