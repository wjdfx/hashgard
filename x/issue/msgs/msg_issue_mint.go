package msgs

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/domain"
)

// MsgIssueMint to allow a registered issuer
// to issue new coins.
type MsgIssueMint struct {
	IssueId string         `json:"issue_id"`
	Issuer  sdk.AccAddress `json:"issuer"`
	Amount  sdk.Int        `json:"amount"`
}

func NewMsgIssueMint(issueId string, issuer sdk.AccAddress, amount sdk.Int) MsgIssueMint {
	return MsgIssueMint{issueId, issuer, amount}
}

// Route Implements Msg.
func (msg MsgIssueMint) Route() string { return domain.RouterKey }

// Type Implements Msg.
func (msg MsgIssueMint) Type() string { return domain.TypeMsgIssueMint }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueMint) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	// Cannot issue zero or negative coins
	if !msg.Amount.IsPositive() {
		return sdk.ErrInvalidCoins("Cannot mint 0 or negative coin amounts")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueMint) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Issuer}
}

func (msg MsgIssueMint) String() string {
	return fmt.Sprintf("MsgIssueMint{%s - %s}", msg.IssueId, msg.Amount.String())
}
