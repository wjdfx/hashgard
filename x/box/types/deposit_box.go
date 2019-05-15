package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type DepositBox struct {
	StartTime          int64            `json:"start_time"`
	EstablishTime      int64            `json:"establish_time"`
	MaturityTime       int64            `json:"maturity_time"`
	BottomLine         sdk.Int          `json:"bottom_line"`
	Interest           BoxToken         `json:"interest"`
	Price              sdk.Int          `json:"price"`
	PerCoupon          sdk.Dec          `json:"per_coupon"`
	Share              sdk.Int          `json:"share"`
	TotalDeposit       sdk.Int          `json:"total_deposit"`
	InterestInjections []AddressDeposit `json:"interest_injections"`
}

type DepositBoxDepositToList []AddressDeposit

//nolint
func (bi DepositBox) String() string {
	return fmt.Sprintf(`DepositInfo:
  StartTime:			%d
  EstablishTime:		%d
  MaturityTime:			%d
  BottomLine:			%s
  Interest:			%s
  Price:			%s
  PerCoupon:			%s
  Share:			%s
  TotalDeposit:			%s
  InterestInjection:			%s`,
		bi.StartTime,
		bi.EstablishTime,
		bi.MaturityTime,
		bi.BottomLine.String(),
		bi.Interest.String(),
		bi.Price.String(),
		bi.PerCoupon.String(),
		bi.Share.String(),
		bi.TotalDeposit.String(),
		bi.InterestInjections)
}

//nolint
func (bi DepositBoxDepositToList) String() string {
	out := fmt.Sprintf("%-44s|%s\n",
		"Address", "Amount")
	for _, box := range bi {
		out += fmt.Sprintf("%-44s|%s\n",
			box.Address.String(), box.Amount.String())
	}
	return strings.TrimSpace(out)
}
