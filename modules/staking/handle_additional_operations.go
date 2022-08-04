package staking

import (
	"github.com/MonikaCat/njuno/types"
	"github.com/rs/zerolog/log"
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

	height, err := m.db.GetLastBlockHeight()
	if err != nil {
		log.Printf("error while getting the latest block height: %s ", err)
	}

	// get validator address from validator_description table
	validatorDescList, err := m.db.GetValidatorsDescription()
	if err != nil {
		log.Printf("error while getting validators description: %s ", err)
	}

	for _, val := range m.validatorsList.Validators {
		for _, desc := range validatorDescList {
			if desc.Identity == val.Validator.Identity && desc.Moniker == val.Validator.Moniker {
				validatorsDesc = append(validatorsDesc, types.NewValidatorCommission(desc.OperatorAddress, val.Validator.Commission, height))

			}
		}
	}

	err = m.db.SaveValidatorCommission(validatorsDesc)
	if err != nil {
		log.Printf("error while saving validator commission: %s ", err)
	}
	return nil
}
