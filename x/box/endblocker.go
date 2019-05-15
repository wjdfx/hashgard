package box

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashgard/hashgard/x/box/tags"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/types"
)

// Called every block, process inflation, update validator set
func EndBlocker(ctx sdk.Context, keeper Keeper) sdk.Tags {
	logger := ctx.Logger().With("module", "x/"+types.ModuleName)
	resTags := sdk.NewTags()

	// fetch active proposals whose voting periods have ended (are passed the block time)
	activeIterator := keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	defer activeIterator.Close()
	for ; activeIterator.Valid(); activeIterator.Next() {
		var boxID string
		var seq int

		keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(activeIterator.Value(), &boxID)
		if strings.Contains(boxID, types.KeyDelimiterString) {
			boxIDs := strings.Split(boxID, types.KeyDelimiterString)
			boxID = boxIDs[0]
			seq, _ = strconv.Atoi(boxIDs[1])
		}
		boxInfo := keeper.GetBox(ctx, boxID)
		if boxInfo == nil {
			panic(fmt.Sprintf("box %s does not exist", boxID))
			//continue
		}

		switch boxInfo.BoxType {
		case types.Lock:
			if err := keeper.ProcessLockBoxByEndBlocker(ctx, boxInfo); err != nil {
				panic(err)
			}
			logger.Info(fmt.Sprintf("lockbox %s (%s) unlocked", boxID, boxInfo.Name))
			resTags = resTags.AppendTag(tags.BoxID, boxID).AppendTag(tags.BoxType, boxInfo.GetBoxType()).AppendTag(tags.BoxStatus, types.LockBoxUnlocked)
		case types.Deposit:
			if err := keeper.ProcessDepositBoxByEndBlocker(ctx, boxInfo); err != nil {
				panic(err)
			}
			logger.Info(fmt.Sprintf("depositbox %s (%s) status:%s", boxID, boxInfo.Name, boxInfo.BoxStatus))
			resTags = resTags.AppendTag(tags.BoxID, boxID).AppendTag(tags.BoxType, boxInfo.GetBoxType()).AppendTag(tags.BoxStatus, boxInfo.BoxStatus)
		case types.Future:
			if err := keeper.ProcessFutureBoxByEndBlocker(ctx, boxInfo, seq); err != nil {
				panic(err)
			}
			logger.Info(fmt.Sprintf("futurebox %s (%s) status:%s,distributed:%s", boxID, boxInfo.Name, boxInfo.BoxStatus,
				fmt.Sprintf("%d", seq)))
			resTags = resTags.AppendTag(tags.BoxID, boxID).
				AppendTag(tags.BoxStatus, boxInfo.BoxStatus).
				AppendTag(tags.Seq, fmt.Sprintf("%d", seq))
		}

	}
	return resTags
}
