package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/utils"
)

//Handle MsgIssueFinishMinting
func HandleMsgIssueFinishMinting(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueFinishMinting) sdk.Result {
	err := keeper.FinishMinting(ctx, msg.Sender, msg.IssueId)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.AppendIssueInfoTag(msg.IssueId),
	}
}
