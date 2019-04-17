package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/utils"
)

//Handle leMsgIssueBurnFromOff
func HandleMsgIssueBurnFromOff(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueBurnFromOff) sdk.Result {
	err := keeper.BurnFromOff(ctx, msg.Operator, msg.IssueId)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.AppendIssueInfoTag(msg.IssueId),
	}
}
