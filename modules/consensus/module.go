package consensus

import (
	"github.com/MonikaCat/njuno/database"

	"github.com/MonikaCat/njuno/modules"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
	_ modules.GenesisModule            = &Module{}
	_ modules.BlockModule              = &Module{}
)

// Module implements the consensus utils
type Module struct {
	db database.Database
}

// NewModule builds a new Module instance
func NewModule(db database.Database) *Module {
	return &Module{
		db: db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "consensus"
}
