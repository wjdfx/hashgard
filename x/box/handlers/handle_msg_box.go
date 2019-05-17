package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/keeper"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/types"
	"github.com/hashgard/hashgard/x/box/utils"
)

//Handle MsgLockBox
func HandleMsgLockBox(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgLockBox) sdk.Result {
	box := &types.BoxInfo{
		Owner:         msg.Sender,
		Name:          msg.Name,
		BoxType:       msg.BoxType,
		TotalAmount:   msg.TotalAmount,
		Description:   msg.Description,
		TradeDisabled: true,
		Lock:          msg.Lock,
	}
	return createBox(ctx, keeper, box)
}

//Handle MsgDepositBox
func HandleMsgDepositBox(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgDepositBox) sdk.Result {
	box := &types.BoxInfo{
		Owner:         msg.Sender,
		Name:          msg.Name,
		BoxType:       msg.BoxType,
		TotalAmount:   msg.TotalAmount,
		Description:   msg.Description,
		TradeDisabled: msg.TradeDisabled,
		Deposit:       msg.Deposit,
	}
	return createBox(ctx, keeper, box)
}

//Handle MsgFutureBox
func HandleMsgFutureBox(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgFutureBox) sdk.Result {
	box := &types.BoxInfo{
		Owner:         msg.Sender,
		Name:          msg.Name,
		BoxType:       msg.BoxType,
		TotalAmount:   msg.TotalAmount,
		Description:   msg.Description,
		TradeDisabled: msg.TradeDisabled,
		Future:        msg.Future,
	}
	return createBox(ctx, keeper, box)
}
func createBox(ctx sdk.Context, keeper keeper.Keeper, box *types.BoxInfo) sdk.Result {
	err := keeper.CreateBox(ctx, box)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(box.BoxId),
		Tags: utils.GetBoxTags(box.BoxId, box.BoxType, box.Owner),
	}
}
