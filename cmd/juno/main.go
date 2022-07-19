package main

import (
	"os"

	"github.com/MonikaCat/njuno/cmd/parse/types"

	"github.com/MonikaCat/njuno/modules/messages"
	"github.com/MonikaCat/njuno/modules/registrar"

	"github.com/MonikaCat/njuno/cmd"
)

func main() {
	// JunoConfig the runner
	config := cmd.NewConfig("juno").
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
