package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/keepers"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/utils"
)

func HandleMsgIssue(ctx sdk.Context, keeper keepers.Keeper, msg msgs.MsgIssue) sdk.Result {
	coinIssueInfo := msg.CoinIssueInfo
	issueID, _, tags, err := keeper.AddIssue(ctx, coinIssueInfo)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(issueID),
		Tags: tags.AppendTags(utils.AppendIssueInfoTag(issueID, coinIssueInfo)),
	}
}
