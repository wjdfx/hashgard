package utils

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/hashgard/hashgard/x/exchange"
)

func QueryOrderById(orderId uint64, cliCtx context.CLIContext, cdc *codec.Codec, queryRoute string) ([]byte, error) {
	params := exchange.NewQueryOrderParams(orderId)
	bz, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/order", queryRoute), bz)
	if err != nil {
		return nil, err
	}
	return res, err
}