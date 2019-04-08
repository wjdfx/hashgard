package exchange

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hashgard/hashgard/x/exchange/keeper"
	"github.com/hashgard/hashgard/x/exchange/queriers"
)

// query endpoints supported by the governance Querier
const (
	QueryOrder				= "order"
	QueryFrozenFund			= "frozen"
	QueryAllOrdersByAddress	= "orders"
)

func NewQuerier(keeper keeper.Keeper, cdc *codec.Codec) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case QueryOrder:
			return queriers.QueryOrder(ctx, cdc, req, keeper)
		case QueryFrozenFund:
			return queriers.QueryFrozenFund(ctx, cdc, req, keeper)
		case QueryAllOrdersByAddress:
			return queriers.QueryOrdersByAddress(ctx, cdc, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown exchange query endpoint")
		}
	}
}