package parse

import (
	"github.com/spf13/cobra"

	parsecmdtypes "github.com/MonikaCat/njuno/cmd/parse/types"

	parseblocks "github.com/MonikaCat/njuno/cmd/parse/blocks"
	parsegenesis "github.com/MonikaCat/njuno/cmd/parse/genesis"
	parsetransactions "github.com/MonikaCat/njuno/cmd/parse/transactions"
)

// NewParseCmd returns the Cobra command allowing to parse chain data without having to re-sync the whole database
func NewParseCmd(parseCfg *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "parse",
		Short:             "Parse data without the need to re-sync the whole database from scratch",
		PersistentPreRunE: runPersistentPreRuns(parsecmdtypes.ReadConfigPreRunE(parseCfg)),
	}

	cmd.AddCommand(
		parseblocks.NewBlocksCmd(parseCfg),
		parsegenesis.NewGenesisCmd(parseCfg),
		parsetransactions.NewTransactionsCmd(parseCfg),
	)

	return cmd
}

func runPersistentPreRuns(preRun func(_ *cobra.Command, _ []string) error) func(_ *cobra.Command, _ []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if root := cmd.Root(); root != nil {
			if root.PersistentPreRunE != nil {
				err := root.PersistentPreRunE(root, args)
				if err != nil {
					return err
				}
			}
		}

		return preRun(cmd, args)
	}
}
