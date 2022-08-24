package utils

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"

	types "github.com/MonikaCat/njuno/types"
	"github.com/MonikaCat/njuno/types/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

// GetLatestValidatorsStatus queries the latest validators status and stores it inside yaml file,
// returning an array of validators
func GetLatestValidatorsStatus() *types.ValidatorsList {
	validatorsCmd := exec.Command("sh", "-c", "~/.njuno/query_validators.sh")

	cmdOutput := &bytes.Buffer{}
	validatorsCmd.Stdout = cmdOutput
	validatorsCmd.Stderr = os.Stderr
	if err := validatorsCmd.Run(); err != nil {
		return nil
	}

	// Get the validators list
	validatorsList := &types.ValidatorsList{}
	yamlFile, err := ioutil.ReadFile(config.Cfg.Parser.ValidatorsListFilePath)
	if err != nil {
		log.Printf("error while reading validators list yaml file: %s ", err)
	}

	err = yaml.Unmarshal(yamlFile, validatorsList)
	if err != nil {
		log.Printf("error while unmarshaling validators list yaml file: %s ", err)
	}

	return validatorsList
}

func ParseValidatorsList(validatorsList *types.ValidatorsList) ([]types.Validator, []types.ValidatorDescription, []types.ValidatorCommission, []types.ValidatorStatus) {
	var validators []types.Validator
	var validatorsDescription []types.ValidatorDescription
	var validatorsCommission []types.ValidatorCommission
	var validatorsStatus []types.ValidatorStatus

	for _, val := range validatorsList.Validators {
		consAddr := sdk.ConsAddress(val.Validator.Address).String()

		validators = append(validators, types.NewValidator(consAddr, val.Validator.Address, 1))
		validatorsDescription = append(validatorsDescription, types.NewValidatorDescription(consAddr, val.Validator.Details, val.Validator.Identity, val.Validator.Moniker, 1))
		validatorsCommission = append(validatorsCommission, types.NewValidatorCommission(consAddr, val.Validator.Commission, val.Validator.MinSelfDelegation, 1))
		validatorsStatus = append(validatorsStatus, types.NewValidatorStatus(consAddr, val.Validator.InActiveSet, val.Validator.Jailed, val.Validator.Tombstoned, 1))
	}

	return validators, validatorsDescription, validatorsCommission, validatorsStatus
}
