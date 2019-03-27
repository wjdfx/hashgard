package issue

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Handle all "issue" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgIssue:
			return handleMsgIssue(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized gov msg type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgIssue(ctx sdk.Context, keeper Keeper, msg MsgIssue) sdk.Result {
	_, coins, tags, err := keeper.AddIssue(ctx, msg)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Data: keeper.cdc.MustMarshalBinaryLengthPrefixed(coins),
		Tags: tags,
	}
}
