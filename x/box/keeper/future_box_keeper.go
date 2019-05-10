package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/types"
)

//Process Future box

func (keeper Keeper) ProcessFutureBoxCreate(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	//TODO
	for _, v := range box.Future.TimeLine {
		keeper.InsertActiveBoxQueue(ctx, v, box.BoxId)
	}

	return nil
}

func (keeper Keeper) ProcessFutureBoxByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	//TODO
	keeper.InsertActiveBoxQueue(ctx, box.Deposit.StartTime, box.BoxId)
	return nil
}
