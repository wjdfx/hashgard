package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/keepers"
	"github.com/hashgard/hashgard/x/issue/msgs"
	issuetags "github.com/hashgard/hashgard/x/issue/tags"
)

func HandleMsgIssue(ctx sdk.Context, keeper keepers.Keeper, msg msgs.MsgIssue) sdk.Result {
	issueID, _, tags, err := keeper.AddIssue(ctx, msg)
	if err != nil {
		return err.Result()
	}
	tags = tags.AppendTag(issuetags.IssueID, issueID).
		AppendTag(issuetags.Name, msg.Name).
		AppendTag(issuetags.TotalSupply, msg.TotalSupply.String())
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(issueID),
		Tags: tags,
	}
}
