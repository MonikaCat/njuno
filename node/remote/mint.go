package remote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MonikaCat/njuno/types"
)

// Inflation implements node.Node
func (cp *Node) Inflation() (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/cosmos/mint/v1beta1/inflation", cp.RESTNode))
	if err != nil {
		return "", fmt.Errorf("error while getting inflation: %s", err)
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error while processing inflation: %s", err)
	}
	var inflation types.InflationResponse
	err = json.Unmarshal(bz, &inflation)
	if err != nil {
		return "", fmt.Errorf("error while unmarshaling inflation: %s", err)
	}

	return inflation.Inflation, nil
}
