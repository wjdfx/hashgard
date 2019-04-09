package queriers

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/types"
)

func GetQueryIssuePath(issueID string, queryRoute string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, queryRoute, types.QueryIssue, issueID)
}
func GetQueryAddressPath(issuer sdk.AccAddress, queryRoute string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, queryRoute, types.QueryIssues, issuer.String())
}

func QueryIssueByID(issueID string, cliCtx context.CLIContext, cdc *codec.Codec, queryRoute string) ([]byte, error) {
	res, err := cliCtx.QueryWithData(GetQueryIssuePath(issueID, queryRoute), nil)
	if err != nil {
		return nil, err
	}
	return res, err
}
func QueryIssuesByAddress(issuer sdk.AccAddress, cliCtx context.CLIContext, cdc *codec.Codec, queryRoute string) ([]byte, error) {
	res, err := cliCtx.QueryWithData(GetQueryAddressPath(issuer, queryRoute), nil)
	if err != nil {
		return nil, err
	}
	return res, err
}
