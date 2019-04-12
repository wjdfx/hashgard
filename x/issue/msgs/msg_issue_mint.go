package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/utils"

	"github.com/hashgard/hashgard/x/issue/types"
)

// MsgIssueMint to allow a registered issuer
// to issue new coins.
type MsgIssueMint struct {
	IssueId  string         `json:"issue_id"`
	From     sdk.AccAddress `json:"from"`
	Amount   sdk.Int        `json:"amount"`
	Decimals uint           `json:"decimals"`
	To       sdk.AccAddress `json:"to"`
}

//New MsgIssueMint Instance
func NewMsgIssueMint(issueId string, from sdk.AccAddress, amount sdk.Int, decimals uint, to sdk.AccAddress) MsgIssueMint {
	return MsgIssueMint{issueId, from, amount, decimals, to}
}

// Route Implements Msg.
func (msg MsgIssueMint) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueMint) Type() string { return types.TypeMsgIssueMint }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueMint) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	// Cannot issue zero or negative coins
	if !msg.Amount.IsPositive() {
		return sdk.ErrInvalidCoins("Cannot mint 0 or negative coin amounts")
	}
	if utils.QuoDecimals(msg.Amount, msg.Decimals).GT(types.CoinMaxTotalSupply) {
		return errors.ErrCoinTotalSupplyMaxValueNotValid()
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueMint) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgIssueMint) String() string {
	return fmt.Sprintf("MsgIssueMint{%s - %s}", msg.IssueId, msg.Amount.String())
}
