package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/tags"
)

//Handle MsgIssueBurnFrom
func HandleMsgIssueBurnFrom(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueBurnFrom) sdk.Result {

	_, err := keeper.BurnFrom(ctx, msg.IssueId, msg.Amount, msg.Operator, msg.From)
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
