package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type FutureBox struct {
	Deposits        []AddressDeposit `json:"deposits"`
	TimeLine        []int64          `json:"time"`
	Receivers       [][]string       `json:"receivers"`
	TotalWithdrawal sdk.Int          `json:"total_withdrawal"`
}

//nolint
func (bi FutureBox) String() string {
	return fmt.Sprintf(`FutureInfo:
  Deposit:			%s			
  TimeLine:			%d
  Receivers:			%s
  TotalWithdrawal:			%s`,
		bi.Deposits, bi.TimeLine, bi.Receivers, bi.TotalWithdrawal.String())
}
