package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/utils"
)

//Handle MsgIssueBurnAnyOff
func HandleMsgIssueBurnAnyOff(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueBurnAnyOff) sdk.Result {
	err := keeper.BurnAnyOff(ctx, msg.Sender, msg.IssueId)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.AppendIssueInfoTag(msg.IssueId),
	}
}
