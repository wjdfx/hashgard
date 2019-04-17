package issue

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/queriers"
	"github.com/hashgard/hashgard/x/issue/types"
)

//New Querier Instance
func NewQuerier(keeper keeper.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryIssue:
			return queriers.QueryIssue(ctx, path[1], req, keeper)
		case types.QuerySearch:
			return queriers.QuerySymbol(ctx, path[1], req, keeper)
		case types.QueryIssues:
			return queriers.QueryIssues(ctx, path[0], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown issue query endpoint")
		}
	}
}
