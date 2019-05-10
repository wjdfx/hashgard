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
		case msgs.MsgBox:
			return handlers.HandleMsgBox(ctx, keeper, msg)
		case msgs.MsgBoxInterest:
			return handlers.HandleMsgBoxInterest(ctx, keeper, msg)
		case msgs.MsgBoxDeposit:
			return handlers.HandleMsgBoxDeposit(ctx, keeper, msg)
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
