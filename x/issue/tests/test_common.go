package tests

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/x/staking"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/stretchr/testify/require"

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
	SenderAccAddr        = sdk.AccAddress(crypto.AddressHash([]byte("senderAddress")))

	CoinIssueInfo = types.CoinIssueInfo{
		Issuer:             IssuerCoinsAccAddr,
		Owner:              IssuerCoinsAccAddr,
		IssueTime:          time.Now(),
		Name:               "testCoin",
		Symbol:             "TEST",
		TotalSupply:        sdk.NewInt(10000),
		Decimals:           types.CoinDecimalsMaxValue,
		BurnOwnerDisabled:  false,
		BurnHolderDisabled: false,
		BurnFromDisabled:   false,
		MintingFinished:    false}
)

// initialize the mock application for this module
func getMockApp(t *testing.T, numGenAccs int, genState issue.GenesisState, genAccs []auth.Account) (
	mapp *mock.App, keeper keeper.Keeper, sk staking.Keeper, addrs []sdk.AccAddress,
	pubKeys []crypto.PubKey, privKeys []crypto.PrivKey) {
	mapp = mock.NewApp()
	msgs.RegisterCodec(mapp.Cdc)
	keyIssue := sdk.NewKVStoreKey(types.StoreKey)

	keyStaking := sdk.NewKVStoreKey(staking.StoreKey)
	tkeyStaking := sdk.NewTransientStoreKey(staking.TStoreKey)

	pk := mapp.ParamsKeeper
	ck := bank.NewBaseKeeper(mapp.AccountKeeper, mapp.ParamsKeeper.Subspace(bank.DefaultParamspace), bank.DefaultCodespace)

	sk = staking.NewKeeper(mapp.Cdc, keyStaking, tkeyStaking, ck, pk.Subspace(staking.DefaultParamspace), staking.DefaultCodespace)
	keeper = issue.NewKeeper(mapp.Cdc, keyIssue, pk, pk.Subspace("testissue"), ck, types.DefaultCodespace)

	mapp.Router().AddRoute(types.RouterKey, issue.NewHandler(keeper))
	mapp.QueryRouter().AddRoute(types.QuerierRoute, issue.NewQuerier(keeper))
	//mapp.SetEndBlocker(getEndBlocker(keeper))
	mapp.SetInitChainer(getInitChainer(mapp, keeper, sk, genState))

	require.NoError(t, mapp.CompleteSetup(keyIssue))

	valTokens := sdk.TokensFromTendermintPower(42)
	if len(genAccs) == 0 {
		genAccs, addrs, pubKeys, privKeys = mock.CreateGenAccounts(numGenAccs,
			sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, valTokens)})
	}

	mock.SetGenesis(mapp, genAccs)

	return mapp, keeper, sk, addrs, pubKeys, privKeys
}
func getInitChainer(mapp *mock.App, keeper keeper.Keeper, stakingKeeper staking.Keeper, genState issue.GenesisState) sdk.InitChainer {

	return func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {

		mapp.InitChainer(ctx, req)

		stakingGenesis := staking.DefaultGenesisState()
		tokens := sdk.TokensFromTendermintPower(100000)
		stakingGenesis.Pool.NotBondedTokens = tokens

		//validators, err := staking.InitGenesis(ctx, stakingKeeper, stakingGenesis)
		//if err != nil {
		//	panic(err)
		//}
		if genState.IsEmpty() {
			issue.InitGenesis(ctx, keeper, issue.DefaultGenesisState())
		} else {
			issue.InitGenesis(ctx, keeper, genState)
		}
		return abci.ResponseInitChain{
			//Validators: validators,
		}
	}
}
