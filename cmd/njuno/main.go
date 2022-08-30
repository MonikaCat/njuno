package main

import (
	"os"

	"github.com/forbole/njuno/cmd/parse/types"

	"github.com/forbole/njuno/modules/messages"
	"github.com/forbole/njuno/modules/registrar"

	"github.com/forbole/njuno/cmd"
)

func main() {
	// NJunoConfig the runner
	config := cmd.NewConfig("njuno").
		WithParseConfig(types.NewConfig().
			WithRegistrar(registrar.NewDefaultRegistrar(
				messages.CosmosMessageAddressesParser,
			)),
		)

	// Run the commands and panic on any error
	exec := cmd.BuildDefaultExecutor(config)
	err := exec.Execute()
	if err != nil {
		os.Exit(1)
	}
}
