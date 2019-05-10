package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/keeper"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/utils"
)

//Handle MsgBox
func HandleMsgBox(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgBox) sdk.Result {

	boxInfo := msg.BoxInfo
	err := keeper.CreateBox(ctx, boxInfo)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(boxInfo.BoxId),
		Tags: utils.GetBoxTags(boxInfo.BoxId, boxInfo.BoxType, boxInfo.Owner),
	}
}
