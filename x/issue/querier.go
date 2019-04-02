package issue

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hashgard/hashgard/x/issue/keepers"
	"github.com/hashgard/hashgard/x/issue/queriers"
)

// query endpoints supported by the governance Querier
const (
	QueryParams = "params"
	QueryIssues = "issues"
	QueryIssue  = "issue"
)

type QueryIssueParams struct {
	IssueID string
}

func NewQuerier(keeper keepers.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case QueryIssue:
			return queriers.QueryIssue(ctx, path[1], req, keeper)
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
