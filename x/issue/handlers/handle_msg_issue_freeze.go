package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/tags"
	"github.com/hashgard/hashgard/x/issue/utils"
)

//Handle MsgIssueFreeze
func HandleMsgIssueFreeze(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueFreeze) sdk.Result {

	err := keeper.Freeze(ctx, msg.GetIssueId(), msg.GetSender(), msg.GetAccAddress(), msg.GetFreezeType(), msg.GetEndTime())

	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.GetIssueTags(msg.IssueId, msg.Sender).AppendTag(tags.FreezeType, msg.GetFreezeType()),
	}
}
