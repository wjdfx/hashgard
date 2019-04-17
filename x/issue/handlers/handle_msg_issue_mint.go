package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/utils"
)

//Handle MsgIssueMint
func HandleMsgIssueMint(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueMint) sdk.Result {

	_, tags, err := keeper.Mint(ctx, msg.IssueId, msg.Amount, msg.Operator, msg.To)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: tags.AppendTags(utils.AppendIssueInfoTag(msg.IssueId)),
	}
}
