package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/params"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/types"
)

// MsgFutureBox to allow a registered boxr
// to box new coins.
type MsgFutureBox struct {
	*params.BoxFutureParams
}

func NewMsgFutureBox(params *params.BoxFutureParams) MsgFutureBox {
	return MsgFutureBox{params}
}

// Route Implements Msg.
func (msg MsgFutureBox) Route() string { return types.RouterKey }

// Type Implements Msg.789
func (msg MsgFutureBox) Type() string { return types.TypeMsgBoxCreateFuture }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgFutureBox) ValidateBasic() sdk.Error {
	if types.Future != msg.BoxType {
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

func (msg MsgFutureBox) validateBox() sdk.Error {
	if msg.Future.TimeLine == nil || msg.Future.Receivers == nil ||
		len(msg.Future.TimeLine) == 0 || len(msg.Future.Receivers) == 0 {
		return errors.ErrNotSupportOperation()
	}
	if len(msg.Future.TimeLine) > types.BoxMaxInstalment {
		return errors.ErrNotEnoughAmount()
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgFutureBox) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgFutureBox) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgFutureBox) String() string {
	return fmt.Sprintf("MsgFutureBox{%s - %s}", "", msg.Sender.String())
}
