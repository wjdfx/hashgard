package queriers

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/types"
)

func GetQueryIssuePath(issueID string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssue, issueID)
}
func GetQueryAddressPath(owner sdk.AccAddress) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssues, owner.String())
}

func QueryIssueByID(issueID string, cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryIssuePath(issueID), nil)
}

func QueryIssuesByAddress(owner sdk.AccAddress, cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryAddressPath(owner), nil)
}
