package handlers

import (
	"fmt"

	"github.com/forbole/njuno/modules/actions/types"

	"github.com/rs/zerolog/log"
)

func TotalDelegationsAmountHandler(ctx *types.Context, payload *types.Payload) (interface{}, error) {
	log.Debug().Str("address", payload.GetAddress()).
		Int64("height", payload.Input.Height).
		Msg("executing account delegations action")

	balance, err := ctx.Node.TotalDelegations(payload.GetAddress())
	if err != nil {
		return nil, fmt.Errorf("error while getting account delegations value: %s", err)
	}

	return types.Coin{
		Amount: balance.Amount.String(),
		Denom:  balance.Denom,
	}, nil
}
