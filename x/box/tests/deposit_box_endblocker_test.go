package tests

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashgard/hashgard/x/box/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/types"
	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestDepositBoxEndBlocker(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 10, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	keeper.GetBankKeeper().SetSendEnabled(ctx, true)
	handler := box.NewHandler(keeper)

	inactiveQueue := keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	boxInfo := createDepositBox(t, ctx, keeper)

	keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.NewCoins(boxInfo.Deposit.Interest.Token))

	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	msgBoxInterest := msgs.NewMsgBoxInterest(boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.Deposit.Interest.Token.Denom,
		boxInfo.Deposit.Interest.Token.Amount.Add(sdk.NewInt(1))), types.Injection)
	res := handler(ctx, msgBoxInterest)
	require.False(t, res.IsOK())

	msgBoxInterest = msgs.NewMsgBoxInterest(boxInfo.BoxId, boxInfo.Owner, boxInfo.Deposit.Interest.Token, types.Injection)
	res = handler(ctx, msgBoxInterest)
	require.True(t, res.IsOK())

	newHeader := ctx.BlockHeader()
	newHeader.Time = time.Unix(boxInfo.Deposit.StartTime, 0)
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.True(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	box.EndBlocker(ctx, keeper)
	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	depositBox := keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, depositBox.BoxStatus, types.BoxDepositing)

	keeper.GetBankKeeper().AddCoins(ctx, TransferAccAddr, sdk.Coins{boxInfo.TotalAmount.Token})

	depositTo := boxInfo.TotalAmount.Token.Amount.Quo(sdk.NewInt(2))

	msgBoxDeposit := msgs.NewMsgBoxDeposit(boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Token.Denom,
		boxInfo.TotalAmount.Token.Amount.Add(sdk.NewInt(1))), types.DepositTo)
	res = handler(ctx, msgBoxDeposit)
	require.False(t, res.IsOK())

	msgBoxDeposit = msgs.NewMsgBoxDeposit(boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Token.Denom,
		boxInfo.Deposit.Price.Add(sdk.NewInt(1))), types.DepositTo)
	res = handler(ctx, msgBoxDeposit)
	require.False(t, res.IsOK())

	msgBoxDeposit = msgs.NewMsgBoxDeposit(boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Token.Denom,
		depositTo), types.DepositTo)
	res = handler(ctx, msgBoxDeposit)
	require.True(t, res.IsOK())

	depositBox = keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, depositBox.Deposit.Share, depositTo.Quo(boxInfo.Deposit.Price))

	newHeader = ctx.BlockHeader()
	newHeader.Time = time.Unix(boxInfo.Deposit.EstablishTime, 0)
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.True(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	box.EndBlocker(ctx, keeper)
	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	depositBox = keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, depositBox.BoxStatus, types.DepositBoxInterest)
	coins := keeper.GetBankKeeper().GetCoins(ctx, TransferAccAddr)
	require.Equal(t, coins.AmountOf(boxInfo.BoxId), depositTo.Quo(depositBox.Deposit.Price))

	newHeader = ctx.BlockHeader()
	newHeader.Time = time.Unix(boxInfo.Deposit.MaturityTime, 0)
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.True(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	box.EndBlocker(ctx, keeper)
	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	depositBox = keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, depositBox.BoxStatus, types.BoxFinished)
	coins = keeper.GetBankKeeper().GetCoins(ctx, TransferAccAddr)

	require.Equal(t, coins.AmountOf(boxInfo.BoxId), sdk.ZeroInt())
	require.Equal(t, coins.AmountOf(boxInfo.TotalAmount.Token.Denom), boxInfo.TotalAmount.Token.Amount)

	amount := depositBox.Deposit.PerCoupon.MulInt(depositTo.Quo(depositBox.Deposit.Price))
	require.Equal(t, coins.AmountOf(boxInfo.Deposit.Interest.Token.Denom),
		utils.MulMaxPrecisionByDecimal(amount, depositBox.Deposit.Interest.Decimals))
}

