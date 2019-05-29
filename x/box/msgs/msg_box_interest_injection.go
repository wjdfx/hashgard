package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/types"
)

// MsgBoxInterestInjection
type MsgBoxInterestInjection struct {
	Id     string         `json:"id"`
	Sender sdk.AccAddress `json:"sender"`
	Amount sdk.Coin       `json:"amount"`
}

//New MsgBoxInterestInjection Instance
func NewMsgBoxInterestInjection(boxId string, sender sdk.AccAddress, interest sdk.Coin) MsgBoxInterestInjection {
	return MsgBoxInterestInjection{boxId, sender, interest}
}

// Route Implements Msg.
func (msg MsgBoxInterestInjection) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgBoxInterestInjection) Type() string { return types.TypeMsgBoxInterestInjection }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgBoxInterestInjection) ValidateBasic() sdk.Error {
	if len(msg.Id) == 0 {
		return errors.ErrUnknownBox("")
	}
	if msg.Amount.IsZero() || msg.Amount.IsNegative() {
		return errors.ErrAmountNotValid(msg.Amount.Denom)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBoxInterestInjection) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBoxInterestInjection) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgBoxInterestInjection) String() string {
	return fmt.Sprintf("MsgBoxInterestInjection{%s}", msg.Id)
}
