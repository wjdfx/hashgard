package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/types"
)

// MsgBoxInterestFetch
type MsgBoxInterestFetch struct {
	Id     string         `json:"id"`
	Sender sdk.AccAddress `json:"sender"`
	Amount sdk.Coin       `json:"amount"`
}

//New MsgBoxInterestFetch Instance
func NewMsgBoxInterestFetch(boxId string, sender sdk.AccAddress, interest sdk.Coin) MsgBoxInterestFetch {
	return MsgBoxInterestFetch{boxId, sender, interest}
}

// Route Implements Msg.
func (msg MsgBoxInterestFetch) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgBoxInterestFetch) Type() string { return types.TypeMsgBoxInterestFetch }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgBoxInterestFetch) ValidateBasic() sdk.Error {
	if len(msg.Id) == 0 {
		return errors.ErrUnknownBox("")
	}
	if msg.Amount.IsZero() || msg.Amount.IsNegative() {
		return errors.ErrAmountNotValid(msg.Amount.Denom)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBoxInterestFetch) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBoxInterestFetch) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgBoxInterestFetch) String() string {
	return fmt.Sprintf("MsgBoxInterestFetch{%s}", msg.Id)
}
