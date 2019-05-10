package box

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/tags"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/types"
)

// Called every block, process inflation, update validator set
func EndBlocker(ctx sdk.Context, keeper Keeper) sdk.Tags {
	logger := ctx.Logger().With("module", "x/"+types.ModuleName)
	resTags := sdk.NewTags()

	// fetch active proposals whose voting periods have ended (are passed the block time)
	activeIterator := keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	defer activeIterator.Close()
	for ; activeIterator.Valid(); activeIterator.Next() {
		var boxID string

		keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(activeIterator.Value(), &boxID)
		boxInfo := keeper.GetBox(ctx, boxID)
		if boxInfo == nil {
			panic(fmt.Sprintf("box %s does not exist", boxID))
		}

		switch boxInfo.BoxType {
		case types.Lock:
			if err := keeper.ProcessLockBoxByEndBlocker(ctx, boxInfo); err != nil {
				panic(err)
			}
			logger.Info(fmt.Sprintf("lockbox %s (%s) unlocked", boxID, boxInfo.Name))
			resTags = resTags.AppendTag(tags.BoxID, boxID).AppendTag(tags.BoxType, boxInfo.GetBoxType()).AppendTag(tags.Status, types.LockBoxUnlocked)
		case types.Deposit:
			if err := keeper.ProcessDepositBoxByEndBlocker(ctx, boxInfo); err != nil {
				panic(err)
			}
			logger.Info(fmt.Sprintf("depositbox %s (%s) status:%s", boxID, boxInfo.Name, boxInfo.Deposit.Status))
			resTags = resTags.AppendTag(tags.BoxID, boxID).AppendTag(tags.BoxType, boxInfo.GetBoxType()).AppendTag(tags.Status, boxInfo.Deposit.Status)
		case types.Future:
			if err := keeper.ProcessFutureBoxByEndBlocker(ctx, boxInfo); err != nil {
				panic(err)
			}
			//TODO
			logger.Info(fmt.Sprintf("futurebox %s (%s):", boxID, boxInfo.Name))
			//resTags = resTags.AppendTag(tags.BoxID, boxID).AppendTag(tags.Status, boxInfo.Deposit.Status)
		}

	}
	return resTags
}
