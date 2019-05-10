package tests

import (
	"testing"

	"github.com/hashgard/hashgard/x/box/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestDepositBoxAdd(t *testing.T) {

	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	boxInfo := GetDepositBoxInfo()

	err := keeper.CreateBox(ctx, boxInfo)
	require.Nil(t, err)
	box := keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, box.Name, boxInfo.Name)

}
func TestDepositBoxFetchInterest(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	boxInfo := GetDepositBoxInfo()

	err := keeper.CreateBox(ctx, boxInfo)
	require.Nil(t, err)

	keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.NewCoins(boxInfo.Deposit.Interest))
	keeper.GetBankKeeper().AddCoins(ctx, TransferAccAddr, sdk.NewCoins(boxInfo.Deposit.Interest))

	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin("test", sdk.NewInt(1000)), types.Injection)
	require.Error(t, err)
	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.Deposit.Interest.Denom, sdk.NewInt(10000)), types.Injection)
	require.Error(t, err)
	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.Deposit.Interest.Denom, sdk.NewInt(500)), types.Injection)
	require.Nil(t, err)
	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.Deposit.Interest.Denom, sdk.NewInt(500)), types.Injection)
	require.Nil(t, err)

	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin("test", sdk.NewInt(1000)), types.Fetch)
	require.Error(t, err)
	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.Deposit.Interest.Denom, sdk.NewInt(10000)), types.Fetch)
	require.Error(t, err)
	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.Deposit.Interest.Denom, sdk.NewInt(500)), types.Fetch)
	require.Nil(t, err)

	coins := keeper.GetBankKeeper().GetCoins(ctx, TransferAccAddr)
	require.Equal(t, coins.AmountOf(boxInfo.Deposit.Interest.Denom), sdk.NewInt(500))

	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.Deposit.Interest.Denom, sdk.NewInt(500)), types.Fetch)
	require.Nil(t, err)
	boxInfo = keeper.GetBox(ctx, boxInfo.BoxId)
	require.Len(t, boxInfo.Deposit.InterestInjection, 1)

}
func TestDepositBoxFetchDeposit(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	boxInfo := GetDepositBoxInfo()

	err := keeper.CreateBox(ctx, boxInfo)
	require.Nil(t, err)

	keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.NewCoins(boxInfo.Deposit.Interest))

	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, boxInfo.Deposit.Interest, types.Injection)
	require.Nil(t, err)

	_, err = keeper.ProcessDepositBoxDeposit(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Denom, sdk.NewInt(10000)), types.DepositTo)
	require.Error(t, err)

	boxInfo = keeper.GetBox(ctx, boxInfo.BoxId)
	err = keeper.ProcessDepositBoxByEndBlocker(ctx, boxInfo)
	require.Nil(t, err)

	_, err = keeper.ProcessDepositBoxDeposit(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Denom, sdk.NewInt(10000)), types.DepositTo)
	require.Error(t, err)

	keeper.GetBankKeeper().AddCoins(ctx, TransferAccAddr, sdk.NewCoins(boxInfo.TotalAmount))

	_, err = keeper.ProcessDepositBoxDeposit(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Denom, sdk.NewInt(100000)), types.DepositTo)
	require.Error(t, err)

	_, err = keeper.ProcessDepositBoxDeposit(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Denom, sdk.NewInt(10000)), types.DepositTo)
	require.Nil(t, err)
	keeper.ProcessDepositBoxDeposit(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Denom, sdk.NewInt(5000)), types.Fetch)
	require.Nil(t, err)

	coins := keeper.GetBankKeeper().GetCoins(ctx, TransferAccAddr)
	require.Equal(t, coins.AmountOf(boxInfo.TotalAmount.Denom), sdk.NewInt(5000))

	keeper.ProcessDepositBoxDeposit(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Denom, sdk.NewInt(5000)), types.Fetch)
	require.Nil(t, err)
	require.False(t, keeper.CheckDepositByAddress(ctx, boxInfo.BoxId, TransferAccAddr))

}
