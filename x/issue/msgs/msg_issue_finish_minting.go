package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/types"
)

// MsgIssueFinishMinting to allow a registered owner
// to issue new coins.
type MsgIssueFinishMinting struct {
	IssueId  string         `json:"issue_id"`
	Operator sdk.AccAddress `json:"operator"`
}

//New MsgIssueFinishMinting Instance
func NewMsgIssueFinishMinting(issueId string, operator sdk.AccAddress) MsgIssueFinishMinting {
	return MsgIssueFinishMinting{issueId, operator}
}

func (ci MsgIssueFinishMinting) GetIssueId() string {
	return ci.IssueId
}
func (ci MsgIssueFinishMinting) SetIssueId(issueId string) {
	ci.IssueId = issueId
}
func (ci MsgIssueFinishMinting) GetOperator() sdk.AccAddress {
	return ci.Operator
}
func (ci MsgIssueFinishMinting) SetOperator(operator sdk.AccAddress) {
	ci.Operator = operator
}

// Route Implements Msg.
func (msg MsgIssueFinishMinting) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueFinishMinting) Type() string { return types.TypeMsgIssueFinishMinting }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueFinishMinting) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueFinishMinting) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueFinishMinting) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Operator}
}

func (msg MsgIssueFinishMinting) String() string {
	return fmt.Sprintf("MsgIssueFinishMinting{%s}", msg.IssueId)
}
