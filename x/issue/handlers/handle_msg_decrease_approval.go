package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue/utils"

	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
)

//Handle MsgIssueDecreaseApproval
func HandleMsgIssueDecreaseApproval(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueDecreaseApproval) sdk.Result {

	err := keeper.DecreaseApproval(ctx, msg.Sender, msg.Spender, msg.IssueId, msg.Amount)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.GetIssueTags(msg.IssueId, msg.Sender),
	}
}
