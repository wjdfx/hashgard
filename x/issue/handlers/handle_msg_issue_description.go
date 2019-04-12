package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/utils"
)

//Handle MsgIssueDescription
func HandleMsgIssueDescription(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueDescription) sdk.Result {
	coinIssueInfo := keeper.GetIssue(ctx, msg.IssueId)
	if coinIssueInfo == nil {
		return errors.ErrUnknownIssue(msg.IssueId).Result()
	}
	if !coinIssueInfo.Owner.Equals(msg.From) {
		return errors.ErrIssuerMismatch(msg.IssueId).Result()
	}

	err := keeper.SetIssueDescription(ctx, msg.IssueId, msg.From, msg.Description)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.AppendIssueInfoTag(msg.IssueId),
	}
}
