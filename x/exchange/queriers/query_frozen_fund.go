package queriers

import (
	abci "github.com/tendermint/tendermint/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/exchange/keeper"
)

type QueryFrozenFundParams struct {
	Seller sdk.AccAddress
}

func NewQueryFrozenFundParams(addr sdk.AccAddress) QueryFrozenFundParams {
	return QueryFrozenFundParams{
		Seller: addr,
	}
}

func QueryFrozenFund(ctx sdk.Context, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, sdk.Error) {
	// 解析请求参数

	// 查询状态

	// 序列化返回结果

	return []byte{}, nil
}