package ibc

import (
	"fmt"

	"github.com/MonikaCat/njuno/modules/utils"
	"github.com/MonikaCat/njuno/types"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RegisterPeriodicOperations implements modules.PeriodicOperationsModule
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "ibc").Msg("setting up periodic tasks")

	// Setup a cron job to run every midnight
	if _, err := scheduler.Every(1).Day().At("00:00").Do(func() {
		utils.WatchMethod(m.updateIBCTransferParams)
	}); err != nil {
		return err
	}

	return nil
}

// updateIBCTransferParams gets the updated ibc transfer params
// and stores them inside the database
func (m *Module) updateIBCTransferParams() error {
	height, err := m.db.GetLastBlockHeight()
	if err != nil {
		return err
	}

	log.Debug().Str("module", "ibc").Int64("height", height).
		Msg("updating ibc transfer params")

	params, err := m.source.IBCTransferParams()
	if err != nil {
		return fmt.Errorf("error while getting ibc transfer params: %s", err)
	}

	return m.db.SaveIBCTransferParams(types.NewIBCTransferParams(params, height))

}
