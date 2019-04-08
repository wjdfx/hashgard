package handlers

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/exchange/keeper"
	"github.com/hashgard/hashgard/x/exchange/msgs"
	"github.com/hashgard/hashgard/x/exchange/tags"
)

func HandleMsgTakeOrder(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgTakeOrder) sdk.Result {
	_, _, soldOut, err := keeper.TakeOrder(ctx, msg.OrderId, msg.Buyer, msg.Value)
	if err != nil {
		return err.Result()
	}

	var status string

	if soldOut {
		status = "inactive"
	} else {
		status = "active"
	}

	resTags := sdk.NewTags(
		tags.OrderId, fmt.Sprintf("%d", msg.OrderId),
		tags.Buyer, msg.Buyer.String(),
		tags.OrderStatus, status,
	)

	return sdk.Result{
		Tags: resTags,
	}
}