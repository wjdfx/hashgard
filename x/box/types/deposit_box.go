package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type DepositBox struct {
	Status            string           `json:"status"`
	StartTime         time.Time        `json:"start_time"`
	EstablishTime     time.Time        `json:"establish_time"`
	MaturityTime      time.Time        `json:"maturity_time"`
	BottomLine        sdk.Int          `json:"bottom_line"`
	Interest          sdk.Coin         `json:"interest"`
	Price             sdk.Int          `json:"price"`
	Coupon            sdk.Int          `json:"coupon"`
	Share             sdk.Int          `json:"share"`
	TotalDeposit      sdk.Int          `json:"total_deposit"`
	InterestInjection []AddressDeposit `json:"interest_injection"`
}

type DepositBoxDepositToList []AddressDeposit

//nolint
func (bi DepositBox) String() string {
	return fmt.Sprintf(`DepositInfo:
  Status:			%s
  StartTime:			%s
  EstablishTime:		%s
  MaturityTime:			%s
  BottomLine:			%s
  Interest:			%s
  Price:			%s
  Coupon:			%s
  Share:			%s
  TotalDeposit:			%s
  InterestInjection:			%s`,
		bi.Status, bi.StartTime.String(), bi.EstablishTime.String(), bi.MaturityTime.String(), bi.BottomLine.String(),
		bi.Interest.String(), bi.Price.String(), bi.Coupon.String(), bi.Share.String(), bi.TotalDeposit.String(), bi.InterestInjection)
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
