package staking

import (
	parsecmdtypes "github.com/MonikaCat/njuno/cmd/parse/types"
	staking "github.com/MonikaCat/njuno/modules/staking/utils"
	"github.com/MonikaCat/njuno/types/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// validatorsCmd returns a Cobra command that allows to fix the validators information
func validatorsCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "validator-list",
		Short: "Fix the information about validators reading the details from .yaml file",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parsecmdtypes.GetParserContext(config.Cfg, parseConfig)
			if err != nil {
				return err
			}

			// query the latest validators status
			validatorsLists := staking.GetLatestValidatorsStatus()

			validators, validatorsDescription, validatorsCommission, validatorsStatus := staking.ParseValidatorsList(validatorsLists)

			err = parseCtx.Database.SaveValidators(validators)
			if err != nil {
				log.Error().Str("module", "staking").Err(err).Int64("height", 1).
					Msg("error while saving validators")
			}

			err = parseCtx.Database.SaveValidatorDescription(validatorsDescription)
			if err != nil {
				log.Error().Str("module", "staking").Err(err).Int64("height", 1).
					Msg("error while saving validators description")
			}

			err = parseCtx.Database.SaveValidatorCommission(validatorsCommission)
			if err != nil {
				log.Error().Str("module", "staking").Err(err).Int64("height", 1).
					Msg("error while saving validators commission")
			}

			err = parseCtx.Database.SaveValidatorsStatus(validatorsStatus)
			if err != nil {
				log.Error().Str("module", "staking").Err(err).Int64("height", 1).
					Msg("error while saving validators status")
			}

			return nil
		},
	}

}
