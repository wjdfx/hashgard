package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/types"
)

// MsgIssueBurnAnyOff to allow a registered owner
// to issue new coins.
type MsgIssueBurnAnyOff struct {
	IssueId  string         `json:"issue_id"`
	Operator sdk.AccAddress `json:"operator"`
}

//New MsgIssueBurnAnyOff Instance
func NewMsgIssueBurnAnyOff(issueId string, operator sdk.AccAddress) MsgIssueBurnAnyOff {
	return MsgIssueBurnAnyOff{issueId, operator}
}

//nolint
func (ci MsgIssueBurnAnyOff) GetIssueId() string {
	return ci.IssueId
}
func (ci MsgIssueBurnAnyOff) SetIssueId(issueId string) {
	ci.IssueId = issueId
}
func (ci MsgIssueBurnAnyOff) GetOperator() sdk.AccAddress {
	return ci.Operator
}
func (ci MsgIssueBurnAnyOff) SetOperator(operator sdk.AccAddress) {
	ci.Operator = operator
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
	return []sdk.AccAddress{msg.Operator}
}

func (msg MsgIssueBurnAnyOff) String() string {
	return fmt.Sprintf("MsgIssueBurnAnyOff{%s}", msg.IssueId)
}
