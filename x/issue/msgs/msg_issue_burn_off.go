package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/types"
)

// MsgIssueBurnOff to allow a registered owner
// to issue new coins.
type MsgIssueBurnOff struct {
	IssueId  string         `json:"issue_id"`
	Operator sdk.AccAddress `json:"operator"`
}

//New MsgIssueBurnOff Instance
func NewMsgIssueBurnOff(issueId string, operator sdk.AccAddress) MsgIssueBurnOff {
	return MsgIssueBurnOff{issueId, operator}
}

//nolint
func (ci MsgIssueBurnOff) GetIssueId() string {
	return ci.IssueId
}
func (ci MsgIssueBurnOff) SetIssueId(issueId string) {
	ci.IssueId = issueId
}
func (ci MsgIssueBurnOff) GetOperator() sdk.AccAddress {
	return ci.Operator
}
func (ci MsgIssueBurnOff) SetOperator(operator sdk.AccAddress) {
	ci.Operator = operator
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
	return []sdk.AccAddress{msg.Operator}
}

func (msg MsgIssueBurnOff) String() string {
	return fmt.Sprintf("MsgIssueBurnOff{%s}", msg.IssueId)
}
