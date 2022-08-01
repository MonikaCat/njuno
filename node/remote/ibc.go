package remote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MonikaCat/njuno/types"
)

// IBCParams implements node.Node
func (cp *Node) IBCParams() (types.IBCTransactionParams, error) {
	resp, err := http.Get(fmt.Sprintf("%s/ibc/apps/transfer/v1/params", cp.RESTNode))
	if err != nil {
		return types.IBCTransactionParams{}, fmt.Errorf("error while getting ibc params: %s", err)
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.IBCTransactionParams{}, fmt.Errorf("error while processing ibc params: %s", err)
	}

	var stakingPool types.IBCTransferParams
	err = json.Unmarshal(bz, &stakingPool)
	if err != nil {
		return types.IBCTransactionParams{}, fmt.Errorf("error while unmarshaling ibc params: %s", err)
	}

	return stakingPool.Params, nil
}
