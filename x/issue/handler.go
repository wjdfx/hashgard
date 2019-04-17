package issue

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/handlers"
	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
)

// Handle all "issue" type messages.
func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case msgs.MsgIssue:
			return handlers.HandleMsgIssue(ctx, keeper, msg)
		case msgs.MsgIssueTransferOwnership:
			return handlers.HandleMsgIssueTransferOwnership(ctx, keeper, msg)
		case msgs.MsgIssueDescription:
			return handlers.HandleMsgIssueDescription(ctx, keeper, msg)
		case msgs.MsgIssueMint:
			return handlers.HandleMsgIssueMint(ctx, keeper, msg)
		case msgs.MsgIssueBurn:
			return handlers.HandleMsgIssueBurn(ctx, keeper, msg)
		case msgs.MsgIssueBurnFrom:
			return handlers.HandleMsgIssueBurnFrom(ctx, keeper, msg)
		case msgs.MsgIssueBurnAny:
			return handlers.HandleMsgIssueBurnAny(ctx, keeper, msg)
		case msgs.MsgIssueBurnOff:
			return handlers.HandleMsgIssueBurnOff(ctx, keeper, msg)
		case msgs.MsgIssueBurnFromOff:
			return handlers.HandleMsgIssueBurnFromOff(ctx, keeper, msg)
		case msgs.MsgIssueBurnAnyOff:
			return handlers.HandleMsgIssueBurnAnyOff(ctx, keeper, msg)
		case msgs.MsgIssueFinishMinting:
			return handlers.HandleMsgIssueFinishMinting(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized gov msg type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
