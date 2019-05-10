package types

import (
	"fmt"
	"time"
)

type FutureBox struct {
	Deposit   []AddressDeposit `json:"deposit"`
	TimeLine  []time.Time      `json:"time_line"`
	Receivers [][]string       `json:"receivers"`
}

//nolint
func (bi FutureBox) String() string {
	return fmt.Sprintf(`FutureInfo:
  Deposit:			%s			
  TimeLine:			%s
  Receivers:			%s`,
		bi.Deposit, bi.TimeLine, bi.Receivers)
}
