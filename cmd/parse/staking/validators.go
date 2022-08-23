package staking

import (
	parsecmdtypes "github.com/MonikaCat/njuno/cmd/parse/types"
	staking "github.com/MonikaCat/njuno/modules/staking/utils"
	"github.com/MonikaCat/njuno/types"
	"github.com/MonikaCat/njuno/types/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
			var validators []types.Validator
			var validatorsDescription []types.ValidatorDescription
			var validatorsCommission []types.ValidatorCommission
			var validatorsStatus []types.ValidatorStatus

			// query the latest validators status
			validatorsLists := staking.GetLatestValidatorsStatus()

			for _, val := range validatorsLists.Validators {
				consAddr := sdk.ConsAddress(val.Validator.Address).String()

				validators = append(validators, types.NewValidator(consAddr, val.Validator.Address, 1))
				validatorsDescription = append(validatorsDescription, types.NewValidatorDescription(consAddr, val.Validator.Details, val.Validator.Identity, val.Validator.Moniker, 1))
				validatorsCommission = append(validatorsCommission, types.NewValidatorCommission(consAddr, val.Validator.Commission, val.Validator.MinSelfDelegation, 1))
				validatorsStatus = append(validatorsStatus, types.NewValidatorStatus(consAddr, val.Validator.InActiveSet, val.Validator.Jailed, val.Validator.Tombstoned, 1))
			}

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
