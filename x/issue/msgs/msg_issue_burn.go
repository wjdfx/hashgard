package msgs

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/types"
)

// MsgIssueBurn to allow a registered owner
// to issue new coins.
type MsgIssueBurn struct {
	IssueId string         `json:"issue_id"`
	From    sdk.AccAddress `json:"from"`
	Amount  sdk.Int        `json:"amount"`
}

//New CreateMsgIssue Instance
func NewMsgIssueBurn(issueId string, owner sdk.AccAddress, amount sdk.Int) MsgIssueBurn {
	return MsgIssueBurn{issueId, owner, amount}
}

// Route Implements Msg.
func (msg MsgIssueBurn) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueBurn) Type() string { return types.TypeMsgIssueBurn }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueBurn) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	// Cannot issue zero or negative coins
	if !msg.Amount.IsPositive() {
		return sdk.ErrInvalidCoins("Cannot Burn 0 or negative coin amounts")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueBurn) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueBurn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgIssueBurn) String() string {
	return fmt.Sprintf("MsgIssueBurn{%s - %s}", msg.IssueId, msg.Amount.String())
}
