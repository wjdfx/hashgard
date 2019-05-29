package tests

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/types"

	"github.com/hashgard/hashgard/x/box"
	"github.com/hashgard/hashgard/x/box/msgs"

	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestLockBoxImportExportQueues(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	handler := box.NewHandler(keeper)

	boxInfo := GetLockBoxInfo()

	keeper.GetBankKeeper().AddCoins(ctx, newBoxInfo.Owner, sdk.NewCoins(boxInfo.TotalAmount.Token))

	msg := msgs.NewMsgLockBox(newBoxInfo.Owner, boxInfo)
	res := handler(ctx, msg)
	require.True(t, res.IsOK())
	var id1 string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &id1)

	keeper.GetBankKeeper().AddCoins(ctx, newBoxInfo.Owner, sdk.NewCoins(boxInfo.TotalAmount.Token))
	msg = msgs.NewMsgLockBox(newBoxInfo.Owner, boxInfo)
	res = handler(ctx, msg)
	require.True(t, res.IsOK())
	var id2 string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &id2)

	genAccs := mapp.AccountKeeper.GetAllAccounts(ctx)

	// Export the state and import it into a new Mock App
	genState := box.ExportGenesis(ctx, keeper)
	mapp2, keeper2, _, _, _, _ := getMockApp(t, 2, genState, genAccs)

	header = abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp2.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx2 := mapp2.BaseApp.NewContext(false, abci.Header{})

	boxInfo1 := keeper2.GetBox(ctx2, id1)
	require.NotNil(t, boxInfo1)
	boxInfo2 := keeper2.GetBox(ctx2, id2)
	require.NotNil(t, boxInfo2)

	require.True(t, boxInfo1.Status == types.LockBoxLocked)
	require.True(t, boxInfo2.Status == types.LockBoxLocked)

	ctx2 = ctx2.WithBlockTime(time.Unix(boxInfo.Lock.EndTime, 0))

	box.EndBlocker(ctx2, keeper2)

	boxInfo1 = keeper2.GetBox(ctx2, id1)
	require.NotNil(t, boxInfo1)
	boxInfo2 = keeper2.GetBox(ctx2, id2)
	require.NotNil(t, boxInfo2)

	require.True(t, boxInfo1.Status == types.LockBoxUnlocked)
	require.True(t, boxInfo2.Status == types.LockBoxUnlocked)
}

