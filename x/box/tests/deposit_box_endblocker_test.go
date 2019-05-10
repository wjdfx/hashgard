package tests

import (
	"fmt"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/types"

	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestDepositBoxEndBlocker(t *testing.T) {

	str := fmt.Sprintf("%saa%s%d", types.IDPreStr, strconv.FormatUint(types.BoxMaxId, 36), 999)
	fmt.Println(str)
	mapp, keeper, _, _, _, _ := getMockApp(t, 10, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	keeper.GetBankKeeper().SetSendEnabled(ctx, true)
	handler := box.NewHandler(keeper)

	inactiveQueue := keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	boxInfo := GetDepositBoxInfo()

	keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.NewCoins(boxInfo.Deposit.Interest))

	msg := msgs.NewMsgBox(boxInfo)

	res := handler(ctx, msg)
	require.True(t, res.IsOK())
	var boxID string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &boxID)
	boxInfo.BoxId = boxID

	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	msgBoxInterest := msgs.NewMsgBoxInterest(boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.Deposit.Interest.Denom, sdk.NewInt(500)), types.Injection)
	res = handler(ctx, msgBoxInterest)
	require.True(t, res.IsOK())

	msgBoxInterest = msgs.NewMsgBoxInterest(boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.Deposit.Interest.Denom, sdk.NewInt(600)), types.Injection)
	res = handler(ctx, msgBoxInterest)
	require.False(t, res.IsOK())

	msgBoxInterest = msgs.NewMsgBoxInterest(boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.Deposit.Interest.Denom, sdk.NewInt(500)), types.Injection)
	res = handler(ctx, msgBoxInterest)
	require.True(t, res.IsOK())

	newHeader := ctx.BlockHeader()
	newHeader.Time = boxInfo.Deposit.StartTime
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.True(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	box.EndBlocker(ctx, keeper)
	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	depositBox := keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, depositBox.Deposit.Status, types.DepositBoxDeposit)

	keeper.GetBankKeeper().AddCoins(ctx, TransferAccAddr, sdk.Coins{boxInfo.TotalAmount})

	msgBoxDeposit := msgs.NewMsgBoxDeposit(boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Denom, sdk.NewInt(50000)), types.DepositTo)
	res = handler(ctx, msgBoxDeposit)
	require.False(t, res.IsOK())

	msgBoxDeposit = msgs.NewMsgBoxDeposit(boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Denom, sdk.NewInt(505)), types.DepositTo)
	res = handler(ctx, msgBoxDeposit)
	require.False(t, res.IsOK())

	msgBoxDeposit = msgs.NewMsgBoxDeposit(boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Denom, sdk.NewInt(500)), types.DepositTo)
	res = handler(ctx, msgBoxDeposit)
	require.True(t, res.IsOK())

	depositBox = keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, depositBox.Deposit.Share, sdk.NewInt(5))

	newHeader = ctx.BlockHeader()
	newHeader.Time = boxInfo.Deposit.EstablishTime
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.True(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	box.EndBlocker(ctx, keeper)
	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	depositBox = keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, depositBox.Deposit.Status, types.DepositBoxInterest)
	coins := keeper.GetBankKeeper().GetCoins(ctx, TransferAccAddr)
	require.Equal(t, coins.AmountOf(boxInfo.BoxId), sdk.NewInt(500))

	newHeader = ctx.BlockHeader()
	newHeader.Time = boxInfo.Deposit.MaturityTime
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.True(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	box.EndBlocker(ctx, keeper)
	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	depositBox = keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, depositBox.Deposit.Status, types.BoxFinished)
	coins = keeper.GetBankKeeper().GetCoins(ctx, TransferAccAddr)
	require.Equal(t, coins.AmountOf(boxInfo.BoxId), sdk.ZeroInt())
	require.Equal(t, coins.AmountOf(boxInfo.TotalAmount.Denom), boxInfo.TotalAmount.Amount)
	require.Equal(t, coins.AmountOf(boxInfo.Deposit.Interest.Denom), sdk.NewInt(50))
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

	inactiveQueue := keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	boxInfo := GetDepositBoxInfo()

	keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.NewCoins(boxInfo.Deposit.Interest))

	msg := msgs.NewMsgBox(boxInfo)

	res := handler(ctx, msg)
	require.True(t, res.IsOK())
	var boxID string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &boxID)
	boxInfo.BoxId = boxID

	msgBoxInterest := msgs.NewMsgBoxInterest(boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.Deposit.Interest.Denom, sdk.NewInt(500)), types.Injection)
	res = handler(ctx, msgBoxInterest)
	require.True(t, res.IsOK())

	coins := keeper.GetBankKeeper().GetCoins(ctx, boxInfo.Owner)
	require.Equal(t, coins.AmountOf(boxInfo.Deposit.Interest.Denom), sdk.NewInt(500))

	newHeader := ctx.BlockHeader()
	newHeader.Time = boxInfo.Deposit.StartTime
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.True(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	box.EndBlocker(ctx, keeper)
	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	depositBox := keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, depositBox.Deposit.Status, types.BoxClosed)

	coins = keeper.GetBankKeeper().GetCoins(ctx, boxInfo.Owner)
	require.Equal(t, coins.AmountOf(boxInfo.Deposit.Interest.Denom), boxInfo.Deposit.Interest.Amount)
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

	inactiveQueue := keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	boxInfo := GetDepositBoxInfo()

	keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.NewCoins(boxInfo.Deposit.Interest))

	msg := msgs.NewMsgBox(boxInfo)

	res := handler(ctx, msg)
	require.True(t, res.IsOK())
	var boxID string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &boxID)
	boxInfo.BoxId = boxID

	msgBoxInterest := msgs.NewMsgBoxInterest(boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.Deposit.Interest.Denom, sdk.NewInt(1000)), types.Injection)
	res = handler(ctx, msgBoxInterest)
	require.True(t, res.IsOK())

	newHeader := ctx.BlockHeader()
	newHeader.Time = boxInfo.Deposit.StartTime
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.True(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	box.EndBlocker(ctx, keeper)
	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	keeper.GetBankKeeper().AddCoins(ctx, TransferAccAddr, sdk.NewCoins(sdk.NewCoin(boxInfo.TotalAmount.Denom, sdk.NewInt(200))))

	msgBoxDeposit := msgs.NewMsgBoxDeposit(boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Denom, sdk.NewInt(100)), types.DepositTo)
	res = handler(ctx, msgBoxDeposit)
	require.True(t, res.IsOK())

	coins := keeper.GetBankKeeper().GetCoins(ctx, TransferAccAddr)
	require.Equal(t, coins.AmountOf(boxInfo.TotalAmount.Denom), sdk.NewInt(100))

	newHeader = ctx.BlockHeader()
	newHeader.Time = boxInfo.Deposit.EstablishTime
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.True(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	box.EndBlocker(ctx, keeper)
	inactiveQueue = keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	depositBox := keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, depositBox.Deposit.Status, types.BoxClosed)

	coins = keeper.GetBankKeeper().GetCoins(ctx, TransferAccAddr)
	require.Equal(t, coins.AmountOf(boxInfo.TotalAmount.Denom), sdk.NewInt(200))
}
