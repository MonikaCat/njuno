package remote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/forbole/njuno/types"
)

// IBCTransferParams implements node.Node
func (cp *Node) IBCTransferParams() (types.IBCTransfer, error) {
	resp, err := http.Get(fmt.Sprintf("%s/ibc/apps/transfer/v1/params", cp.RESTNode))
	if err != nil {
		return types.IBCTransfer{}, fmt.Errorf("error while getting ibc transfer params: %s", err)
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.IBCTransfer{}, fmt.Errorf("error while processing ibc transfer params: %s", err)
	}

	var params types.IBCTransfer
	err = json.Unmarshal(bz, &params)
	if err != nil {
		return types.IBCTransfer{}, fmt.Errorf("error while unmarshaling ibc transfer params: %s", err)
	}

	return params, nil
}
