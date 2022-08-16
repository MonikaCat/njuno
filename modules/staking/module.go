package staking

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/MonikaCat/njuno/database"
	"github.com/MonikaCat/njuno/logging"
	"github.com/MonikaCat/njuno/modules"
	source "github.com/MonikaCat/njuno/node"
	"github.com/MonikaCat/njuno/types"
)

var (
	_ modules.Module      = &Module{}
	_ modules.BlockModule = &Module{}
)

// Module represents the staking module
type Module struct {
	cdc            codec.Marshaler
	db             database.Database
	logger         logging.Logger
	source         source.Node
	validatorsList *types.ValidatorsList
}

func NewModule(cdc codec.Marshaler, db database.Database, logger logging.Logger, source source.Node, validatorsList *types.ValidatorsList) *Module {
	return &Module{
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
