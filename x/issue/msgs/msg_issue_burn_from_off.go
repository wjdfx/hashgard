package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/types"
)

// MsgIssueBurnFromOff to allow a registered owner
// to issue new coins.
type MsgIssueBurnFromOff struct {
	IssueId string         `json:"issue_id"`
	From    sdk.AccAddress `json:"owner"`
}

//New MsgIssueBurnFromOff Instance
func NewMsgIssueBurnFromOff(issueId string, from sdk.AccAddress) MsgIssueBurnFromOff {
	return MsgIssueBurnFromOff{issueId, from}
}

//nolint
func (ci MsgIssueBurnFromOff) GetIssueId() string {
	return ci.IssueId
}
func (ci MsgIssueBurnFromOff) SetIssueId(issueId string) {
	ci.IssueId = issueId
}
func (ci MsgIssueBurnFromOff) GetFrom() sdk.AccAddress {
	return ci.From
}
func (ci MsgIssueBurnFromOff) SetFrom(from sdk.AccAddress) {
	ci.From = from
}

// Route Implements Msg.
func (msg MsgIssueBurnFromOff) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueBurnFromOff) Type() string { return types.TypeMsgIssueBurnFromOff }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueBurnFromOff) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueBurnFromOff) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueBurnFromOff) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgIssueBurnFromOff) String() string {
	return fmt.Sprintf("MsgIssueBurnFromOff{%s}", msg.IssueId)
}
