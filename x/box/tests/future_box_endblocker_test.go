package tests

import (
	"testing"
	"time"

	"github.com/hashgard/hashgard/x/box/utils"

	"github.com/hashgard/hashgard/x/box/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box"
	"github.com/hashgard/hashgard/x/box/msgs"

	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestFutureBoxEndBlocker(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 10, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	keeper.GetBankKeeper().SetSendEnabled(ctx, true)
	handler := box.NewHandler(keeper)

	boxInfo := createFutureBox(t, ctx, keeper)

	keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.NewCoins(boxInfo.TotalAmount.Token))

	msgDeposit := msgs.NewMsgBoxDeposit(boxInfo.BoxId, boxInfo.Owner, boxInfo.TotalAmount.Token, types.DepositTo)
	res := handler(ctx, msgDeposit)
	require.True(t, res.IsOK())

	newBoxInfo := keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, newBoxInfo.BoxStatus, types.BoxActived)

	var address sdk.AccAddress

	coins := keeper.GetDepositedCoins(ctx, boxInfo.BoxId)
	require.True(t, coins.IsEqual(sdk.NewCoins(boxInfo.TotalAmount.Token)))

	for _, v := range boxInfo.Future.Receivers {
		for j, rec := range v {
			if j == 0 {
				address, _ = sdk.AccAddressFromBech32(rec)
				coins = keeper.GetBankKeeper().GetCoins(ctx, address)
				//fmt.Println(address.String() + ":" + coins.String())
				continue
			}
			amount, _ := sdk.NewIntFromString(rec)
			boxDenom := utils.GetCoinDenomByFutureBoxSeq(boxInfo.BoxId, j)
			require.Equal(t, coins.AmountOf(boxDenom), amount)
		}
	}
	for _, v := range boxInfo.Future.TimeLine {
		newHeader := ctx.BlockHeader()
		newHeader.Time = time.Unix(v, 0)
		ctx = ctx.WithBlockHeader(newHeader)

		inactiveQueue := keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
		require.True(t, inactiveQueue.Valid())
		inactiveQueue.Close()

		box.EndBlocker(ctx, keeper)
	}

	inactiveQueue := keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()
	newBoxInfo = keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, newBoxInfo.BoxStatus, types.BoxFinished)
	require.Equal(t, newBoxInfo.Future.Distributed, newBoxInfo.Future.TimeLine)
	for _, v := range boxInfo.Future.Receivers {
		totalAmount := sdk.ZeroInt()
		for j, rec := range v {
			if j == 0 {
				address, _ = sdk.AccAddressFromBech32(rec)
				coins = keeper.GetBankKeeper().GetCoins(ctx, address)
				//fmt.Println(coins)
				continue
			}
			amount, _ := sdk.NewIntFromString(rec)
			totalAmount = totalAmount.Add(amount)
		}
		require.Equal(t, coins.AmountOf(boxInfo.TotalAmount.Token.Denom), totalAmount)
	}
	coins = keeper.GetDepositedCoins(ctx, boxInfo.BoxId)
	require.True(t, coins.IsZero())

}
