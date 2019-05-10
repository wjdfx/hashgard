package types

import (
	"fmt"
	"time"
)

type LockBox struct {
	Status  string    `json:"status"`
	EndTime time.Time `json:"end_time"`
}

//nolint
func (bi LockBox) String() string {
	return fmt.Sprintf(`LockInfo:
  Status:			%s
  EndTime:			%s`,
		bi.Status, bi.EndTime)
}
