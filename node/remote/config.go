package remote

import (
	"fmt"
)

// Details represents a node details for a remote node
type Details struct {
	RPC  *RPCConfig  `yaml:"rpc"`
	REST *RESTConfig `yaml:"rest"`
}

func NewDetails(rpc *RPCConfig, rest *RESTConfig) *Details {
	return &Details{
		RPC:  rpc,
		REST: rest,
	}
}

func DefaultDetails() *Details {
	return NewDetails(DefaultRPCConfig(), DefaultRESTConfig())
}

// Validate implements node.Details
func (d *Details) Validate() error {
	if d.RPC == nil {
		return fmt.Errorf("rpc config cannot be null")
	}
	if d.REST == nil {
		return fmt.Errorf("rest config cannot be null")
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// RPCConfig contains the configuration for the RPC endpoint
type RPCConfig struct {
	ClientName     string `yaml:"client_name"`
	Address        string `yaml:"address"`
	MaxConnections int    `yaml:"max_connections"`
}

// NewRPCConfig allows to build a new RPCConfig instance
func NewRPCConfig(clientName, address string, maxConnections int) *RPCConfig {
	return &RPCConfig{
		ClientName:     clientName,
		Address:        address,
		MaxConnections: maxConnections,
	}
}

// DefaultRPCConfig returns the default instance of RPCConfig
func DefaultRPCConfig() *RPCConfig {
	return NewRPCConfig("njuno", "http://localhost:26657", 20)
}

// --------------------------------------------------------------------------------------------------------------------

// RESTConfig contains the configuration for the REST endpoint
type RESTConfig struct {
	Address string `yaml:"address"`
}

// NewRESTConfig allows to build a new RESTConfig instance
func NewRESTConfig(address string) *RESTConfig {
	return &RESTConfig{
		Address: address,
	}
}

// DefaultRESTConfig returns the default instance of RESTConfig
func DefaultRESTConfig() *RESTConfig {
	return NewRESTConfig("http://localhost:26657")
}
