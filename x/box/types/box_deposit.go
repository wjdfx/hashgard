package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BoxDeposit struct {
	Amount   sdk.Int `json:"amount"`
	Interest sdk.Int `json:"interest"`
}

func NewZeroBoxDeposit() *BoxDeposit {
	return &BoxDeposit{Amount: sdk.ZeroInt(), Interest: sdk.ZeroInt()}
}
func NewBoxDeposit(amount sdk.Int) *BoxDeposit {
	return &BoxDeposit{Amount: amount, Interest: sdk.ZeroInt()}
}

//nolint
func (bi BoxDeposit) String() string {
	return fmt.Sprintf(`
  Amount:			%s
  Interest:			%s`,
		bi.Amount.String(), bi.Interest.String())
}
