package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/tags"

	"github.com/hashgard/hashgard/x/box/keeper"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/utils"
)

//Handle MsgBoxDeposit
func HandleMsgBoxDeposit(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgBoxDeposit) sdk.Result {
	boxInfo, err := keeper.ProcessDepositToBox(ctx, msg.BoxId, msg.Sender, msg.Deposit, msg.Operation)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.BoxId),
		Tags: utils.GetBoxTags(msg.BoxId, boxInfo.BoxType, msg.Sender).AppendTag(tags.Operation, msg.Operation),
	}
}
