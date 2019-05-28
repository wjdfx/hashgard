package handlers

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/exchange/keeper"
	"github.com/hashgard/hashgard/x/exchange/msgs"
	"github.com/hashgard/hashgard/x/exchange/tags"
)

func HandleMsgCreateOrder(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgCreateOrder) sdk.Result {
	order, err := keeper.CreateOrder(ctx, msg.Seller, msg.Supply, msg.Target)

	if err != nil {
		return err.Result()
	}

	resTags := sdk.NewTags(
		tags.Category, tags.TxCategory,
		tags.OrderId, fmt.Sprintf("%d", order.OrderId),
		tags.Sender, order.Seller.String(),
		tags.SupplyDenom, order.Supply.Denom,
		tags.TargetDenom, order.Target.Denom,
	)

	return sdk.Result{
		Tags: resTags,
	}
}
