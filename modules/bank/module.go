package bank

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/forbole/njuno/database"
	"github.com/forbole/njuno/logging"
	"github.com/forbole/njuno/modules"
	source "github.com/forbole/njuno/node"
)

var (
	_ modules.Module      = &Module{}
	_ modules.BlockModule = &Module{}
)

// Module represents the bank module
type Module struct {
	cdc    codec.Marshaler
	db     database.Database
	logger logging.Logger
	source source.Node
}

func NewModule(cdc codec.Marshaler, db database.Database, logger logging.Logger, source source.Node) *Module {
	return &Module{
		cdc:    cdc,
		db:     db,
		logger: logger,
		source: source,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "bank"
}
