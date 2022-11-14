package remote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// StkingPool implements node.Node
func (cp *Node) StakingPool() (stakingtypes.Pool, error) {
	resp, err := http.Get(fmt.Sprintf("%s/cosmos/staking/v1beta1/pool", cp.RESTNode))
	if err != nil {
		return stakingtypes.Pool{}, fmt.Errorf("error while getting staking pool: %s", err)
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return stakingtypes.Pool{}, fmt.Errorf("error while processing staking pool: %s", err)
	}

	var stakingPool stakingtypes.Pool
	err = json.Unmarshal(bz, &stakingPool)
	if err != nil {
		return stakingtypes.Pool{}, fmt.Errorf("error while unmarshaling staking pool: %s", err)
	}

	return stakingPool, nil
}

// -------------------------------------------------------------------------------------------------------------------

// Validators implements node.Node
func (cp *Node) Validators(height int64) (*tmctypes.ResultValidators, error) {
	vals := &tmctypes.ResultValidators{
		BlockHeight: height,
	}

	page := 1
	perPage := 100 // maximum 100 entries per page
	stop := false
	for !stop {
		result, err := cp.client.Validators(cp.ctx, &height, &page, &perPage)
		if err != nil {
			return nil, err
		}
		vals.Validators = append(vals.Validators, result.Validators...)
		vals.Count += result.Count
		vals.Total = result.Total
		page++
		stop = vals.Count == vals.Total
	}

	return vals, nil
}

// TotalDelegations implements node.Node
func (cp *Node) TotalDelegations(address string) (sdk.Coin, error) {
	resp, err := http.Get(fmt.Sprintf("%s/cosmos/staking/v1beta1/delegations/%s", cp.RESTNode, address))
	if err != nil {
		return sdk.Coin{}, fmt.Errorf("error while getting total delegations value of address %s: %s", address, err)
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sdk.Coin{}, fmt.Errorf("error while processing total delegations value of address %s: %s", address, err)
	}

	var delegation stakingtypes.DelegationResponse
	err = json.Unmarshal(bz, &delegation)
	if err != nil {
		return sdk.Coin{}, fmt.Errorf("error while unmarshaling total delegations value of address %s: %s", address, err)
	}

	return delegation.Balance, nil
}
