package issue

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"testing"
)

// Parameter store key
var (

	// TODO: Find another way to implement this without using accounts, or find a cleaner way to implement it using accounts.
	IssuerCoinsAccAddr   = sdk.AccAddress(crypto.AddressHash([]byte("issuerCoins")))
	ReceiverCoinsAccAddr = sdk.AccAddress(crypto.AddressHash([]byte("receiverCoins")))
	coinIssueInfo        = CoinIssueInfo{
		IssuerCoinsAccAddr,
		"test",
		"tst",
		sdk.NewInt(10000),
		DefaultDecimals,
		false}
)

func TestAddIssue(t *testing.T) {

	mapp, keeper, _, _, _ := getMockApp(t, 0, GenesisState{}, nil)
	mapp.BeginBlock(abci.RequestBeginBlock{})
	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	mapp.InitChainer(ctx, abci.RequestInitChain{})

	issueID, _, _, err := keeper.AddIssue(ctx, NewMsgIssue(IssuerCoinsAccAddr, &coinIssueInfo))

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

	mapp, keeper, _, _, _ := getMockApp(t, 0, GenesisState{}, nil)
	mapp.BeginBlock(abci.RequestBeginBlock{})
	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	mapp.InitChainer(ctx, abci.RequestInitChain{})

	issueID, _, _, err := keeper.AddIssue(ctx, NewMsgIssue(IssuerCoinsAccAddr, &coinIssueInfo))

	require.Nil(t, err)
	keeper.Mint(ctx, issueID, sdk.NewInt(10000), ReceiverCoinsAccAddr)

	//require.True(t,res.IsOK())
	//
	//str:=sdk.TagsToStringTags(res.Tags).String()
	//
	//require.True(t, len(str)>0)

	coinIssue := keeper.GetIssue(ctx, issueID)

	require.True(t, coinIssue.TotalSupply.Equal(sdk.NewInt(20000)))

	acc := mapp.AccountKeeper.GetAccount(ctx, ReceiverCoinsAccAddr)

	amount := acc.GetCoins().AmountOf(issueID)
	flag1 := amount.Equal(sdk.NewInt(10000))

	require.True(t, flag1)
	//
	//keeper.burn(ctx,issueID,sdk.NewInt(5000),IssuerCoinsAccAddr)
	//
	//coinIssueInfo1:=keeper.getIssue(ctx,issueID)
	//
	//require.True(t,coinIssueInfo1.TotalSupply.Equal(sdk.NewInt(5000)))
}

func TestBurn(t *testing.T) {

	mapp, keeper, _, _, _ := getMockApp(t, 0, GenesisState{}, nil)
	mapp.BeginBlock(abci.RequestBeginBlock{})
	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	mapp.InitChainer(ctx, abci.RequestInitChain{})

	issueID, _, _, err := keeper.AddIssue(ctx, NewMsgIssue(IssuerCoinsAccAddr, &coinIssueInfo))

	require.Nil(t, err)

	keeper.Burn(ctx, issueID, sdk.NewInt(5000), IssuerCoinsAccAddr)

	coinIssue := keeper.GetIssue(ctx, issueID)

	require.True(t, coinIssue.TotalSupply.Equal(sdk.NewInt(5000)))

	acc := mapp.AccountKeeper.GetAccount(ctx, IssuerCoinsAccAddr)

	amount := acc.GetCoins().AmountOf(issueID)
	flag1 := amount.Equal(sdk.NewInt(5000))

	require.True(t, flag1)
}
