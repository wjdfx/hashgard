package tests

import (
	"testing"

	"github.com/hashgard/hashgard/x/box/msgs"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"

	"github.com/hashgard/hashgard/x/box/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func createDepositBox(t *testing.T, ctx sdk.Context, keeper box.Keeper) *types.BoxInfo {
	boxInfo := GetDepositBoxInfo()

	handler := box.NewHandler(keeper)
	msg := msgs.NewMsgDepositBox(boxInfo)
	res := handler(ctx, msg)
	require.True(t, res.IsOK())

	var boxID string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &boxID)

	box := keeper.GetBox(ctx, boxID)
	require.Equal(t, box.Name, boxInfo.Name)

	return box
}

func TestDepositBoxFetchInterest(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	boxInfo := createDepositBox(t, ctx, keeper)

	keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.NewCoins(boxInfo.Deposit.Interest.Token))
	keeper.GetBankKeeper().AddCoins(ctx, TransferAccAddr, sdk.NewCoins(boxInfo.Deposit.Interest.Token))

	injection := boxInfo.Deposit.Interest.Token.Amount.Quo(sdk.NewInt(2))

	_, err := keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin("error",
		issueutils.MulDecimals(sdk.NewInt(1000), TestTokenDecimals)), types.Injection)
	require.Error(t, err)
	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.Deposit.Interest.Token.Denom,
		boxInfo.Deposit.Interest.Token.Amount.Add(sdk.NewInt(1))), types.Injection)
	require.Error(t, err)
	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.Deposit.Interest.Token.Denom,
		injection), types.Injection)
	require.Nil(t, err)
	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.Deposit.Interest.Token.Denom,
		injection), types.Injection)
	require.Nil(t, err)

	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin("error",
		issueutils.MulDecimals(sdk.NewInt(1000), TestTokenDecimals)), types.Fetch)
	require.Error(t, err)
	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.Deposit.Interest.Token.Denom,
		injection.Mul(sdk.NewInt(10))), types.Fetch)
	require.Error(t, err)

	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.Deposit.Interest.Token.Denom,
		injection), types.Fetch)
	require.Nil(t, err)

	coins := keeper.GetBankKeeper().GetCoins(ctx, TransferAccAddr)
	require.Equal(t, coins.AmountOf(boxInfo.Deposit.Interest.Token.Denom), boxInfo.Deposit.Interest.Token.Amount)

	boxInfo = keeper.GetBox(ctx, boxInfo.BoxId)
	require.Len(t, boxInfo.Deposit.InterestInjections, 1)

	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.Deposit.Interest.Token.Denom,
		injection), types.Fetch)
	require.Nil(t, err)

	coins = keeper.GetBankKeeper().GetCoins(ctx, boxInfo.Owner)
	require.Equal(t, coins.AmountOf(boxInfo.Deposit.Interest.Token.Denom), boxInfo.Deposit.Interest.Token.Amount)

}
func TestDepositBoxFetchDeposit(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	boxInfo := createDepositBox(t, ctx, keeper)

	keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.NewCoins(boxInfo.Deposit.Interest.Token))

	_, err := keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, boxInfo.Deposit.Interest.Token, types.Injection)
	require.Nil(t, err)

	_, err = keeper.ProcessDepositToBox(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Token.Denom, sdk.NewInt(10000)), types.DepositTo)
	require.Error(t, err)

	boxInfo = keeper.GetBox(ctx, boxInfo.BoxId)
	err = keeper.ProcessDepositBoxByEndBlocker(ctx, boxInfo)
	require.Nil(t, err)

	_, err = keeper.ProcessDepositToBox(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Token.Denom, sdk.NewInt(10000)), types.DepositTo)
	require.Error(t, err)

	keeper.GetBankKeeper().AddCoins(ctx, TransferAccAddr, sdk.NewCoins(boxInfo.TotalAmount.Token))

	depositTo := issueutils.MulDecimals(sdk.NewInt(1000), TestTokenDecimals)
	fetch := issueutils.MulDecimals(sdk.NewInt(500), TestTokenDecimals)

	_, err = keeper.ProcessDepositToBox(ctx, boxInfo.BoxId, TransferAccAddr,
		sdk.NewCoin(boxInfo.TotalAmount.Token.Denom, issueutils.MulDecimals(sdk.NewInt(100000), TestTokenDecimals)), types.DepositTo)
	require.Error(t, err)

	_, err = keeper.ProcessDepositToBox(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Token.Denom, depositTo), types.DepositTo)
	require.Nil(t, err)

	_, err = keeper.ProcessDepositToBox(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Token.Denom,
		issueutils.MulDecimals(sdk.NewInt(10000), TestTokenDecimals)), types.Fetch)
	require.Error(t, err)

	_, err = keeper.ProcessDepositToBox(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Token.Denom, fetch), types.Fetch)
	require.Nil(t, err)

	amount := keeper.GetDepositByAddress(ctx, boxInfo.BoxId, TransferAccAddr)
	require.Equal(t, amount, depositTo.Sub(fetch))

	coins := keeper.GetBankKeeper().GetCoins(ctx, TransferAccAddr)
	require.Equal(t, coins.AmountOf(boxInfo.TotalAmount.Token.Denom), boxInfo.TotalAmount.Token.Amount.Sub(depositTo).Add(fetch))

}
