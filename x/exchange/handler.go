package exchange

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/exchange/handlers"
	"github.com/hashgard/hashgard/x/exchange/keeper"
	"github.com/hashgard/hashgard/x/exchange/msgs"
)

func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case msgs.MsgCreateOrder:
			return handlers.HandleMsgCreateOrder(ctx, keeper, msg)
		case msgs.MsgTakeOrder:
			return handlers.HandleMsgTakeOrder(ctx, keeper, msg)
		case msgs.MsgWithdrawalOrder:
			return handlers.HandleMsgWithdrawalOrder(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized exchange msg type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
