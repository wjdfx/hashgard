package handlers

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/domain"
	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/keepers"
	"github.com/hashgard/hashgard/x/issue/msgs"
	issuetags "github.com/hashgard/hashgard/x/issue/tags"
	"github.com/hashgard/hashgard/x/issue/utils"
)

func HandleMsgIssueFinishMinting(ctx sdk.Context, keeper keepers.Keeper, msg msgs.MsgIssueFinishMinting) sdk.Result {
	coinIssueInfo := keeper.GetIssue(ctx, msg.IssueId)
	if !coinIssueInfo.Issuer.Equals(msg.Issuer) {
		return errors.ErrCanNotMint(domain.DefaultCodespace, msg.IssueId).Result()
	}
	tags := utils.AppendIssueInfoTag(msg.IssueId, *coinIssueInfo)
	if coinIssueInfo.MintingFinished {
		tags = tags.AppendTag(issuetags.MintingFinished, fmt.Sprintf("%t", coinIssueInfo.MintingFinished))
		return sdk.Result{
			Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
			Tags: tags,
		}
	}
	coinIssueInfo = keeper.FinishMinting(ctx, msg.IssueId)
	tags = tags.AppendTag(issuetags.MintingFinished, fmt.Sprintf("%t", coinIssueInfo.MintingFinished))
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: tags,
	}
}
