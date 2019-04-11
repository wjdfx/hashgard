package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/utils"
)

//Handle MsgIssueBurn
func HandleMsgIssueBurn(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueBurn) sdk.Result {
	coinIssueInfo := keeper.GetIssue(ctx, msg.IssueId)
	if coinIssueInfo == nil {
		return errors.ErrUnknownIssue(msg.IssueId).Result()
	}
	_, tags, err := keeper.Burn(ctx, coinIssueInfo, msg.Amount, msg.From)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: tags.AppendTags(utils.AppendIssueInfoTag(msg.IssueId, coinIssueInfo)),
	}
}
