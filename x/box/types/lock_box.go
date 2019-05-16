package types

import (
	"fmt"
)

type LockBox struct {
	EndTime int64 `json:"end_time"`
}

//nolint
func (bi LockBox) String() string {
	return fmt.Sprintf(`LockInfo:
  EndTime:			%d`,
		bi.EndTime)
}
