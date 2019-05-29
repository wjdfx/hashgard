package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/types"
)

// MsgBoxDepositFetch
type MsgBoxDepositFetch struct {
	Id     string         `json:"id"`
	Sender sdk.AccAddress `json:"sender"`
	Amount sdk.Coin       `json:"amount"`
}

//New MsgBoxDepositFetch Instance
func NewMsgBoxDepositFetch(boxId string, sender sdk.AccAddress, amount sdk.Coin) MsgBoxDepositFetch {
	return MsgBoxDepositFetch{boxId, sender, amount}
}

// Route Implements Msg.
func (msg MsgBoxDepositFetch) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgBoxDepositFetch) Type() string { return types.TypeMsgBoxDepositFetch }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgBoxDepositFetch) ValidateBasic() sdk.Error {
	if len(msg.Id) == 0 {
		return errors.ErrUnknownBox("")
	}
	if msg.Amount.IsZero() || msg.Amount.IsNegative() {
		return errors.ErrAmountNotValid(msg.Amount.Denom)
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBoxDepositFetch) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBoxDepositFetch) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgBoxDepositFetch) String() string {
	return fmt.Sprintf("MsgBoxDepositFetch{%s}", msg.Id)
}
