package staking

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/MonikaCat/njuno/database"
	"github.com/MonikaCat/njuno/logging"
	"github.com/MonikaCat/njuno/modules"
	source "github.com/MonikaCat/njuno/node"
	"github.com/MonikaCat/njuno/types"
	"github.com/MonikaCat/njuno/types/config"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.BlockModule              = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

// Module represents the staking module
type Module struct {
	cfg            config.Config
	cdc            codec.Marshaler
	db             database.Database
	logger         logging.Logger
	source         source.Node
	validatorsList *types.ValidatorsList
}

func NewModule(cfg config.Config, cdc codec.Marshaler, db database.Database, logger logging.Logger, source source.Node, validatorsList *types.ValidatorsList) *Module {
	return &Module{
		cfg:            cfg,
		cdc:            cdc,
		db:             db,
		logger:         logger,
		source:         source,
		validatorsList: validatorsList,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "staking"
}
