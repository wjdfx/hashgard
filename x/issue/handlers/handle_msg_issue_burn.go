package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/domain"
	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/keepers"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/utils"
)

func HandleMsgIssueBurn(ctx sdk.Context, keeper keepers.Keeper, msg msgs.MsgIssueBurn) sdk.Result {
	coinIssueInfo := keeper.GetIssue(ctx, msg.IssueId)
	if coinIssueInfo == nil {
		return errors.ErrUnknownIssue(domain.DefaultCodespace, msg.IssueId).Result()
	}
	//if !coinIssueInfo.Issuer.Equals(msg.Issuer) {
	//	return errors.ErrIssuerMismatch(domain.DefaultCodespace, msg.IssueId).Result()
	//}
	_, tags, error := keeper.Burn(ctx, coinIssueInfo, msg.Amount, msg.From)
	if error != nil {
		return error.Result()
	}
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: tags.AppendTags(utils.AppendIssueInfoTag(msg.IssueId, coinIssueInfo)),
	}
}
