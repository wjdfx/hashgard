package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/types"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"

	"github.com/hashgard/hashgard/x/box"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func createFutureBox(t *testing.T, ctx sdk.Context, keeper box.Keeper) *types.BoxInfo {
	boxInfo := GetFutureBoxInfo()

	handler := box.NewHandler(keeper)
	msg := msgs.NewMsgFutureBox(newBoxInfo.Owner, boxInfo)
	res := handler(ctx, msg)
	require.True(t, res.IsOK())

	var boxID string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &boxID)

	box := keeper.GetBox(ctx, boxID)
	require.Equal(t, box.Name, boxInfo.Name)

	return box
}

func TestFutureBoxAdd(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	boxInfo := createFutureBox(t, ctx, keeper)

	err := keeper.CreateBox(ctx, boxInfo)
	require.Nil(t, err)
	box := keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, boxInfo.Name, box.Name)
}

func TestFutureBoxFetchDeposit(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	boxInfo := createFutureBox(t, ctx, keeper)

	err := keeper.CreateBox(ctx, boxInfo)
	require.Nil(t, err)
	box := keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, boxInfo.Name, box.Name)

	keeper.GetBankKeeper().AddCoins(ctx, TransferAccAddr, sdk.NewCoins(boxInfo.TotalAmount.Token))

	depositTo := issueutils.MulDecimals(sdk.NewInt(1000), TestTokenDecimals)
	fetch := issueutils.MulDecimals(sdk.NewInt(500), TestTokenDecimals)

	_, err = keeper.ProcessDepositToBox(ctx, boxInfo.BoxId, TransferAccAddr,
		sdk.NewCoin(boxInfo.TotalAmount.Token.Denom,
			issueutils.MulDecimals(sdk.NewInt(10000), TestTokenDecimals)), types.DepositTo)
	require.Error(t, err)

	_, err = keeper.ProcessDepositToBox(ctx, boxInfo.BoxId, TransferAccAddr,
		sdk.NewCoin(boxInfo.TotalAmount.Token.Denom, depositTo), types.DepositTo)
	require.Nil(t, err)

	_, err = keeper.ProcessDepositToBox(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Token.Denom,
		issueutils.MulDecimals(sdk.NewInt(5000), TestTokenDecimals)), types.Fetch)
	require.Error(t, err)

	_, err = keeper.ProcessDepositToBox(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Token.Denom, fetch), types.Fetch)
	require.Nil(t, err)

	newBoxInfo := keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, newBoxInfo.Future.Deposits[0].Amount, depositTo.Sub(fetch))

	coins := keeper.GetBankKeeper().GetCoins(ctx, TransferAccAddr)
	require.Equal(t, coins.AmountOf(boxInfo.TotalAmount.Token.Denom), boxInfo.TotalAmount.Token.Amount.Sub(depositTo).Add(fetch))
}
