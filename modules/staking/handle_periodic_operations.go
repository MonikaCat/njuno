package staking

import (
	"fmt"

	staking "github.com/MonikaCat/njuno/modules/staking/utils"
	"github.com/MonikaCat/njuno/modules/utils"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RegisterPeriodicOperations implements modules.PeriodicOperationsModule
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "staking").Msg("setting up periodic tasks")

	// Fetch updated validators info every 5 mins
	if _, err := scheduler.Every(5).Minutes().Do(func() {
		utils.WatchMethod(m.updateValidatorsInfo)
	}); err != nil {
		return fmt.Errorf("error while setting up staking period operations: %s", err)
	}

	return nil
}

// updateValidatorsInfo allows to get the latest validators infos from yaml file
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

	return nil
}
