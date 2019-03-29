package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/keepers"
	"github.com/hashgard/hashgard/x/issue/msgs"
)

func HandleMsgIssue(ctx sdk.Context, keeper keepers.Keeper, msg msgs.MsgIssue) sdk.Result {
	issueID, coins, tags, err := keeper.AddIssue(ctx, msg)
	if err != nil {
		return err.Result()
	}
	tags = tags.AppendTag("issue-id", issueID)
	tags = tags.AppendTag("denom", msg.Name)
	tags = tags.AppendTag("total-supply", msg.TotalSupply.String())
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(coins),
		Tags: tags,
	}
}
