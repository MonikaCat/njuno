package types

import (
	"github.com/cosmos/cosmos-sdk/simapp"

	"github.com/forbole/njuno/logging"
	"github.com/forbole/njuno/types/config"

	"github.com/forbole/njuno/database"
	"github.com/forbole/njuno/database/builder"
	"github.com/forbole/njuno/modules/registrar"
)

// Config contains all the configuration for the "parse" command
type Config struct {
	registrar             registrar.Registrar
	configParser          config.Parser
	encodingConfigBuilder EncodingConfigBuilder
	setupCfg              SdkConfigSetup
	buildDb               database.Builder
	logger                logging.Logger
}

// NewConfig allows to build a new Config instance
func NewConfig() *Config {
	return &Config{}
}

// WithRegistrar sets the modules registrar to be used
func (cfg *Config) WithRegistrar(r registrar.Registrar) *Config {
	cfg.registrar = r
	return cfg
}

// GetRegistrar returns the modules registrar to be used
func (cfg *Config) GetRegistrar() registrar.Registrar {
	if cfg.registrar == nil {
		return &registrar.EmptyRegistrar{}
	}
	return cfg.registrar
}

// GetConfigParser returns the configuration parser to be used
func (cfg *Config) GetConfigParser() config.Parser {
	if cfg.configParser == nil {
		return config.DefaultConfigParser
	}
	return cfg.configParser
}

// GetEncodingConfigBuilder returns the encoding config builder to be used
func (cfg *Config) GetEncodingConfigBuilder() EncodingConfigBuilder {
	if cfg.encodingConfigBuilder == nil {
		return simapp.MakeTestEncodingConfig
	}
	return cfg.encodingConfigBuilder
}

// GetSetupConfig returns the SDK configuration builder to use
func (cfg *Config) GetSetupConfig() SdkConfigSetup {
	if cfg.setupCfg == nil {
		return DefaultConfigSetup
	}
	return cfg.setupCfg
}

// GetDBBuilder returns the database builder to be used
func (cfg *Config) GetDBBuilder() database.Builder {
	if cfg.buildDb == nil {
		return builder.Builder
	}
	return cfg.buildDb
}

// GetLogger returns the logger to be used when parsing the data
func (cfg *Config) GetLogger() logging.Logger {
	if cfg.logger == nil {
		return logging.DefaultLogger()
	}
	return cfg.logger
}
