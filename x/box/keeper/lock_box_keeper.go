package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/types"
)

//Process lock box

func (keeper Keeper) ProcessLockBoxCreate(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if box.Lock.EndTime < ctx.BlockHeader().Time.Unix() {
		return errors.ErrTimeNotValid("EndTime")
	}
	if err := keeper.SendDepositedCoin(ctx, box.Owner, sdk.Coins{box.TotalAmount.Token}, box.BoxId); err != nil {
		return err
	}
	_, err := keeper.ck.AddCoins(ctx, box.Owner, sdk.Coins{sdk.NewCoin(box.BoxId, box.TotalAmount.Token.Amount)})
	if err != nil {
		return err
	}
	keeper.InsertActiveBoxQueue(ctx, box.Lock.EndTime, box.BoxId)
	box.BoxStatus = types.LockBoxLocked
	return nil
}
func (keeper Keeper) ProcessLockBoxByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if box.BoxStatus == types.LockBoxUnlocked {
		return nil
	}
	_, err := keeper.ck.SubtractCoins(ctx, box.Owner, sdk.Coins{sdk.NewCoin(box.BoxId, box.TotalAmount.Token.Amount)})
	if err != nil {
		return err
	}
	if err := keeper.FetchDepositedCoin(ctx, box.Owner, sdk.Coins{box.TotalAmount.Token}, box.BoxId); err != nil {
		return err
	}
	keeper.RemoveFromActiveBoxQueue(ctx, box.Lock.EndTime, box.BoxId)
	box.BoxStatus = types.LockBoxUnlocked
	keeper.setBox(ctx, box)
	return nil
}
