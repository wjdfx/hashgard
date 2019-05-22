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
	WithdrawalShare    sdk.Int          `json:"withdrawal_share"`
	WithdrawalInterest sdk.Int          `json:"withdrawal_interest"`
	InterestInjections []AddressDeposit `json:"interest_injections"`
}

type DepositBoxDepositInterest struct {
	Address  sdk.AccAddress `json:"address"`
	Amount   sdk.Int        `json:"amount"`
	Interest sdk.Int        `json:"interest"`
}

type DepositBoxDepositInterestList []DepositBoxDepositInterest

//nolint
func (bi DepositBoxDepositInterest) String() string {
	return fmt.Sprintf(`
  Address:			%s
  Amount:			%s
  Interest:			%s`,
		bi.Address.String(), bi.Amount.String(), bi.Interest.String())
}

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
  WithdrawalShare:			%s,
  WithdrawalInterest:			%s,
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
		bi.WithdrawalShare.String(),
		bi.WithdrawalInterest.String(),
		bi.InterestInjections)
}

//nolint
func (bi DepositBoxDepositInterestList) String() string {
	out := fmt.Sprintf("%-44s|%-40s|%s\n",
		"Address", "Amount", "Interest")
	for _, box := range bi {
		out += fmt.Sprintf("%-44s|%-40s|%s\n",
			box.Address.String(), box.Amount.String(), box.Interest.String())
	}
	return strings.TrimSpace(out)
}
