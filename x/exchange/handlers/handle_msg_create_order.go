package handlers

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/exchange/keeper"
	"github.com/hashgard/hashgard/x/exchange/msgs"
	"github.com/hashgard/hashgard/x/exchange/tags"
)

func HandleMsgCreateOrder(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgCreateOrder) sdk.Result {
	order, err := keeper.CreateOrder(ctx , msg.Seller, msg.Supply, msg.Target)

	if err != nil {
		return err.Result()
	}

	resTags := sdk.NewTags(
		tags.OrderId, fmt.Sprintf("%d", order.OrderId),
		tags.Seller, order.Seller.String(),
		tags.SupplyToken, order.Supply.Denom,
		tags.TargetToken, order.Target.Denom,
	)

	return sdk.Result{
		Tags: resTags,
	}
}