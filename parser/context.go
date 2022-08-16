package parser

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/MonikaCat/njuno/logging"
	"github.com/MonikaCat/njuno/node"
	"github.com/MonikaCat/njuno/types"

	"github.com/MonikaCat/njuno/database"
	"github.com/MonikaCat/njuno/modules"
)

// Context represents the context that is shared among different workers
type Context struct {
	EncodingConfig *params.EncodingConfig
	Node           node.Node
	Database       database.Database
	Logger         logging.Logger
	Modules        []modules.Module
	ValidatorsList *types.ValidatorsList
}

// NewContext builds a new Context instance
func NewContext(
	encodingConfig *params.EncodingConfig, proxy node.Node, db database.Database,
	logger logging.Logger, modules []modules.Module, validatorsList *types.ValidatorsList,
) *Context {
	return &Context{
		EncodingConfig: encodingConfig,
		Node:           proxy,
		Database:       db,
		Modules:        modules,
		Logger:         logger,
		ValidatorsList: validatorsList,
	}
}
