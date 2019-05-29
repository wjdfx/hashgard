package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/keeper"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/utils"
)

//Handle MsgBoxDescription
func HandleMsgBoxDescription(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgBoxDescription) sdk.Result {
	boxInfo, err := keeper.SetBoxDescription(ctx, msg.Id, msg.Sender, msg.Description)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.Id),
		Tags: utils.GetBoxTags(msg.Id, boxInfo.BoxType, msg.Sender),
	}
}
