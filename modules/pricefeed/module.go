package pricefeed

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/forbole/njuno/database"
	"github.com/forbole/njuno/logging"
	"github.com/forbole/njuno/modules"
	"github.com/forbole/njuno/modules/token"
	source "github.com/forbole/njuno/node"
	"github.com/forbole/njuno/types/config"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

// Module represents the pricefeed module
type Module struct {
	cfg    *token.Config
	cdc    codec.Marshaler
	db     database.Database
	logger logging.Logger
	source source.Node
}

func NewModule(cfg config.Config, cdc codec.Marshaler, db database.Database, logger logging.Logger, source source.Node) *Module {
	bz, err := cfg.GetBytes()
	if err != nil {
		panic(err)
	}

	pricefeedCfg, err := token.ParseConfig(bz)
	if err != nil {
		panic(err)
	}

	return &Module{
		cfg:    pricefeedCfg,
		cdc:    cdc,
		db:     db,
		logger: logger,
		source: source,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "pricefeed"
}
