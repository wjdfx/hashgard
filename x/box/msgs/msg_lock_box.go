package msgs

import (
	"fmt"
	"time"

	"github.com/hashgard/hashgard/x/box/params"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/types"
)

// MsgLockBox to allow a registered boxr
// to box new coins.
type MsgLockBox struct {
	*params.BoxLockParams
}

func NewMsgLockBox(params *params.BoxLockParams) MsgLockBox {
	return MsgLockBox{params}
}

// Route Implements Msg.
func (msg MsgLockBox) Route() string { return types.RouterKey }

// Type Implements Msg.789
func (msg MsgLockBox) Type() string { return types.TypeMsgBox }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgLockBox) ValidateBasic() sdk.Error {
	if types.Lock != msg.BoxType {
		return errors.ErrUnknownBoxType()
	}

	if len(msg.Sender) == 0 {
		return sdk.ErrInvalidAddress("Sender address cannot be empty")
	}

	if msg.TotalAmount.Token.IsZero() || msg.TotalAmount.Token.Amount.IsNegative() {
		return errors.ErrAmountNotValid("Token amount")
	}
	if len(msg.Name) > types.BoxNameMaxLength {
		return errors.ErrBoxNameNotValid()
	}
	if len(msg.Description) > types.BoxDescriptionMaxLength {
		return errors.ErrBoxDescriptionMaxLengthNotValid()
	}
	if err := msg.validateBox(); err != nil {
		return err
	}
	return nil
}
func (msg MsgLockBox) validateBox() sdk.Error {

	if msg.Lock.EndTime < time.Now().Unix() {
		return errors.ErrTimeNotValid("EndTime")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgLockBox) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgLockBox) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgLockBox) String() string {
	return fmt.Sprintf("MsgLockBox{%s - %s}", "", msg.Sender.String())
}
