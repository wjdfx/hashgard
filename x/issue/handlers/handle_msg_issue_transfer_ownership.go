package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/tags"
)

//Handle MsgIssueMint
func HandleMsgIssueTransferOwnership(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueTransferOwnership) sdk.Result {

	err := keeper.TransferOwnership(ctx, msg.IssueId, msg.Operator, msg.To)
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
