package pruning

import (
	"github.com/MonikaCat/njuno/types/config"

	"github.com/MonikaCat/njuno/logging"

	"github.com/MonikaCat/njuno/database"
	"github.com/MonikaCat/njuno/modules"
)

var (
	_ modules.Module                     = &Module{}
	_ modules.BlockModule                = &Module{}
	_ modules.AdditionalOperationsModule = &Module{}
)

// Module represents the pruning module allowing to clean the database periodically
type Module struct {
	cfg    *Config
	db     database.Database
	logger logging.Logger
}

// NewModule builds a new Module instance
func NewModule(cfg config.Config, db database.Database, logger logging.Logger) *Module {
	bz, err := cfg.GetBytes()
	if err != nil {
		panic(err)
	}

	pruningCfg, err := ParseConfig(bz)
	if err != nil {
		panic(err)
	}

	return &Module{
		cfg:    pruningCfg,
		db:     db,
		logger: logger,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "pruning"
}

// RunAdditionalOperations implements
func (m *Module) RunAdditionalOperations() error {
	return RunAdditionalOperations(m.cfg)
}
