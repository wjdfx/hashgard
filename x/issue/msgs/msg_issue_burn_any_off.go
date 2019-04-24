package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/types"
)

// MsgIssueBurnAnyOff to allow a registered owner
// to issue new coins.
type MsgIssueBurnAnyOff struct {
	IssueId string         `json:"issue_id"`
	Sender  sdk.AccAddress `json:"sender"`
}

//New MsgIssueBurnAnyOff Instance
func NewMsgIssueBurnAnyOff(issueId string, sender sdk.AccAddress) MsgIssueBurnAnyOff {
	return MsgIssueBurnAnyOff{issueId, sender}
}

//nolint
func (ci MsgIssueBurnAnyOff) GetIssueId() string {
	return ci.IssueId
}
func (ci MsgIssueBurnAnyOff) SetIssueId(issueId string) {
	ci.IssueId = issueId
}
func (ci MsgIssueBurnAnyOff) GetSender() sdk.AccAddress {
	return ci.Sender
}
func (ci MsgIssueBurnAnyOff) SetSender(sender sdk.AccAddress) {
	ci.Sender = sender
}

// Route Implements Msg.
func (msg MsgIssueBurnAnyOff) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueBurnAnyOff) Type() string { return types.TypeMsgIssueBurnAnyOff }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueBurnAnyOff) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueBurnAnyOff) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueBurnAnyOff) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgIssueBurnAnyOff) String() string {
	return fmt.Sprintf("MsgIssueBurnAnyOff{%s}", msg.IssueId)
}
