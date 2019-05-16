package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/types"
)

// MsgBoxInterest
type MsgBoxInterest struct {
	BoxId     string         `json:"box_id"`
	Sender    sdk.AccAddress `json:"sender"`
	Interest  sdk.Coin       `json:"interest"`
	Operation string         `json:"operation"`
}

//New MsgBoxInterest Instance
func NewMsgBoxInterest(boxId string, sender sdk.AccAddress, interest sdk.Coin, operation string) MsgBoxInterest {
	return MsgBoxInterest{boxId, sender, interest, operation}
}

// Route Implements Msg.
func (msg MsgBoxInterest) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgBoxInterest) Type() string { return types.TypeMsgBoxInterest }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgBoxInterest) ValidateBasic() sdk.Error {
	if len(msg.BoxId) == 0 {
		return errors.ErrUnknownBox("")
	}
	if err := types.CheckInterestOperation(msg.Operation); err != nil {
		return err
	}
	if msg.Interest.IsZero() || msg.Interest.IsNegative() {
		return errors.ErrAmountNotValid(msg.Interest.Denom)
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBoxInterest) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBoxInterest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgBoxInterest) String() string {
	return fmt.Sprintf("MsgBoxInterest{%s}", msg.BoxId)
}
