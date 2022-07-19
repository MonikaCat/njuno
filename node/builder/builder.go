package builder

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/MonikaCat/njuno/node"
	nodeconfig "github.com/MonikaCat/njuno/node/config"
	"github.com/MonikaCat/njuno/node/remote"
)

func BuildNode(cfg nodeconfig.Config, encodingConfig *params.EncodingConfig) (node.Node, error) {
	switch cfg.Type {
	case nodeconfig.TypeRemote:
		return remote.NewNode(cfg.Details.(*remote.Details), encodingConfig.Marshaler)
	case nodeconfig.TypeNone:
		return nil, nil

	default:
		return nil, fmt.Errorf("invalid node type: %s", cfg.Type)
	}
}
