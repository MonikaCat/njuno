package staking

import (
	"fmt"

	parsecmdtypes "github.com/MonikaCat/njuno/cmd/parse/types"
	"github.com/MonikaCat/njuno/types"
	"github.com/MonikaCat/njuno/types/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// validatorsCmd returns a Cobra command that allows to fix the validators information
func validatorsCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "validators",
		Short: "Fix the information about validators reading the details from .yaml file",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parsecmdtypes.GetParserContext(config.Cfg, parseConfig)
			if err != nil {
				return err
			}

			var validators []types.Validator
			var validatorsDescription []types.ValidatorDescription
			var validatorsCommission []types.ValidatorCommission

			for _, val := range parseCtx.ValidatorsList.Validators {
				consAddr := sdk.ConsAddress(val.Validator.Hex).String()
				validatorAddress, err := sdk.ValAddressFromHex(val.Validator.Hex)
				if err != nil {
					fmt.Printf("failed to convert validator address from hex: %s", err)
				}
				validators = append(validators, types.NewValidator(consAddr, validatorAddress.String(), "", val.Validator.Address, 1))
				validatorsDescription = append(validatorsDescription, types.NewValidatorDescription(consAddr, val.Validator.Details, val.Validator.Identity, val.Validator.Moniker, 1))
				validatorsCommission = append(validatorsCommission, types.NewValidatorCommission(consAddr, val.Validator.Commission, 1))
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
			return nil
		},
	}

}
