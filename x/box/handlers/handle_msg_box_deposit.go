package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/types"

	"github.com/hashgard/hashgard/x/box/keeper"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/utils"
)

//Handle MsgBoxDeposit
func HandleMsgBoxDepositTo(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgBoxDepositTo) sdk.Result {
	boxInfo, err := keeper.ProcessDepositToBox(ctx, msg.Id, msg.Sender, msg.Amount, types.DepositTo)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.Id),
		Tags: utils.GetBoxTags(msg.Id, boxInfo.BoxType, msg.Sender),
	}
}

//Handle MsgBoxDepositTo
func HandleMsgBoxDepositFetch(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgBoxDepositFetch) sdk.Result {
	boxInfo, err := keeper.ProcessDepositToBox(ctx, msg.Id, msg.Sender, msg.Amount, types.Fetch)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.Id),
		Tags: utils.GetBoxTags(msg.Id, boxInfo.BoxType, msg.Sender),
	}
}
