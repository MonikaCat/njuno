package v3

import (
	databaseconfig "github.com/MonikaCat/njuno/database/config"
	loggingconfig "github.com/MonikaCat/njuno/logging/config"
	"github.com/MonikaCat/njuno/modules/pruning"
	"github.com/MonikaCat/njuno/modules/telemetry"
	nodeconfig "github.com/MonikaCat/njuno/node/config"
	parserconfig "github.com/MonikaCat/njuno/parser/config"
	pricefeedconfig "github.com/MonikaCat/njuno/pricefeed"
	"github.com/MonikaCat/njuno/types/config"
)

// Config defines all necessary juno configuration parameters.
type Config struct {
	Chain    config.ChainConfig    `yaml:"chain"`
	Node     nodeconfig.Config     `yaml:"node"`
	Parser   parserconfig.Config   `yaml:"parsing"`
	Database databaseconfig.Config `yaml:"database"`
	Logging  loggingconfig.Config  `yaml:"logging"`

	// The following are there to support modules which config are present if they are enabled

	Telemetry *telemetry.Config       `yaml:"telemetry,omitempty"`
	Pruning   *pruning.Config         `yaml:"pruning,omitempty"`
	PriceFeed *pricefeedconfig.Config `yaml:"pricefeed,omitempty"`
}