func TestDepositBoxNotEnoughIteratorEndBlocker(t *testing.T) {
	str := fmt.Sprintf("%saa%s%d", types.IDPreStr, strconv.FormatUint(types.BoxMaxId, 36), 999)
	fmt.Println(str)
	mapp, keeper, _, _, _, _ := getMockApp(t, 10, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	keeper.GetBankKeeper().SetSendEnabled(ctx, true)
	handler := box.NewHandler(keeper)

	inactiveQueue := keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	boxInfo := createDepositBox(t, ctx, keeper)

	keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.NewCoins(boxInfo.Deposit.Interest.Token))

	injection := boxInfo.Deposit.Interest.Token.Amount.Quo(sdk.NewInt(2))

	msgBoxInterest := msgs.NewMsgBoxInterest(boxInfo.BoxId, boxInfo.Owner,
		sdk.NewCoin(boxInfo.Deposit.Interest.Token.Denom,
			injection), types.Injection)
	res := handler(ctx, msgBoxInterest)
	require.True(t, res.IsOK())

	coins := keeper.GetBankKeeper().GetCoins(ctx, boxInfo.Owner)
	require.Equal(t, coins.AmountOf(boxInfo.Deposit.Interest.Token.Denom),
		boxInfo.Deposit.Interest.Token.Amount.Sub(injection))

	newHeader := ctx.BlockHeader()
	newHeader.Time = time.Unix(boxInfo.Deposit.StartTime, 0)
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.True(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	box.EndBlocker(ctx, keeper)
	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	depositBox := keeper.GetBox(ctx, boxInfo.BoxId)
	require.Nil(t, depositBox)
	//require.Equal(t, depositBox.BoxStatus, types.BoxClosed)

	coins = keeper.GetBankKeeper().GetCoins(ctx, boxInfo.Owner)
	require.Equal(t, coins.AmountOf(boxInfo.Deposit.Interest.Token.Denom), boxInfo.Deposit.Interest.Token.Amount)
}
func TestDepositBoxNotEnoughDepositEndBlocker(t *testing.T) {
	str := fmt.Sprintf("%saa%s%d", types.IDPreStr, strconv.FormatUint(types.BoxMaxId, 36), 999)
	fmt.Println(str)
	mapp, keeper, _, _, _, _ := getMockApp(t, 10, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	keeper.GetBankKeeper().SetSendEnabled(ctx, true)
	handler := box.NewHandler(keeper)

	inactiveQueue := keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	boxInfo := createDepositBox(t, ctx, keeper)

	keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.NewCoins(boxInfo.Deposit.Interest.Token))

	msgBoxInterest := msgs.NewMsgBoxInterest(boxInfo.BoxId, boxInfo.Owner, boxInfo.Deposit.Interest.Token, types.Injection)
	res := handler(ctx, msgBoxInterest)
	require.True(t, res.IsOK())

	newHeader := ctx.BlockHeader()
	newHeader.Time = time.Unix(boxInfo.Deposit.StartTime, 0)
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.True(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	box.EndBlocker(ctx, keeper)
	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	keeper.GetBankKeeper().AddCoins(ctx, TransferAccAddr, sdk.NewCoins(sdk.NewCoin(boxInfo.TotalAmount.Token.Denom,
		boxInfo.Deposit.BottomLine)))

	deposit := boxInfo.Deposit.BottomLine.Quo(sdk.NewInt(2))

	msgBoxDeposit := msgs.NewMsgBoxDeposit(boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Token.Denom,
		deposit), types.DepositTo)
	res = handler(ctx, msgBoxDeposit)
	require.True(t, res.IsOK())

	coins := keeper.GetBankKeeper().GetCoins(ctx, TransferAccAddr)
	require.Equal(t, coins.AmountOf(boxInfo.TotalAmount.Token.Denom), boxInfo.Deposit.BottomLine.Sub(deposit))

	newHeader = ctx.BlockHeader()
	newHeader.Time = time.Unix(boxInfo.Deposit.EstablishTime, 0)
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.True(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	box.EndBlocker(ctx, keeper)
	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	depositBox := keeper.GetBox(ctx, boxInfo.BoxId)
	require.Nil(t, depositBox)
	//require.Equal(t, depositBox.BoxStatus, types.BoxClosed)

	coins = keeper.GetBankKeeper().GetCoins(ctx, TransferAccAddr)
	require.Equal(t, coins.AmountOf(boxInfo.TotalAmount.Token.Denom), boxInfo.Deposit.BottomLine)
}
