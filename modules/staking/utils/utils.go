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

// GetLatestValidatorsList queries the latest validators list, stores it inside yaml file,
// and returns an array of validators
func GetLatestValidatorsList() *types.ValidatorsList {
	validatorsCmd := exec.Command("sh", "-c", "~/.njuno/validators_query.sh")

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

// ParseValidatorsList parses the validators list and returns arrays of validators,
// validators description, validators commission and validators status
func ParseValidatorsList(validatorsList *types.ValidatorsList, height int64) ([]types.Validator, []types.ValidatorCommission, []types.ValidatorDescription, []types.ValidatorStatus, []types.ValidatorVotingPower) {
	var validators []types.Validator
	var validatorsCommission []types.ValidatorCommission
	var validatorsDescription []types.ValidatorDescription
	var validatorsStatus []types.ValidatorStatus
	var validatorsVP []types.ValidatorVotingPower

	for _, val := range validatorsList.Validators {
		consAddr := sdk.ConsAddress(val.Validator.Address)

		validators = append(validators, types.NewValidator(consAddr.String(), val.Validator.Address, height))
		validatorsCommission = append(validatorsCommission, types.NewValidatorCommission(consAddr.String(), val.Validator.Address, val.Validator.Commission, val.Validator.MinSelfDelegation, height))
		validatorsDescription = append(validatorsDescription, types.NewValidatorDescription(consAddr.String(), val.Validator.Address, val.Validator.Details, val.Validator.Identity, val.Validator.Moniker, height))
		validatorsStatus = append(validatorsStatus, types.NewValidatorStatus(consAddr.String(), val.Validator.Address, val.Validator.InActiveSet, val.Validator.Jailed, val.Validator.Tombstoned, height))
		validatorsVP = append(validatorsVP, types.NewValidatorVotingPower(consAddr.String(), val.Validator.Address, val.Validator.VotingPower, height))
	}

	return validators, validatorsCommission, validatorsDescription, validatorsStatus, validatorsVP
}
