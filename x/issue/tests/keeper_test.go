package tests

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"testing"

	"github.com/hashgard/hashgard/x/issue"
	"github.com/hashgard/hashgard/x/issue/domain"
	"github.com/hashgard/hashgard/x/issue/msgs"
)

// Parameter store key
var (
	// TODO: Find another way to implement this without using accounts, or find a cleaner way to implement it using accounts.
	IssuerCoinsAccAddr   = sdk.AccAddress(crypto.AddressHash([]byte("issuerCoins")))
	ReceiverCoinsAccAddr = sdk.AccAddress(crypto.AddressHash([]byte("receiverCoins")))
	coinIssueInfo        = domain.CoinIssueInfo{
		"",
		IssuerCoinsAccAddr,
		"test",
		"tst",
		sdk.NewInt(10000),
		domain.DefaultDecimals,
		false}
)

func TestAddIssue(t *testing.T) {

	mapp, keeper, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)
	mapp.BeginBlock(abci.RequestBeginBlock{})
	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	mapp.InitChainer(ctx, abci.RequestInitChain{})
	issueID, _, _, err := keeper.AddIssue(ctx, msgs.NewMsgIssue(&coinIssueInfo))
	require.Nil(t, err)
	coin := sdk.Coin{Denom: issueID, Amount: sdk.NewInt(5000)}
	keeper.SendCoins(ctx, IssuerCoinsAccAddr, ReceiverCoinsAccAddr,
		sdk.Coins{coin})
	coinIssue := keeper.GetIssue(ctx, issueID)
	require.True(t, coinIssue.TotalSupply.Equal(coinIssueInfo.TotalSupply))
	acc := mapp.AccountKeeper.GetAccount(ctx, ReceiverCoinsAccAddr)
	amount := acc.GetCoins().AmountOf(issueID)
	flag1 := amount.Equal(coin.Amount)
	require.True(t, flag1)
}

func TestMint(t *testing.T) {

	mapp, keeper, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)
	mapp.BeginBlock(abci.RequestBeginBlock{})
	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	mapp.InitChainer(ctx, abci.RequestInitChain{})
	issueID, _, _, err := keeper.AddIssue(ctx, msgs.NewMsgIssue(&coinIssueInfo))
	require.Nil(t, err)
	keeper.Mint(ctx, issueID, sdk.NewInt(10000))
	coinIssue := keeper.GetIssue(ctx, issueID)
	require.True(t, coinIssue.TotalSupply.Equal(sdk.NewInt(20000)))
}

func TestBurn(t *testing.T) {

	mapp, keeper, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)
	mapp.BeginBlock(abci.RequestBeginBlock{})
	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	mapp.InitChainer(ctx, abci.RequestInitChain{})

	issueID, _, _, err := keeper.AddIssue(ctx, msgs.NewMsgIssue(&coinIssueInfo))
	require.Nil(t, err)
	keeper.Burn(ctx, issueID, sdk.NewInt(5000))
	coinIssue := keeper.GetIssue(ctx, issueID)
	require.True(t, coinIssue.TotalSupply.Equal(sdk.NewInt(5000)))
}
