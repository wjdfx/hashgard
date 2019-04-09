package queriers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/types"
)

func QueryIssue(ctx sdk.Context, issueID string, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, sdk.Error) {
	issue := keeper.GetIssue(ctx, issueID)
	if issue == nil {
		return nil, errors.ErrUnknownIssue(types.DefaultCodespace, issueID)
	}

	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), issue)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
func QueryIssues(ctx sdk.Context, accAddress string, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, sdk.Error) {
	issues := keeper.GetIssues(ctx, accAddress)
	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), issues)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
