package queriers

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hashgard/hashgard/x/exchange/keeper"
	"github.com/hashgard/hashgard/x/exchange/types"
)

type QueryOrderParams struct {
	OrderId uint64
}

func NewQueryOrderParams(orderId uint64) QueryOrderParams {
	return QueryOrderParams{
		OrderId: orderId,
	}
}

func QueryOrder(ctx sdk.Context, cdc *codec.Codec, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, sdk.Error) {
	// 解析请求参数
	var params QueryOrderParams
	err := cdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return []byte{}, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	// 查询状态
	order, ok := keeper.GetOrder(ctx, params.OrderId)
	if !ok {
		return nil, sdk.NewError(types.DefaultCodespace, types.CodeOrderNotExist, fmt.Sprintf("this orderId is not exist : %d", params.OrderId))
	}

	// 序列化返回结果
	bz, err := codec.MarshalJSONIndent(cdc, order)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("could not marshal result to JSON: %s", err))
	}

	return bz, nil
}