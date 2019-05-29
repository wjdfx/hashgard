package box

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/handlers"
	"github.com/hashgard/hashgard/x/box/msgs"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/keeper"
)

// Handle all "box" type messages.
func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case msgs.MsgLockBox:
			return handlers.HandleMsgLockBox(ctx, keeper, msg)
		case msgs.MsgDepositBox:
			return handlers.HandleMsgDepositBox(ctx, keeper, msg)
		case msgs.MsgFutureBox:
			return handlers.HandleMsgFutureBox(ctx, keeper, msg)
		case msgs.MsgBoxInterestInjection:
			return handlers.HandleMsgBoxInterestInjection(ctx, keeper, msg)
		case msgs.MsgBoxInterestFetch:
			return handlers.HandleMsgBoxInterestFetch(ctx, keeper, msg)
		case msgs.MsgBoxDepositTo:
			return handlers.HandleMsgBoxDepositTo(ctx, keeper, msg)
		case msgs.MsgBoxDepositFetch:
			return handlers.HandleMsgBoxDepositFetch(ctx, keeper, msg)
		case msgs.MsgBoxWithdraw:
			return handlers.HandleMsgBoxWithdraw(ctx, keeper, msg)
		case msgs.MsgBoxDescription:
			return handlers.HandleMsgBoxDescription(ctx, keeper, msg)
		case msgs.MsgBoxDisableFeature:
			return handlers.HandleMsgBoxDisableFeature(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized gov msg type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
