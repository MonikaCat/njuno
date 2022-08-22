package registrar

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/MonikaCat/njuno/node"
	"github.com/MonikaCat/njuno/types"

	"github.com/MonikaCat/njuno/modules/actions"
	"github.com/MonikaCat/njuno/modules/bank"
	"github.com/MonikaCat/njuno/modules/consensus"
	"github.com/MonikaCat/njuno/modules/ibc"
	"github.com/MonikaCat/njuno/modules/mint"
	"github.com/MonikaCat/njuno/modules/pricefeed"
	"github.com/MonikaCat/njuno/modules/staking"
	"github.com/MonikaCat/njuno/modules/telemetry"
	"github.com/MonikaCat/njuno/modules/token"

	"github.com/MonikaCat/njuno/logging"

	"github.com/MonikaCat/njuno/types/config"

	"github.com/MonikaCat/njuno/modules/pruning"

	"github.com/MonikaCat/njuno/modules"
	"github.com/MonikaCat/njuno/modules/messages"

	"github.com/MonikaCat/njuno/database"
)

// Context represents the context of the modules registrar
type Context struct {
	NJunoConfig    config.Config
	SDKConfig      *sdk.Config
	EncodingConfig *params.EncodingConfig
	Database       database.Database
	Proxy          node.Node
	Logger         logging.Logger
	ValidatorsList *types.ValidatorsList
}

// NewContext allows to build a new Context instance
func NewContext(
	parsingConfig config.Config, sdkConfig *sdk.Config, encodingConfig *params.EncodingConfig,
	database database.Database, proxy node.Node, logger logging.Logger, validatorsList *types.ValidatorsList,
) Context {
	return Context{
		NJunoConfig:    parsingConfig,
		SDKConfig:      sdkConfig,
		EncodingConfig: encodingConfig,
		Database:       database,
		Proxy:          proxy,
		Logger:         logger,
		ValidatorsList: validatorsList,
	}
}

// Registrar represents a modules registrar. This allows to build a list of modules that can later be used by
// specifying their names inside the YAML configuration file.
type Registrar interface {
	BuildModules(context Context) modules.Modules
}

// ------------------------------------------------------------------------------------------------------------------

var (
	_ Registrar = &EmptyRegistrar{}
)

// EmptyRegistrar represents a Registrar which does not register any custom module
type EmptyRegistrar struct{}

// BuildModules implements Registrar
func (*EmptyRegistrar) BuildModules(_ Context) modules.Modules {
	return nil
}

// ------------------------------------------------------------------------------------------------------------------

var (
	_ Registrar = &DefaultRegistrar{}
)

// DefaultRegistrar represents a registrar that allows to handle the default nJuno modules
type DefaultRegistrar struct {
	parser messages.MessageAddressesParser
}

// NewDefaultRegistrar builds a new DefaultRegistrar
func NewDefaultRegistrar(parser messages.MessageAddressesParser) *DefaultRegistrar {
	return &DefaultRegistrar{
		parser: parser,
	}
}

// BuildModules implements Registrar
func (r *DefaultRegistrar) BuildModules(ctx Context) modules.Modules {
	return modules.Modules{
		actions.NewModule(ctx.NJunoConfig, ctx.EncodingConfig),
		bank.NewModule(ctx.EncodingConfig.Marshaler, ctx.Database, ctx.Logger, ctx.Proxy),
		consensus.NewModule(ctx.Database),
		ibc.NewModule(ctx.EncodingConfig.Marshaler, ctx.Database, ctx.Logger, ctx.Proxy),
		mint.NewModule(ctx.EncodingConfig.Marshaler, ctx.Database, ctx.Logger, ctx.Proxy),
		pricefeed.NewModule(ctx.NJunoConfig, ctx.EncodingConfig.Marshaler, ctx.Database, ctx.Logger, ctx.Proxy),
		pruning.NewModule(ctx.NJunoConfig, ctx.Database, ctx.Logger),
		staking.NewModule(ctx.NJunoConfig, ctx.EncodingConfig.Marshaler, ctx.Database, ctx.Logger, ctx.Proxy, ctx.ValidatorsList),
		telemetry.NewModule(ctx.NJunoConfig),
		token.NewModule(ctx.NJunoConfig, ctx.EncodingConfig.Marshaler, ctx.Database, ctx.Logger, ctx.Proxy),
	}
}

// ------------------------------------------------------------------------------------------------------------------

// GetModules returns the list of module implementations based on the given module names.
// For each module name that is specified but not found, a warning log is printed.
func GetModules(mods modules.Modules, names []string, logger logging.Logger) []modules.Module {
	var modulesImpls []modules.Module
	for _, name := range names {
		module, found := mods.FindByName(name)
		if found {
			modulesImpls = append(modulesImpls, module)
		} else {
			logger.Error("Module is required but not registered. Be sure to register it using registrar.RegisterModule", "module", name)
		}
	}
	return modulesImpls
}
