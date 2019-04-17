package queriers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/params"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/keeper"
)

func QueryIssue(ctx sdk.Context, issueID string, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, sdk.Error) {
	issue := keeper.GetIssue(ctx, issueID)
	if issue == nil {
		return nil, errors.ErrUnknownIssue(issueID)
	}

	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), issue)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
func QuerySymbol(ctx sdk.Context, symbol string, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, sdk.Error) {
	issue := keeper.SearchIssues(ctx, symbol)
	if issue == nil {
		return nil, errors.ErrUnknownIssue(symbol)
	}

	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), issue)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
func QueryIssues(ctx sdk.Context, accAddress string, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, sdk.Error) {
	var params params.IssueQueryParams
	err := keeper.Getcdc().UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(sdk.AppendMsgToErr("incorrectly formatted request data", err.Error()))
	}

	issues := keeper.List(ctx, params)
	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), issues)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