func TestDepositBoxImportExportQueues(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	handler := box.NewHandler(keeper)

	boxInfo := GetDepositBoxInfo()

	msg := msgs.NewMsgDepositBox(newBoxInfo.Owner, boxInfo)
	res := handler(ctx, msg)
	require.True(t, res.IsOK())
	var id1 string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &id1)

	keeper.GetBankKeeper().AddCoins(ctx, newBoxInfo.Owner, sdk.NewCoins(boxInfo.Deposit.Interest.Token))
	msgBoxInterest := msgs.NewMsgBoxInterest(id1, newBoxInfo.Owner, boxInfo.Deposit.Interest.Token, types.Injection)
	res = handler(ctx, msgBoxInterest)
	require.True(t, res.IsOK())

	ctx = ctx.WithBlockTime(time.Unix(boxInfo.Deposit.StartTime, 0))
	box.EndBlocker(ctx, keeper)

	keeper.GetBankKeeper().AddCoins(ctx, TransferAccAddr, sdk.Coins{boxInfo.TotalAmount.Token})
	msgBoxDeposit := msgs.NewMsgBoxDeposit(id1, TransferAccAddr, boxInfo.TotalAmount.Token, types.DepositTo)
	res = handler(ctx, msgBoxDeposit)
	require.True(t, res.IsOK())

	ctx = ctx.WithBlockTime(time.Unix(boxInfo.Deposit.EstablishTime, 0))
	box.EndBlocker(ctx, keeper)

	boxInfo = GetDepositBoxInfo()

	msg = msgs.NewMsgDepositBox(newBoxInfo.Owner, boxInfo)
	res = handler(ctx, msg)
	require.True(t, res.IsOK())
	var id2 string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &id2)

	keeper.GetBankKeeper().AddCoins(ctx, newBoxInfo.Owner, sdk.NewCoins(boxInfo.Deposit.Interest.Token))
	msg = msgs.NewMsgDepositBox(newBoxInfo.Owner, boxInfo)
	res = handler(ctx, msg)
	require.True(t, res.IsOK())
	var id3 string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &id3)

	msgBoxInterest = msgs.NewMsgBoxInterest(id3, newBoxInfo.Owner, boxInfo.Deposit.Interest.Token, types.Injection)
	res = handler(ctx, msgBoxInterest)
	require.True(t, res.IsOK())

	genAccs := mapp.AccountKeeper.GetAllAccounts(ctx)

	// Export the state and import it into a new Mock App
	genState := box.ExportGenesis(ctx, keeper)
	mapp2, keeper2, _, _, _, _ := getMockApp(t, 2, genState, genAccs)

	header = abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp2.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx2 := mapp2.BaseApp.NewContext(false, abci.Header{})

	boxInfo1 := keeper2.GetBox(ctx2, id1)
	require.NotNil(t, boxInfo1)
	boxInfo2 := keeper2.GetBox(ctx2, id2)
	require.NotNil(t, boxInfo2)
	boxInfo3 := keeper2.GetBox(ctx2, id3)
	require.NotNil(t, boxInfo3)

	require.True(t, boxInfo1.Status == types.DepositBoxInterest)
	require.True(t, boxInfo2.Status == types.BoxCreated)
	require.True(t, boxInfo3.Status == types.BoxCreated)

	ctx2 = ctx2.WithBlockTime(time.Unix(boxInfo.Deposit.MaturityTime, 0))
	box.EndBlocker(ctx2, keeper2)

	boxInfo1 = keeper2.GetBox(ctx2, id1)
	require.NotNil(t, boxInfo1)
	boxInfo2 = keeper2.GetBox(ctx2, id2)
	require.Nil(t, boxInfo2)
	boxInfo3 = keeper2.GetBox(ctx2, id3)
	require.NotNil(t, boxInfo3)

	require.True(t, boxInfo1.Status == types.BoxFinished)
	require.True(t, boxInfo3.Status == types.BoxDepositing)
}

func TestFutureBoxImportExportQueues(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	handler := box.NewHandler(keeper)

	boxInfo := GetFutureBoxInfo()

	msg := msgs.NewMsgFutureBox(newBoxInfo.Owner, boxInfo)
	res := handler(ctx, msg)
	require.True(t, res.IsOK())
	var id1 string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &id1)

	keeper.GetBankKeeper().AddCoins(ctx, newBoxInfo.Owner, sdk.NewCoins(boxInfo.TotalAmount.Token))
	msgDeposit := msgs.NewMsgBoxDeposit(id1, newBoxInfo.Owner, boxInfo.TotalAmount.Token, types.DepositTo)
	res = handler(ctx, msgDeposit)
	require.True(t, res.IsOK())

	boxInfo = GetFutureBoxInfo()

	msg = msgs.NewMsgFutureBox(newBoxInfo.Owner, boxInfo)
	res = handler(ctx, msg)
	require.True(t, res.IsOK())
	var id2 string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &id2)

	genAccs := mapp.AccountKeeper.GetAllAccounts(ctx)

	// Export the state and import it into a new Mock App
	genState := box.ExportGenesis(ctx, keeper)
	mapp2, keeper2, _, _, _, _ := getMockApp(t, 2, genState, genAccs)

	header = abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp2.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx2 := mapp2.BaseApp.NewContext(false, abci.Header{})

	boxInfo1 := keeper2.GetBox(ctx2, id1)
	require.NotNil(t, boxInfo1)
	boxInfo2 := keeper2.GetBox(ctx2, id2)
	require.NotNil(t, boxInfo2)

	require.True(t, boxInfo1.Status == types.BoxActived)
	require.True(t, boxInfo2.Status == types.BoxDepositing)

	ctx2 = ctx2.WithBlockTime(time.Unix(boxInfo.Future.TimeLine[len(boxInfo.Future.TimeLine)-1], 0))
	box.EndBlocker(ctx2, keeper2)

	boxInfo1 = keeper2.GetBox(ctx2, id1)
	require.NotNil(t, boxInfo1)
	boxInfo2 = keeper2.GetBox(ctx2, id2)
	require.Nil(t, boxInfo2)

	require.True(t, boxInfo1.Status == types.BoxFinished)

}
