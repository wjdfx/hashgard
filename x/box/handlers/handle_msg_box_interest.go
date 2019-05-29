package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/keeper"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/utils"
)

//Handle MsgBoxInterestInjection
func HandleMsgBoxInterestInjection(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgBoxInterestInjection) sdk.Result {
	boxInfo, err := keeper.InjectionDepositBoxInterest(ctx, msg.Id, msg.Sender, msg.Amount)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.Id),
		Tags: utils.GetBoxTags(msg.Id, boxInfo.BoxType, msg.Sender),
	}
}

//Handle MsgBoxInterestFetch
func HandleMsgBoxInterestFetch(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgBoxInterestFetch) sdk.Result {
	boxInfo, err := keeper.FetchInterestFromDepositBox(ctx, msg.Id, msg.Sender, msg.Amount)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.Id),
		Tags: utils.GetBoxTags(msg.Id, boxInfo.BoxType, msg.Sender),
	}
}
