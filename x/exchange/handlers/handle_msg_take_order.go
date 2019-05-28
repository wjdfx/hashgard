package handlers

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/exchange/keeper"
	"github.com/hashgard/hashgard/x/exchange/msgs"
	"github.com/hashgard/hashgard/x/exchange/tags"
)

func HandleMsgTakeOrder(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgTakeOrder) sdk.Result {
	supplyTurnover, targetTurnover, soldOut, err := keeper.TakeOrder(ctx, msg.OrderId, msg.Buyer, msg.Value)
	if err != nil {
		return err.Result()
	}

	resTags := sdk.NewTags(
		tags.Category, tags.TxCategory,
		tags.OrderId, fmt.Sprintf("%d", msg.OrderId),
		tags.Sender, msg.Buyer.String(),
		tags.SupplyTurnover, supplyTurnover,
		tags.TargetTurnover, targetTurnover,
	)

	if soldOut {
		resTags = resTags.AppendTag(tags.OrderStatus, "inactive")
	}

	return sdk.Result{
		Tags: resTags,
	}
}
