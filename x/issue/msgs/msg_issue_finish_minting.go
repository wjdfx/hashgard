package msgs

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/types"
)

// MsgIssueFinishMinting to allow a registered owner
// to issue new coins.
type MsgIssueFinishMinting struct {
	IssueId string         `json:"issue_id"`
	From    sdk.AccAddress `json:"owner"`
}

//New MsgIssueFinishMinting Instance
func NewMsgIssueFinishMinting(issueId string, from sdk.AccAddress) MsgIssueFinishMinting {
	return MsgIssueFinishMinting{issueId, from}
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
	return []sdk.AccAddress{msg.From}
}

func (msg MsgIssueFinishMinting) String() string {
	return fmt.Sprintf("MsgIssueFinishMinting{%s}", msg.IssueId)
}
