package utils

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"

	types "github.com/MonikaCat/njuno/types"
	"github.com/MonikaCat/njuno/types/config"
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
