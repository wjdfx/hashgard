package queriers

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hashgard/hashgard/x/exchange/keeper"
	"github.com/hashgard/hashgard/x/exchange/types"
)

type QueryFrozenFundParams struct {
	Seller sdk.AccAddress
}

func NewQueryFrozenFundParams(addr sdk.AccAddress) QueryFrozenFundParams {
	return QueryFrozenFundParams{
		Seller: addr,
	}
}

func QueryFrozenFund(ctx sdk.Context, cdc *codec.Codec, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, sdk.Error) {
	var params QueryFrozenFundParams
	err := cdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return []byte{}, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	fund, err := keeper.GetFrozenFundByAddr(ctx, params.Seller)
	if err != nil {
		return nil, sdk.NewError(types.DefaultCodespace, types.CodeOrderNotExist, err.Error())
	}

	bz, err := codec.MarshalJSONIndent(cdc, fund)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("could not marshal result to JSON: %s", err))
	}

	return bz, nil
}