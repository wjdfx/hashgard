package tests

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/mock"

	"github.com/hashgard/hashgard/x/issue"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/types"

	"github.com/hashgard/hashgard/x/issue/keeper"
)

var (
	IssuerCoinsAccAddr   = sdk.AccAddress(crypto.AddressHash([]byte("issuerCoins")))
	ReceiverCoinsAccAddr = sdk.AccAddress(crypto.AddressHash([]byte("receiverCoins")))

	CoinIssueInfo = types.CoinIssueInfo{
		Issuer:          IssuerCoinsAccAddr,
		Name:            "test",
		Symbol:          "tst",
		TotalSupply:     sdk.NewInt(10000),
		Decimals:        types.DefaultDecimals,
		MintingFinished: false}
)

// initialize the mock application for this module
func getMockApp(t *testing.T, numGenAccs int, genState issue.GenesisState, genAccs []auth.Account) (
	mapp *mock.App, keeper keeper.Keeper, addrs []sdk.AccAddress,
	pubKeys []crypto.PubKey, privKeys []crypto.PrivKey) {
	mapp = mock.NewApp()
	msgs.RegisterCodec(mapp.Cdc)
	keyIssue := sdk.NewKVStoreKey(types.StoreKey)
	pk := mapp.ParamsKeeper
	ck := bank.NewBaseKeeper(mapp.AccountKeeper, mapp.ParamsKeeper.Subspace(bank.DefaultParamspace), bank.DefaultCodespace)
	keeper = issue.NewKeeper(mapp.Cdc, keyIssue, pk, pk.Subspace("testissue"), ck, types.DefaultCodespace)
	mapp.Router().AddRoute(types.RouterKey, issue.NewHandler(keeper))
	mapp.QueryRouter().AddRoute(types.QuerierRoute, issue.NewQuerier(keeper))
	//mapp.SetEndBlocker(getEndBlocker(keeper))

	require.NoError(t, mapp.CompleteSetup(keyIssue))

	valTokens := sdk.TokensFromTendermintPower(42)
	if genAccs == nil || len(genAccs) == 0 {
		genAccs, addrs, pubKeys, privKeys = mock.CreateGenAccounts(numGenAccs,
			sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, valTokens)})
	}

	mock.SetGenesis(mapp, genAccs)

	return mapp, keeper, addrs, pubKeys, privKeys
}
