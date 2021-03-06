package queriers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/keeper"
	"github.com/hashgard/hashgard/x/box/params"
	abci "github.com/tendermint/tendermint/abci/types"
)

func QueryBox(ctx sdk.Context, boxID string, keeper keeper.Keeper) ([]byte, sdk.Error) {
	box := keeper.GetBox(ctx, boxID)
	if box == nil {
		return nil, errors.ErrUnknownBox(boxID)
	}

	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), box)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}

func QueryName(ctx sdk.Context, boxType string, name string, keeper keeper.Keeper) ([]byte, sdk.Error) {
	box := keeper.SearchBox(ctx, boxType, name)
	if box == nil {
		return nil, errors.ErrUnknownBox(name)
	}

	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), box)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
func QueryDepositAmountFromDepositBox(ctx sdk.Context, boxID string, accAddress string, keeper keeper.Keeper) ([]byte, sdk.Error) {
	address, err := sdk.AccAddressFromBech32(accAddress)
	if err != nil {
		return nil, sdk.ErrInvalidAddress(accAddress)
	}
	amount := keeper.GetDepositByAddress(ctx, boxID, address)

	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), amount)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
func QueryList(ctx sdk.Context, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, sdk.Error) {
	var params params.BoxQueryParams
	err := keeper.Getcdc().UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(sdk.AppendMsgToErr("incorrectly formatted request data", err.Error()))
	}

	boxs := keeper.List(ctx, params)
	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), boxs)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
func QueryDepositList(ctx sdk.Context, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, sdk.Error) {
	var params params.BoxQueryDepositListParams
	err := keeper.Getcdc().UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(sdk.AppendMsgToErr("incorrectly formatted request data", err.Error()))
	}

	boxs := keeper.QueryDepositListFromDepositBox(ctx, params.BoxId, params.Owner)
	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), boxs)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
