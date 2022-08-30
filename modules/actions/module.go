package actions

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/forbole/njuno/modules"
	"github.com/forbole/njuno/node"
	"github.com/forbole/njuno/node/builder"
	nodeconfig "github.com/forbole/njuno/node/config"
	"github.com/forbole/njuno/types/config"
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
	nJunoNode, err := builder.BuildNode(nodeCfg, encodingConfig)
	if err != nil {
		panic(err)
	}

	return &Module{
		cfg:  actionsCfg,
		node: nJunoNode,
	}
}

func (m *Module) Name() string {
	return "actions"
}
