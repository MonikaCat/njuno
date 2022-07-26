package actions

import (
	"github.com/MonikaCat/njuno/modules"
	"github.com/MonikaCat/njuno/node"
	"github.com/MonikaCat/njuno/node/builder"
	nodeconfig "github.com/MonikaCat/njuno/node/config"
	"github.com/MonikaCat/njuno/types/config"
	"github.com/cosmos/cosmos-sdk/simapp/params"
)

var (
	_ modules.Module                     = &Module{}
	_ modules.AdditionalOperationsModule = &Module{}
)

type Module struct {
	cfg  *Config
	node node.Node
}

func NewModule(cfg config.Config, encodingConfig *params.EncodingConfig) *Module {
	bz, err := cfg.GetBytes()
	if err != nil {
		panic(err)
	}

	actionsCfg, err := ParseConfig(bz)
	if err != nil {
		panic(err)
	}

	nodeCfg := cfg.Node
	if actionsCfg.Node != nil {
		nodeCfg = nodeconfig.NewConfig(nodeconfig.TypeRemote, actionsCfg.Node)
	}

	// Build the node
	junoNode, err := builder.BuildNode(nodeCfg, encodingConfig)
	if err != nil {
		panic(err)
	}

	return &Module{
		cfg:  actionsCfg,
		node: junoNode,
	}
}

func (m *Module) Name() string {
	return "actions"
}
