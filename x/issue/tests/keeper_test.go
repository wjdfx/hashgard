package tests

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"

	"github.com/hashgard/hashgard/x/issue"
)

func TestAddIssue(t *testing.T) {

	mapp, keeper, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)
	mapp.BeginBlock(abci.RequestBeginBlock{})
	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	mapp.InitChainer(ctx, abci.RequestInitChain{})

	var _, _, err = keeper.AddIssue(ctx, &CoinIssueInfo)
	require.Nil(t, err)
	coinIssue := keeper.GetIssue(ctx, CoinIssueInfo.IssueId)
	require.Equal(t, coinIssue.TotalSupply, CoinIssueInfo.TotalSupply)
	coin := sdk.Coin{Denom: CoinIssueInfo.IssueId, Amount: sdk.NewInt(5000)}
	keeper.SendCoins(ctx, IssuerCoinsAccAddr, ReceiverCoinsAccAddr,
		sdk.Coins{coin})
	coinIssue = keeper.GetIssue(ctx, CoinIssueInfo.IssueId)
	require.True(t, coinIssue.TotalSupply.Equal(CoinIssueInfo.TotalSupply))
	acc := mapp.AccountKeeper.GetAccount(ctx, ReceiverCoinsAccAddr)
	amount := acc.GetCoins().AmountOf(CoinIssueInfo.IssueId)
	flag1 := amount.Equal(coin.Amount)
	require.True(t, flag1)
}

func TestMint(t *testing.T) {

	mapp, keeper, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)
	mapp.BeginBlock(abci.RequestBeginBlock{})
	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	mapp.InitChainer(ctx, abci.RequestInitChain{})
	_, _, err := keeper.AddIssue(ctx, &CoinIssueInfo)
	require.Nil(t, err)
	keeper.Mint(ctx, &CoinIssueInfo, sdk.NewInt(10000), IssuerCoinsAccAddr)
	coinIssue := keeper.GetIssue(ctx, CoinIssueInfo.IssueId)
	require.True(t, coinIssue.TotalSupply.Equal(sdk.NewInt(20000)))
}

func TestBurn(t *testing.T) {

	mapp, keeper, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)
	mapp.BeginBlock(abci.RequestBeginBlock{})
	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	mapp.InitChainer(ctx, abci.RequestInitChain{})

	_, _, err := keeper.AddIssue(ctx, &CoinIssueInfo)
	require.Nil(t, err)
	keeper.Burn(ctx, &CoinIssueInfo, sdk.NewInt(5000), IssuerCoinsAccAddr)
	coinIssue := keeper.GetIssue(ctx, CoinIssueInfo.IssueId)
	require.True(t, coinIssue.TotalSupply.Equal(sdk.NewInt(5000)))
}
