package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/tags"
	"github.com/hashgard/hashgard/x/issue/types"

	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/utils"
)

//Handle MsgIssueDisableFeature
func HandleMsgIssueDisableFeature(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueDisableFeature) sdk.Result {
	var err sdk.Error
	switch msg.GetFeature() {

	case types.BurnOwner:
		err = keeper.DisableBurnOwner(ctx, msg.Sender, msg.IssueId)
	case types.BurnHolder:
		err = keeper.DisableBurnHolder(ctx, msg.Sender, msg.IssueId)
	case types.BurnFrom:
		err = keeper.DisableBurnFrom(ctx, msg.Sender, msg.IssueId)
	case types.Minting:
		err = keeper.FinishMinting(ctx, msg.Sender, msg.IssueId)
	default:
		err = errors.ErrUnknownFeatures()
	}

	if err != nil {
		return err.Result()

	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.AppendIssueInfoTag(msg.IssueId).AppendTag(tags.Feature, msg.GetFeature()),
	}
}
