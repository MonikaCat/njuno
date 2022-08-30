package staking

import (
	parsecmdtypes "github.com/forbole/njuno/cmd/parse/types"
	staking "github.com/forbole/njuno/modules/staking/utils"
	"github.com/forbole/njuno/types/config"
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

			// query the latest validators list
			validatorsLists := staking.GetLatestValidatorsList()

			// parse validators list
			validators, validatorsCommission, validatorsDescription, validatorsStatus, validatorsVP := staking.ParseValidatorsList(validatorsLists, 1)

			err = parseCtx.Database.SaveValidators(validators)
			if err != nil {
				log.Error().Str("module", "staking").Err(err).Int64("height", 1).
					Msg("error while saving validators")
			}

			err = parseCtx.Database.SaveValidatorCommission(validatorsCommission)
			if err != nil {
				log.Error().Str("module", "staking").Err(err).Int64("height", 1).
					Msg("error while saving validators commission")
			}

			err = parseCtx.Database.SaveValidatorDescription(validatorsDescription)
			if err != nil {
				log.Error().Str("module", "staking").Err(err).Int64("height", 1).
					Msg("error while saving validators description")
			}

			err = parseCtx.Database.SaveValidatorsStatus(validatorsStatus)
			if err != nil {
				log.Error().Str("module", "staking").Err(err).Int64("height", 1).
					Msg("error while saving validators status")
			}

			err = parseCtx.Database.SaveValidatorsVotingPower(validatorsVP)
			if err != nil {
				log.Error().Str("module", "staking").Err(err).Int64("height", 1).
					Msg("error while saving validators voting power")
			}

			return nil
		},
	}

}
