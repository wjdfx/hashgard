package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/domain"
	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/keepers"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/utils"
)

func HandleMsgIssueMint(ctx sdk.Context, keeper keepers.Keeper, msg msgs.MsgIssueMint) sdk.Result {
	coinIssueInfo := keeper.GetIssue(ctx, msg.IssueId)
	if coinIssueInfo.MintingFinished || !coinIssueInfo.Issuer.Equals(msg.Issuer) {
		return errors.ErrCanNotMint(domain.DefaultCodespace, msg.IssueId).Result()
	}
	coinIssueInfo = keeper.Mint(ctx, msg.IssueId, msg.Amount)
	tags := utils.AppendIssueInfoTag(msg.IssueId, *coinIssueInfo)
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: tags,
	}
}
