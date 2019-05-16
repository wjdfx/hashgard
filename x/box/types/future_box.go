package types

import (
	"fmt"
)

type FutureBox struct {
	MiniMultiple uint             `json:"mini_multiple"`
	Deposits     []AddressDeposit `json:"deposits"`
	TimeLine     []int64          `json:"time"`
	Receivers    [][]string       `json:"receivers"`
	Distributed  []int64          `json:"distributed"`
}

//nolint
func (bi FutureBox) String() string {
	return fmt.Sprintf(`FutureInfo:
  MiniMultiple:			%d
  Deposit:			%s			
  TimeLine:			%d
  Distributed:			%d
  Receivers:			%s`,
		bi.MiniMultiple, bi.Deposits, bi.Distributed, bi.TimeLine, bi.Receivers)
}
