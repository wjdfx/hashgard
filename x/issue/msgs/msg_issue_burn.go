package msgs

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/domain"
)

// MsgIssueBurn to allow a registered issuer
// to issue new coins.
type MsgIssueBurn struct {
	IssueId string         `json:"issue_id"`
	From    sdk.AccAddress `json:"from"`
	Amount  sdk.Int        `json:"amount"`
}

func NewMsgIssueBurn(issueId string, issuer sdk.AccAddress, amount sdk.Int) MsgIssueBurn {
	return MsgIssueBurn{issueId, issuer, amount}
}

// Route Implements Msg.
func (msg MsgIssueBurn) Route() string { return domain.RouterKey }

// Type Implements Msg.
func (msg MsgIssueBurn) Type() string { return domain.TypeMsgIssueBurn }

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
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueBurn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgIssueBurn) String() string {
	return fmt.Sprintf("MsgIssueBurn{%s - %s}", msg.IssueId, msg.Amount.String())
}
