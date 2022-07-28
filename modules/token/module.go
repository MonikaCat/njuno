package token

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/MonikaCat/njuno/database"
	"github.com/MonikaCat/njuno/logging"
	"github.com/MonikaCat/njuno/modules"
	source "github.com/MonikaCat/njuno/node"
	"github.com/MonikaCat/njuno/types/config"
)

var (
	_ modules.Module                     = &Module{}
	_ modules.AdditionalOperationsModule = &Module{}
)

// Module represents the pricefeed module
type Module struct {
	cfg    *Config
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

	pricefeedCfg, err := ParseConfig(bz)
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
	return "token"
}
