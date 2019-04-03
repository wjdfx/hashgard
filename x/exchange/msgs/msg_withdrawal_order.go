package msgs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/exchange/types"
)

var _ sdk.Msg = MsgWithdrawalOrder{}

type MsgWithdrawalOrder struct {
	OrderId 	uint64			`json:"order_id"`
	Seller		sdk.AccAddress	`json:"seller"`
}

func NewMsgWithdrawalOrder(orderId uint64, seller sdk.AccAddress) MsgWithdrawalOrder {
	return MsgWithdrawalOrder{
		OrderId:	orderId,
		Seller:		seller,
	}
}


// implement Msg interface
func (msg MsgWithdrawalOrder) Route() string {
	return types.RouterKey
}

func (msg MsgWithdrawalOrder) Type() string {
	return "withdrawal_order"
}

func (msg MsgWithdrawalOrder) ValidateBasic() sdk.Error {
	if msg.OrderId <= 0 {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "order_id is invalid")
	}
	if msg.Seller.Empty() {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "seller address is nil")
	}

	return nil
}

func (msg MsgWithdrawalOrder) GetSignBytes() []byte {
	bz := types.MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgWithdrawalOrder) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Seller}
}
