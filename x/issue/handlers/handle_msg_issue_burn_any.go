package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/utils"
)

//Handle MsgIssueBurnAny
func HandleMsgIssueBurnAny(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueBurnAny) sdk.Result {

	_, tags, err := keeper.BurnFrom(ctx, msg.IssueId, msg.Amount, msg.Sender, msg.From)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: tags.AppendTags(utils.AppendIssueInfoTag(msg.IssueId)),
	}
}
