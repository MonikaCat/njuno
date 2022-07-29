package staking

import (
	"io/ioutil"

	"github.com/MonikaCat/njuno/types"
	"github.com/MonikaCat/njuno/types/config"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

// RunAdditionalOperations implements modules.AdditionalOperationsModule
func (m *Module) RunAdditionalOperations() error {
	err := m.saveValidatorsCommission()
	if err != nil {
		return err
	}

	return nil
}

// saveValidatorsCommission stores validator commision in database 
func (m *Module) saveValidatorsCommission() error {
	var validatorsDesc []types.ValidatorCommission

	cfg := config.Cfg.Parser
	validatorList := &types.ValidatorsList{}
	yamlFile, err := ioutil.ReadFile(cfg.ValidatorsListFilePath)
	if err != nil {
		log.Printf("error while reading yaml file: %s ", err)
	}
	err = yaml.Unmarshal(yamlFile, validatorList)
	if err != nil {
		log.Printf("error while unmarshaling yaml file: %s ", err)
	}

	height, err := m.db.GetLastBlockHeight()
	if err != nil {
		log.Printf("error while getting the latest block height: %s ", err)
	}

	for _, val := range validatorList.Validators {
		validatorsDesc = append(validatorsDesc, types.NewValidatorCommission(val.Validator.Address, val.Validator.Commission, height))
	}

	err = m.db.SaveValidatorCommission(validatorsDesc)
	if err != nil {
		log.Printf("error while saving validator commission: %s ", err)
	}
	return nil
}
