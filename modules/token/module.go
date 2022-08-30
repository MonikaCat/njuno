package token

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/forbole/njuno/database"
	"github.com/forbole/njuno/logging"
	"github.com/forbole/njuno/modules"
	source "github.com/forbole/njuno/node"
	"github.com/forbole/njuno/types/config"
)

var (
	_ modules.Module                     = &Module{}
	_ modules.AdditionalOperationsModule = &Module{}
)

// Module represents the token module
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

	tokenCfg, err := ParseConfig(bz)
	if err != nil {
		panic(err)
	}

	return &Module{
		cfg:    tokenCfg,
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
