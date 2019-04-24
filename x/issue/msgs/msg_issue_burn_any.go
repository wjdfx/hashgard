package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/types"
)

// MsgIssueBurnAny to allow a registered owner
// to issue new coins.
type MsgIssueBurnAny struct {
	IssueId string         `json:"issue_id"`
	Sender  sdk.AccAddress `json:"sender"`
	From    sdk.AccAddress `json:"from"`
	Amount  sdk.Int        `json:"amount"`
}

//New CreateMsgIssue Instance
func NewMsgIssueBurnAny(issueId string, sender sdk.AccAddress, from sdk.AccAddress, amount sdk.Int) MsgIssueBurnAny {
	return MsgIssueBurnAny{issueId, sender, from, amount}
}

// Route Implements Msg.
func (msg MsgIssueBurnAny) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueBurnAny) Type() string { return types.TypeMsgIssueBurnAny }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueBurnAny) ValidateBasic() sdk.Error {
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
func (msg MsgIssueBurnAny) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueBurnAny) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgIssueBurnAny) String() string {
	return fmt.Sprintf("MsgIssueBurnAny{%s - %s}", msg.IssueId, msg.Amount.String())
}
