package queriers

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/types"
)

func GetQueryIssuePath(issueID string, queryRoute string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, queryRoute, types.QueryIssue, issueID)
}
func GetQueryAddressPath(owner sdk.AccAddress, queryRoute string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, queryRoute, types.QueryIssues, owner.String())
}

func QueryIssueByID(issueID string, cliCtx context.CLIContext, queryRoute string) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryIssuePath(issueID, queryRoute), nil)
}

func QueryIssuesByAddress(owner sdk.AccAddress, cliCtx context.CLIContext, queryRoute string) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryAddressPath(owner, queryRoute), nil)
}
