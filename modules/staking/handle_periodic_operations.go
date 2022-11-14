package staking

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	staking "github.com/forbole/njuno/modules/staking/utils"
	"github.com/forbole/njuno/modules/utils"
	types "github.com/forbole/njuno/types"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RegisterPeriodicOperations implements modules.PeriodicOperationsModule
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "staking").Msg("setting up periodic tasks")

	// Fetch updated validators and staking pool info every 5 mins
	if _, err := scheduler.Every(5).Minutes().Do(func() {
		utils.WatchMethod(m.updateValidatorsInfo)
	}); err != nil {
		return fmt.Errorf("error while setting up staking period operations: %s", err)
	}

	return nil
}

// updateValidatorsInfo allows to parse the latest validators infos from yaml file,
// update staking pool and store it inside the database
func (m *Module) updateValidatorsInfo() error {
	log.Debug().
		Str("module", "staking").
		Str("operation", "staking").
		Msg("updating validators info")

	height, err := m.db.GetLastBlockHeight()
	if err != nil {
		return fmt.Errorf("error while getting latest block height, error: %s", err)
	}

	// query the latest validators list
	validatorsLists := staking.GetLatestValidatorsList()

	// parse validators list
	validators, validatorsCommission, validatorsDescription, validatorsStatus, validatorsVP := staking.ParseValidatorsList(validatorsLists, height)

	err = m.db.SaveValidators(validators)
	if err != nil {
		log.Error().Str("module", "staking").Err(err).Int64("height", height).
			Msg("error while saving validators")
	}

	err = m.db.SaveValidatorCommission(validatorsCommission)
	if err != nil {
		log.Error().Str("module", "staking").Err(err).Int64("height", height).
			Msg("error while saving validators commission")
	}

	err = m.db.SaveValidatorDescription(validatorsDescription)
	if err != nil {
		log.Error().Str("module", "staking").Err(err).Int64("height", height).
			Msg("error while saving validators description")
	}

	err = m.db.SaveValidatorsStatus(validatorsStatus)
	if err != nil {
		log.Error().Str("module", "staking").Err(err).Int64("height", height).
			Msg("error while saving validators status")
	}

	err = m.db.SaveValidatorsVotingPower(validatorsVP)
	if err != nil {
		log.Error().Str("module", "staking").Err(err).Int64("height", height).
			Msg("error while saving validators voting power")
	}

	// update staking pool after updating validators VP
	// as bonded tokens correspond to overall voting power
	go m.updateStakingPool(height, validatorsVP)

	return nil
}

// updateStakingPool reads the current staking pool and stores its value inside the database
func (m *Module) updateStakingPool(height int64, validatorsVP []types.ValidatorVotingPower) {
	log.Debug().Str("module", "staking").Int64("height", height).
		Msg("updating staking pool")

	// Hard code total supply of 21M NOM
	var totalTokenSupply int64 = 21000000000000
	var bondedTokens int64

	// Calculate the overall voting power as total bonded tokens value
	for _, vp := range validatorsVP {
		v, err := strconv.ParseInt(vp.VotingPower, 10, 64)
		if err != nil {
			fmt.Errorf("failed to parse voting power from string to int: %s", err)
		}
		bondedTokens += v
	}

	// Calculate not bonded tokens by substracting bonded tokens
	// from the total token supply value
	var notBondedTokens = totalTokenSupply - bondedTokens

	err := m.db.SaveStakingPool(types.NewStakingPool(sdk.NewInt(bondedTokens), sdk.NewInt(notBondedTokens), height))
	if err != nil {
		log.Error().Str("module", "staking").Err(err).Int64("height", height).
			Msg("error while saving staking pool")
		return
	}
}
