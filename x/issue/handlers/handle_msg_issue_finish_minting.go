package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/utils"
)

//Handle MsgIssueFinishMinting
func HandleMsgIssueFinishMinting(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueFinishMinting) sdk.Result {
	coinIssueInfo := keeper.GetIssue(ctx, msg.IssueId)
	if coinIssueInfo == nil {
		return errors.ErrUnknownIssue(msg.IssueId).Result()
	}
	if !coinIssueInfo.Owner.Equals(msg.From) {
		return errors.ErrIssuerMismatch(msg.IssueId).Result()
	}
	if coinIssueInfo.MintingFinished {
		return sdk.Result{
			Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
			Tags: utils.AppendIssueInfoTag(msg.IssueId, coinIssueInfo),
		}
	}
	coinIssueInfo = keeper.FinishMinting(ctx, msg.IssueId)
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.AppendIssueInfoTag(msg.IssueId, coinIssueInfo),
	}
}
