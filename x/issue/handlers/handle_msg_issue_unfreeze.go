package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/tags"
	"github.com/hashgard/hashgard/x/issue/utils"
)

//Handle MsgIssueUnFreeze
func HandleMsgIssueUnFreeze(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueUnFreeze) sdk.Result {

	err := keeper.UnFreeze(ctx, msg.GetIssueId(), msg.GetSender(), msg.GetAccAddress(), msg.GetFreezeType())

	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.AppendIssueInfoTag(msg.IssueId).AppendTag(tags.FreezeType, msg.GetFreezeType()),
	}
}
