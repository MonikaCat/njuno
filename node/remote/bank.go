package remote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MonikaCat/njuno/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// AccountBalance implements node.Node
func (cp *Node) AccountBalance(address string) (sdk.Coins, error) {
	resp, err := http.Get(fmt.Sprintf("%s/cosmos/bank/v1beta1/balances/%s", cp.RESTNode, address))
	if err != nil {
		return sdk.Coins{}, fmt.Errorf("error while getting account balance of address %s: %s", address, err)
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sdk.Coins{}, fmt.Errorf("error while processing account balance of address %s: %s", address, err)
	}

	var balance types.QueryAllBalancesResponse
	err = json.Unmarshal(bz, &balance)
	if err != nil {
		return sdk.Coins{}, fmt.Errorf("error while unmarshaling account balance of address %s: %s", address, err)
	}

	return balance.Balances, nil
}

// -------------------------------------------------------------------------------------------------------------------

// Supply implements node.Node
func (cp *Node) Supply() (sdk.Coins, error) {
	resp, err := http.Get(fmt.Sprintf("%s/cosmos/bank/v1beta1/supply/unom", cp.RESTNode))
	if err != nil {
		return sdk.Coins{}, fmt.Errorf("error while getting total supply: %s", err)
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sdk.Coins{}, fmt.Errorf("error while processing total supply: %s", err)
	}

	var supply *banktypes.QuerySupplyOfResponse
	err = json.Unmarshal(bz, &supply)
	if err != nil {
		return sdk.Coins{}, fmt.Errorf("error while unmarshaling supply: %s", err)
	}
	var totalSupply []sdk.Coin
	totalSupply = append(totalSupply, supply.Amount)

	return totalSupply, nil
}
