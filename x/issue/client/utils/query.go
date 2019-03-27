package utils

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/hashgard/hashgard/x/issue"
)

func QueryIssueByID(issueID string, cliCtx context.CLIContext, cdc *codec.Codec, queryRoute string) ([]byte, error) {
	params := issue.NewQueryIssueParams(issueID)
	bz, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/issue", queryRoute), bz)
	if err != nil {
		return nil, err
	}
	return res, err
}
