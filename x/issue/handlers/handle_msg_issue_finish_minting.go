package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/domain"
	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/keepers"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/utils"
)

func HandleMsgIssueFinishMinting(ctx sdk.Context, keeper keepers.Keeper, msg msgs.MsgIssueFinishMinting) sdk.Result {
	coinIssueInfo := keeper.GetIssue(ctx, msg.IssueId)
	if coinIssueInfo == nil {
		return errors.ErrUnknownIssue(domain.DefaultCodespace, msg.IssueId).Result()
	}
	if !coinIssueInfo.Issuer.Equals(msg.From) {
		return errors.ErrIssuerMismatch(domain.DefaultCodespace, msg.IssueId).Result()
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
