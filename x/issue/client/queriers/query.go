package queriers

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue/params"
	"github.com/hashgard/hashgard/x/issue/types"
)

func getQueryIssuePath(issueID string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssue, issueID)
}
func getQueryIssueAllowancePath(issueID string, owner sdk.AccAddress, spender sdk.AccAddress) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryAllowance, issueID, owner.String(), spender.String())
}
func getQueryIssueSearchPath(symbol string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QuerySearch, symbol)
}
func getQueryIssuesPath() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssues)
}

func QueryIssueBySymbol(symbol string, cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(getQueryIssueSearchPath(symbol), nil)
}

func QueryIssueByID(issueID string, cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(getQueryIssuePath(issueID), nil)
}
func QueryIssueAllowance(issueID string, owner sdk.AccAddress, spender sdk.AccAddress, cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(getQueryIssueAllowancePath(issueID, owner, spender), nil)
}

func QueryIssuesList(params params.IssueQueryParams, cdc *codec.Codec, cliCtx context.CLIContext) ([]byte, error) {
	bz, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	return cliCtx.QueryWithData(getQueryIssuesPath(), bz)
}
