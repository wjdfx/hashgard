package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/tags"
)

//Handle MsgIssueApprove
func HandleMsgIssueApprove(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueApprove) sdk.Result {

	err := keeper.Approve(ctx, msg.Sender, msg.Spender, msg.IssueId, msg.Amount)
	if err != nil {
		return err.Result()
	}

	resTags := sdk.NewTags(
		tags.Category, tags.TxCategory,
		tags.IssueID, msg.IssueId,
		tags.Sender, msg.Operator.String(),
	)

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: resTags,
	}
}
