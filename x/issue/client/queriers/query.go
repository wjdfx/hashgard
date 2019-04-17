package queriers

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/hashgard/hashgard/x/issue/params"
	"github.com/hashgard/hashgard/x/issue/types"
)

func GetQueryIssuePath(issueID string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssue, issueID)
}
func GetQueryIssueSearchPath(symbol string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QuerySearch, symbol)
}
func GetQueryIssuesPath() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssues)
}

func QueryIssueBySymbol(symbol string, cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryIssueSearchPath(symbol), nil)
}

func QueryIssueByID(issueID string, cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryIssuePath(issueID), nil)
}

func QueryIssuesList(params params.IssueQueryParams, cdc *codec.Codec, cliCtx context.CLIContext) ([]byte, error) {
	bz, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	return cliCtx.QueryWithData(GetQueryIssuesPath(), bz)
}
