package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/types"
	"github.com/hashgard/hashgard/x/issue/utils"
)

//Handle MsgIssueMint
func HandleMsgIssueMint(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueMint) sdk.Result {
	coinIssueInfo := keeper.GetIssue(ctx, msg.IssueId)
	if coinIssueInfo == nil {
		return errors.ErrUnknownIssue(types.DefaultCodespace, msg.IssueId).Result()
	}
	if !coinIssueInfo.Issuer.Equals(msg.From) {
		return errors.ErrIssuerMismatch(types.DefaultCodespace, msg.IssueId).Result()
	}
	if coinIssueInfo.MintingFinished {
		return errors.ErrCanNotMint(types.DefaultCodespace, msg.IssueId).Result()
	}
	_, tags, err := keeper.Mint(ctx, coinIssueInfo, msg.Amount, msg.To)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: tags.AppendTags(utils.AppendIssueInfoTag(msg.IssueId, coinIssueInfo)),
	}
}
