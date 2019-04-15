package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/types"
)

// MsgIssueBurnOff to allow a registered owner
// to issue new coins.
type MsgIssueBurnOff struct {
	IssueId string         `json:"issue_id"`
	From    sdk.AccAddress `json:"owner"`
}

//New MsgIssueBurnOff Instance
func NewMsgIssueBurnOff(issueId string, from sdk.AccAddress) MsgIssueBurnOff {
	return MsgIssueBurnOff{issueId, from}
}

//nolint
func (ci MsgIssueBurnOff) GetIssueId() string {
	return ci.IssueId
}
func (ci MsgIssueBurnOff) SetIssueId(issueId string) {
	ci.IssueId = issueId
}
func (ci MsgIssueBurnOff) GetFrom() sdk.AccAddress {
	return ci.From
}
func (ci MsgIssueBurnOff) SetFrom(from sdk.AccAddress) {
	ci.From = from
}

// Route Implements Msg.
func (msg MsgIssueBurnOff) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueBurnOff) Type() string { return types.TypeMsgIssueBurnOff }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueBurnOff) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueBurnOff) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueBurnOff) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgIssueBurnOff) String() string {
	return fmt.Sprintf("MsgIssueBurnOff{%s}", msg.IssueId)
}
