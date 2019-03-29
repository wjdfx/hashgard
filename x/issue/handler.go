package issue

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/handlers"
	"github.com/hashgard/hashgard/x/issue/keepers"
	"github.com/hashgard/hashgard/x/issue/msgs"
)

// Handle all "issue" type messages.
func NewHandler(keeper keepers.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case msgs.MsgIssue:
			return handlers.HandleMsgIssue(ctx, keeper, msg)
		case msgs.MsgIssueMint:
			return handlers.HandleMsgIssueMint(ctx, keeper, msg)
		case msgs.MsgIssueBurn:
			return handlers.HandleMsgIssueBurn(ctx, keeper, msg)
		case msgs.MsgIssueFinishMinting:
			return handlers.HandleMsgIssueFinishMinting(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized gov msg type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
