package config_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	nodeconfig "github.com/MonikaCat/njuno/node/config"
	"github.com/MonikaCat/njuno/node/remote"
)

func TestConfig_UnmarshalYAML(t *testing.T) {
	var remoteData = `
type: "remote"
config:
  rpc:
    client_name: "njuno"
    max_connections: 1
    address: "http://localhost:26657"
`

	var config nodeconfig.Config
	err := yaml.Unmarshal([]byte(remoteData), &config)
	require.NoError(t, err)
	require.IsType(t, &remote.Details{}, config.Details)
}

func TestConfig_MarshalYAML(t *testing.T) {
	config := nodeconfig.Config{
		Type: nodeconfig.TypeRemote,
		Details: &remote.Details{
			RPC: &remote.RPCConfig{
				ClientName:     "njuno",
				Address:        "http://localhost:26657",
				MaxConnections: 10,
			},
		},
	}
	bz, err := yaml.Marshal(&config)
	require.NoError(t, err)

	expected := `
type: remote
config:
    rpc:
        client_name: njuno
        address: http://localhost:26657
        max_connections: 10
`
	require.Equal(t, strings.TrimLeft(expected, "\n"), string(bz))
}
