package types

import (
	"fmt"
	"time"
)

const (
	FreezeIn       = "in"
	FreezeOut      = "out"
	FreezeInAndOut = "in-out"

	UnFreezeEndTime int64 = 0
)

var FreezeType = map[string]int{FreezeIn: 1, FreezeOut: 1, FreezeInAndOut: 1}

type Freeze struct {
	OutEndTime int64 `json:"out_end_time"`
	InEndTime  int64 `json:"in_end_time"`
}

func NewFreeze(outEndTime int64, inEndTime int64) Freeze {
	return Freeze{outEndTime, inEndTime}
}

func (ci Freeze) String() string {

	return fmt.Sprintf(`Freeze:\n
	Out-end-time:			%T
	In-end-time:			%T`,
		time.Unix(ci.OutEndTime, 0), time.Unix(ci.InEndTime, 0))
}
