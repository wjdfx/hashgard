package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/utils"
)

//Handle MsgIssue
func HandleMsgIssue(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssue) sdk.Result {
	coinIssueInfo := msg.CoinIssueInfo
	_, tags, err := keeper.AddIssue(ctx, coinIssueInfo)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: tags.AppendTags(utils.AppendIssueInfoTag(coinIssueInfo.IssueId)),
	}
}
