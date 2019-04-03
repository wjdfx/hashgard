package queriers

import (
	abci "github.com/tendermint/tendermint/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/exchange/keeper"
)

type QueryOrdersParams struct {
	Seller sdk.AccAddress
}

func NewQueryOrdersParams(addr sdk.AccAddress) QueryOrdersParams {
	return QueryOrdersParams{
		Seller: addr,
	}
}

func QueryOrdersByAddress(ctx sdk.Context, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, sdk.Error) {
	// 解析请求参数

	// 查询状态

	// 序列化返回结果

	return []byte{}, nil
}