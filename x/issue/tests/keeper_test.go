package tests

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hashgard/hashgard/x/issue"
)

func TestAddIssue(t *testing.T) {

	mapp, keeper, _, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	_, _, err := keeper.AddIssue(ctx, &CoinIssueInfo)
	require.Nil(t, err)
	coinIssue := keeper.GetIssue(ctx, CoinIssueInfo.IssueId)
	require.Equal(t, coinIssue.TotalSupply, CoinIssueInfo.TotalSupply)
	coin := sdk.Coin{Denom: CoinIssueInfo.IssueId, Amount: sdk.NewInt(5000)}
	_, err = keeper.SendCoins(ctx, IssuerCoinsAccAddr, ReceiverCoinsAccAddr,
		sdk.NewCoins(coin))
	require.Nil(t, err)
	coinIssue = keeper.GetIssue(ctx, CoinIssueInfo.IssueId)
	require.True(t, coinIssue.TotalSupply.Equal(CoinIssueInfo.TotalSupply))
	acc := mapp.AccountKeeper.GetAccount(ctx, ReceiverCoinsAccAddr)
	amount := acc.GetCoins().AmountOf(CoinIssueInfo.IssueId)
	flag1 := amount.Equal(coin.Amount)
	require.True(t, flag1)
}

func TestGetIssues(t *testing.T) {
	fmt.Print(time.Now().Unix())
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	cap := 10
	for i := 0; i < cap; i++ {
		_, _, err := keeper.AddIssue(ctx, &CoinIssueInfo)
		require.Nil(t, err)
	}
	issues := keeper.GetIssues(ctx, CoinIssueInfo.Issuer.String())

	require.Len(t, issues, cap)
}

func TestMint(t *testing.T) {

	mapp, keeper, _, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	CoinIssueInfo.TotalSupply = sdk.NewInt(10000)
	_, _, err := keeper.AddIssue(ctx, &CoinIssueInfo)
	require.Nil(t, err)
	_, _, err = keeper.Mint(ctx, CoinIssueInfo.IssueId, sdk.NewInt(10000), IssuerCoinsAccAddr, IssuerCoinsAccAddr)
	require.Nil(t, err)
	coinIssue := keeper.GetIssue(ctx, CoinIssueInfo.IssueId)
	require.True(t, coinIssue.TotalSupply.Equal(sdk.NewInt(20000)))
}

func TestBurnOwner(t *testing.T) {

	mapp, keeper, _, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	CoinIssueInfo.TotalSupply = sdk.NewInt(10000)

	_, _, err := keeper.AddIssue(ctx, &CoinIssueInfo)
	require.Nil(t, err)

	_, _, err = keeper.BurnOwner(ctx, CoinIssueInfo.IssueId, sdk.NewInt(5000), IssuerCoinsAccAddr)
	require.Nil(t, err)

	err = keeper.DisableBurnOwner(ctx, CoinIssueInfo.Owner, CoinIssueInfo.IssueId)
	require.Nil(t, err)

	_, _, err = keeper.BurnOwner(ctx, CoinIssueInfo.IssueId, sdk.NewInt(5000), IssuerCoinsAccAddr)
	require.Error(t, err)

}

func TestBurnHolder(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	CoinIssueInfo.TotalSupply = sdk.NewInt(10000)

	_, _, err := keeper.AddIssue(ctx, &CoinIssueInfo)
	require.Nil(t, err)

	_, err = keeper.SendCoins(ctx, IssuerCoinsAccAddr, ReceiverCoinsAccAddr, sdk.NewCoins(sdk.NewCoin(CoinIssueInfo.IssueId, sdk.NewInt(10000))))
	require.Nil(t, err)

	_, _, err = keeper.BurnHolder(ctx, CoinIssueInfo.IssueId, sdk.NewInt(5000), ReceiverCoinsAccAddr)
	require.Nil(t, err)

	err = keeper.DisableBurnHolder(ctx, CoinIssueInfo.Owner, CoinIssueInfo.IssueId)
	require.Nil(t, err)

	_, _, err = keeper.BurnHolder(ctx, CoinIssueInfo.IssueId, sdk.NewInt(5000), ReceiverCoinsAccAddr)
	require.Error(t, err)

}

func TestBurnFrom(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	CoinIssueInfo.TotalSupply = sdk.NewInt(10000)

	_, _, err := keeper.AddIssue(ctx, &CoinIssueInfo)
	require.Nil(t, err)

	_, err = keeper.SendCoins(ctx, IssuerCoinsAccAddr, ReceiverCoinsAccAddr, sdk.NewCoins(sdk.NewCoin(CoinIssueInfo.IssueId, sdk.NewInt(10000))))
	require.Nil(t, err)

	_, _, err = keeper.BurnFrom(ctx, CoinIssueInfo.IssueId, sdk.NewInt(5000), IssuerCoinsAccAddr, ReceiverCoinsAccAddr)
	require.Nil(t, err)

	err = keeper.DisableBurnFrom(ctx, CoinIssueInfo.Owner, CoinIssueInfo.IssueId)
	require.Nil(t, err)

	_, _, err = keeper.BurnFrom(ctx, CoinIssueInfo.IssueId, sdk.NewInt(5000), ReceiverCoinsAccAddr, ReceiverCoinsAccAddr)
	require.Error(t, err)
}
