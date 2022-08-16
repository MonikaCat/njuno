package utils

import (
	"fmt"
	"strings"

	"github.com/MonikaCat/njuno/node"

	tmjson "github.com/tendermint/tendermint/libs/json"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmtypes "github.com/tendermint/tendermint/types"
)

// ReadGenesisFileGenesisDoc reads the genesis file located at the given path
func ReadGenesisFileGenesisDoc(genesisPath string) (*tmtypes.GenesisDoc, error) {
	var genesisDoc *tmtypes.GenesisDoc
	bz, err := tmos.ReadFile(genesisPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read genesis file: %s", err)
	}

	err = tmjson.Unmarshal(bz, &genesisDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal genesis doc: %s", err)
	}

	return genesisDoc, nil
}

// GetGenesisDoc reads the genesis from node or file and returns genesis doc
func GetGenesisDoc(genesisPath string, node node.Node) (*tmtypes.GenesisDoc, error) {
	var genesisDoc *tmtypes.GenesisDoc
	if strings.TrimSpace(genesisPath) != "" {
		genDoc, err := ReadGenesisFileGenesisDoc(genesisPath)
		if err != nil {
			return nil, fmt.Errorf("error while reading genesis file: %s", err)
		}
		genesisDoc = genDoc

	} else {
		response, err := node.Genesis()
		if err != nil {
			return nil, fmt.Errorf("failed to get genesis: %s", err)
		}
		genesisDoc = response.Genesis
	}

	return genesisDoc, nil
}
