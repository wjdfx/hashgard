package queriers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hashgard/hashgard/x/issue/domain"
	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/keepers"
)

func QueryIssue(ctx sdk.Context, issueID string, req abci.RequestQuery, keeper keepers.Keeper) ([]byte, sdk.Error) {
	//var params QueryIssueParams
	//err := keeper.Getcdc().UnmarshalJSON(req.Data, &params)
	//if err != nil {
	//	return nil, sdk.ErrUnknownRequest(sdk.AppendMsgToErr("incorrectly formatted request data", err.Error()))
	//}

	issue := keeper.GetIssue(ctx, issueID)
	if issue == nil {
		return nil, errors.ErrUnknownIssue(domain.DefaultCodespace, issueID)
	}

	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), issue)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
