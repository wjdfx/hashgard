package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/tags"

	"github.com/hashgard/hashgard/x/box/keeper"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/utils"
)

//Handle MsgBoxInterest
func HandleMsgBoxInterest(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgBoxInterest) sdk.Result {
	boxInfo, err := keeper.ProcessDepositBoxInterest(ctx, msg.BoxId, msg.Sender, msg.Interest, msg.Operation)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.BoxId),
		Tags: utils.GetBoxTags(msg.BoxId, boxInfo.BoxType, msg.Sender).AppendTag(tags.Operation, msg.Operation),
	}
}
