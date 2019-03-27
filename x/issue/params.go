package issue

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Param around issue for issue
type IssueParams struct {
	MinDeposit sdk.Coins `json:"min_deposit"`
}

func (dp IssueParams) String() string {
	return fmt.Sprintf(`Issue Params:Min Deposit:%s`, dp.MinDeposit)
}

// Checks equality of IssueParams
func (dp IssueParams) Equal(dp2 IssueParams) bool {
	return dp.MinDeposit.IsEqual(dp2.MinDeposit)
}

// Params returns all of the issue params
type Params struct {
	IssueParams IssueParams `json:"issue_params"`
}

func (iss Params) String() string {
	return iss.IssueParams.String()
}
func NewParams(ip IssueParams) Params {
	return Params{
		IssueParams: ip,
	}
}
