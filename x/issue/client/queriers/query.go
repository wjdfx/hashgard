package queriers

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/hashgard/hashgard/x/issue/domain"
)

const (
	QueryIssue = "issue"
)

func GetQueryIssuePath(issueID string, queryRoute string) string {
	//return strings.Join([]string{domain.Custom, queryRoute, QueryIssue, issueID}, "/")
	return fmt.Sprintf("%s/%s/%s/%s", domain.Custom, queryRoute, QueryIssue, issueID)
}

func QueryIssueByID(issueID string, cliCtx context.CLIContext, cdc *codec.Codec, queryRoute string) ([]byte, error) {
	//params := queriers.NewQueryIssueParams(issueID)
	//bz, err := cdc.MarshalJSON(params)
	//if err != nil {
	//	return nil, err
	//}
	res, err := cliCtx.QueryWithData(GetQueryIssuePath(issueID, queryRoute), nil)
	if err != nil {
		return nil, err
	}
	return res, err
}
