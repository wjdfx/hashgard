package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/types"
)

// MsgBoxDepositTo
type MsgBoxDepositTo struct {
	Id     string         `json:"id"`
	Sender sdk.AccAddress `json:"sender"`
	Amount sdk.Coin       `json:"amount"`
}

//New MsgBoxDepositTo Instance
func NewMsgBoxDepositTo(boxId string, sender sdk.AccAddress, amount sdk.Coin) MsgBoxDepositTo {
	return MsgBoxDepositTo{boxId, sender, amount}
}

// Route Implements Msg.
func (msg MsgBoxDepositTo) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgBoxDepositTo) Type() string { return types.TypeMsgBoxDepositTo }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgBoxDepositTo) ValidateBasic() sdk.Error {
	if len(msg.Id) == 0 {
		return errors.ErrUnknownBox("")
	}
	if msg.Amount.IsZero() || msg.Amount.IsNegative() {
		return errors.ErrAmountNotValid(msg.Amount.Denom)
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBoxDepositTo) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBoxDepositTo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgBoxDepositTo) String() string {
	return fmt.Sprintf("MsgBoxDepositTo{%s}", msg.Id)
}
