package issue

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/queriers"
	"github.com/hashgard/hashgard/x/issue/types"
)

// query endpoints supported by the governance Querier
type QueryIssueParams struct {
	IssueID string
}

//New Querier Instance
func NewQuerier(keeper keeper.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryIssue:
			return queriers.QueryIssue(ctx, path[1], req, keeper)
		case types.QueryIssues:
			return queriers.QueryIssues(ctx, path[1], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown issue query endpoint")
		}
	}
}

// creates a new instance of QueryIssueParams
func NewQueryIssueParams(IssueID string) QueryIssueParams {
	return QueryIssueParams{
		IssueID: IssueID,
	}
}
