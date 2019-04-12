package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/issue/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/types"
)

// MsgIssueDescription to allow a registered owner
// to issue new coins.
type MsgIssueDescription struct {
	IssueId     string         `json:"issue_id"`
	From        sdk.AccAddress `json:"owner"`
	Description []byte         `json:"description"`
}

//New MsgIssueDescription Instance
func NewMsgIssueDescription(issueId string, from sdk.AccAddress, description []byte) MsgIssueDescription {
	return MsgIssueDescription{issueId, from, description}
}

// Route Implements Msg.
func (msg MsgIssueDescription) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueDescription) Type() string { return types.TypeMsgIssueDescription }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueDescription) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	if len(msg.Description) > types.CoinDescriptionMaxLength {
		return errors.ErrCoinDescriptionMaxLengthNotValid()
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueDescription) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueDescription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgIssueDescription) String() string {
	return fmt.Sprintf("MsgIssueDescription{%s}", msg.IssueId)
}
