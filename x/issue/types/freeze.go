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

type IssueFreeze struct {
	OutEndTime int64 `json:"out_end_time"`
	InEndTime  int64 `json:"in_end_time"`
}

func NewIssueFreeze(outEndTime int64, inEndTime int64) IssueFreeze {
	return IssueFreeze{outEndTime, inEndTime}
}

func (ci IssueFreeze) String() string {

	return fmt.Sprintf(`Freeze:\n
	Out-end-time:			%T
	In-end-time:			%T`,
		time.Unix(ci.OutEndTime, 0), time.Unix(ci.InEndTime, 0))
}
