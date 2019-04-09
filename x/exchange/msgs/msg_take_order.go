package msgs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/exchange/types"
)

var _ sdk.Msg = MsgTakeOrder{}

type MsgTakeOrder struct {
	OrderId 	uint64			`json:"order_id"`
	Buyer		sdk.AccAddress	`json:"buyer"`
	Value		sdk.Coin		`json:"value"`
}

func NewMsgTakeOrder(orderId uint64, buyer sdk.AccAddress, val sdk.Coin) MsgTakeOrder {
	return MsgTakeOrder{
		OrderId:	orderId,
		Buyer:		buyer,
		Value:		val,
	}
}


// implement Msg interface
func (msg MsgTakeOrder) Route() string {
	return types.RouterKey
}

func (msg MsgTakeOrder) Type() string {
	return "take_order"
}

func (msg MsgTakeOrder) ValidateBasic() sdk.Error {
	if msg.OrderId <= 0 {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "order_id is invalid")
	}
	if msg.Value.Amount.LTE(sdk.ZeroInt()) {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "value is invalid: " + msg.Value.String())
	}

	return nil
}

func (msg MsgTakeOrder) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgTakeOrder) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}