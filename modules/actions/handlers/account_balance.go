package handlers

import (
	"fmt"

	"github.com/MonikaCat/njuno/modules/actions/types"

	"github.com/rs/zerolog/log"
)

func AccountBalanceHandler(ctx *types.Context, payload *types.Payload) (interface{}, error) {
	log.Debug().Str("address", payload.GetAddress()).
		Int64("height", payload.Input.Height).
		Msg("executing account balance action")

	balance, err := ctx.Node.AccountBalance(payload.GetAddress())
	if err != nil {
		return nil, fmt.Errorf("error while getting account balance: %s", err)
	}

	return types.Balance{
		Coins: types.ConvertCoins(balance),
	}, nil
}
