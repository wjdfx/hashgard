package msgs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgFlag interface {
	sdk.Msg

	GetIssueId() string
	SetIssueId(string)

	GetFrom() sdk.AccAddress
	SetFrom(sdk.AccAddress)
}
