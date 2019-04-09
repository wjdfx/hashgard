package params

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Param Config issue for issue
type IssueConfigParams struct {
	MinDeposit sdk.Coins `json:"min_deposit"`
}

func (dp IssueConfigParams) String() string {
	return fmt.Sprintf(`Issue Params:Min Deposit:%s`, dp.MinDeposit)
}

// Checks equality of IssueConfigParams
func (dp IssueConfigParams) Equal(dp2 IssueConfigParams) bool {
	return dp.MinDeposit.IsEqual(dp2.MinDeposit)
}

// Params returns all of the issue params
type Params struct {
	IssueConfigParams IssueConfigParams `json:"issue_params"`
}

func (iss Params) String() string {
	return iss.IssueConfigParams.String()
}
func NewParams(ip IssueConfigParams) Params {
	return Params{
		IssueConfigParams: ip,
	}
}
