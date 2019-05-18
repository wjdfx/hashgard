package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/keeper"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/tags"
	"github.com/hashgard/hashgard/x/box/utils"
)

//Handle MsgBoxDisableFeature
func HandleMsgBoxDisableFeature(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgBoxDisableFeature) sdk.Result {
	boxInfo, err := keeper.DisableFeature(ctx, msg.Sender, msg.BoxId, msg.Feature)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.BoxId),
		Tags: utils.GetBoxTags(msg.BoxId, boxInfo.BoxType, msg.Sender).AppendTag(tags.Feature, msg.GetFeature()),
	}
}
