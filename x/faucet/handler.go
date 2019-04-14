package faucet

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgFaucetSend:
			return handleMsgFaucetSend(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized gov msg type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgFaucetSend(ctx sdk.Context, keeper Keeper, msg MsgFaucetSend) sdk.Result {
	receiverBalance := keeper.bankKeeper.GetCoins(ctx, msg.Receiver)
	limitCoins := keeper.GetLimitCoins(ctx)
	if receiverBalance.IsAllGTE(limitCoins) {
		return sdk.NewError(keeper.codespace, CodeHaveTooMany, fmt.Sprintf("your balance is more than the faucet limit(%s)", limitCoins.String())).Result()
	}

	faucetOrigin := keeper.GetFaucetOrigin(ctx)
	if faucetOrigin.Empty() {
		return sdk.NewError(keeper.codespace, CodeInvalidGenesis, "faucet origin never set").Result()
	}

	sendCoins := keeper.GetSendCoins(ctx)
	tags, err := keeper.bankKeeper.SendCoins(ctx, faucetOrigin, msg.Receiver, sendCoins)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Tags: tags,
	}
}