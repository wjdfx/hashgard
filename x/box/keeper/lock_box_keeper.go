package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/types"
)

//Process lock box

func (keeper Keeper) ProcessLockBoxCreate(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	_, err := keeper.ck.SubtractCoins(ctx, box.Owner, sdk.Coins{box.TotalAmount})
	if err != nil {
		return err
	}
	_, err = keeper.ck.AddCoins(ctx, box.Owner, sdk.Coins{sdk.NewCoin(box.BoxId, box.TotalAmount.Amount)})
	if err != nil {
		return err
	}
	keeper.InsertActiveBoxQueue(ctx, box.Lock.EndTime, box.BoxId)
	box.Lock.Status = types.LockBoxLocked
	return nil
}
func (keeper Keeper) ProcessLockBoxByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if box.Lock.Status == types.LockBoxUnlocked {
		return nil
	}
	_, err := keeper.ck.SubtractCoins(ctx, box.Owner, sdk.Coins{sdk.NewCoin(box.BoxId, box.TotalAmount.Amount)})
	if err != nil {
		return err
	}
	_, err = keeper.ck.AddCoins(ctx, box.Owner, sdk.Coins{box.TotalAmount})
	if err != nil {
		return err
	}
	keeper.RemoveFromActiveBoxQueue(ctx, box.Lock.EndTime, box.BoxId)
	box.Lock.Status = types.LockBoxUnlocked
	keeper.setBox(ctx, box)
	return nil
}
