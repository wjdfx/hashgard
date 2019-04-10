package msgs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/exchange/types"
)

var _ sdk.Msg = MsgCreateOrder{}

type MsgCreateOrder struct {
	Seller 		sdk.AccAddress	`json:"seller"`
	Supply		sdk.Coin		`json:"supply"`
	Target		sdk.Coin		`json:"target"`
}

func NewMsgCreateOrder(seller sdk.AccAddress, supply sdk.Coin, target sdk.Coin) MsgCreateOrder {
	return MsgCreateOrder{
		Seller:	seller,
		Supply:	supply,
		Target:	target,
	}
}


// implement Msg interface
func (msg MsgCreateOrder) Route() string {
	return types.RouterKey
}

func (msg MsgCreateOrder) Type() string {
	return "create_order"
}

func (msg MsgCreateOrder) ValidateBasic() sdk.Error {
	if msg.Seller.Empty() {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "seller address is nil")
	}
	if msg.Supply.Amount.LTE(sdk.ZeroInt()) {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "supply amount is invalid: " + msg.Supply.String())
	}
	if msg.Target.Amount.LTE(sdk.ZeroInt()) {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "target amount is invalid: " + msg.Target.String())
	}

	return nil
}

func (msg MsgCreateOrder) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateOrder) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Seller}
}