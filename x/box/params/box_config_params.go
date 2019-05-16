package params

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Param Config issue for issue
type BoxConfigParams struct {
	MinDeposit    sdk.Coins `json:"min_deposit"`
	StartingBoxId uint64    `json:"starting_box_id"`
}

func (dp BoxConfigParams) String() string {
	return fmt.Sprintf(`Box Params:Min Deposit:%s`, dp.MinDeposit)
}

// Checks equality of BoxConfigParams
func (dp BoxConfigParams) Equal(dp2 BoxConfigParams) bool {
	return dp.MinDeposit.IsEqual(dp2.MinDeposit)
}

// Params returns all of the issue params
type Params struct {
	BoxConfigParams BoxConfigParams `json:"issue_params"`
}

func (iss Params) String() string {
	return iss.BoxConfigParams.String()
}
func NewParams(config BoxConfigParams) Params {
	return Params{
		BoxConfigParams: config,
	}
}
