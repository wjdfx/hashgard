package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/types"
)

// MsgBoxDeposit
type MsgBoxDeposit struct {
	BoxId     string         `json:"box_id"`
	Sender    sdk.AccAddress `json:"sender"`
	Deposit   sdk.Coin       `json:"deposit"`
	Operation string         `json:"operation"`
}

//New MsgBoxDeposit Instance
func NewMsgBoxDeposit(boxId string, sender sdk.AccAddress, deposit sdk.Coin, operation string) MsgBoxDeposit {
	return MsgBoxDeposit{boxId, sender, deposit, operation}
}

// Route Implements Msg.
func (msg MsgBoxDeposit) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgBoxDeposit) Type() string { return types.TypeMsgBoxDeposit }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgBoxDeposit) ValidateBasic() sdk.Error {
	if len(msg.BoxId) == 0 {
		return errors.ErrUnknownBox("")
	}
	if err := types.CheckDepositOperation(msg.Operation); err != nil {
		return err
	}
	if msg.Deposit.IsZero() || msg.Deposit.IsNegative() {
		return errors.ErrAmountNotValid(msg.Deposit.Denom)
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBoxDeposit) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBoxDeposit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgBoxDeposit) String() string {
	return fmt.Sprintf("MsgBoxDeposit{%s}", msg.BoxId)
}
