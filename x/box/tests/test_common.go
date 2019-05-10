package tests

import (
	"testing"
	"time"

	"github.com/hashgard/hashgard/x/box/types"

	"github.com/cosmos/cosmos-sdk/x/staking"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/mock"

	"github.com/hashgard/hashgard/x/box"
	"github.com/hashgard/hashgard/x/box/msgs"

	"github.com/hashgard/hashgard/x/box/keeper"
)

var (
	TransferAccAddr = sdk.AccAddress(crypto.AddressHash([]byte("transferAddress")))
	SenderAccAddr   = sdk.AccAddress(crypto.AddressHash([]byte("senderAddress")))

	newBoxInfo = types.BoxInfo{
		Owner:         SenderAccAddr,
		CreatedTime:   time.Now(),
		Name:          "testBox",
		TotalAmount:   sdk.NewCoin("text", sdk.NewInt(10000)),
		BoxType:       types.Lock,
		Description:   "{}",
		TradeDisabled: true,
	}
)

func GetLockBoxInfo() *types.BoxInfo {
	boxInfo := newBoxInfo
	boxInfo.TotalAmount.Amount = sdk.NewInt(10000)
	boxInfo.BoxType = types.Lock
	boxInfo.Lock.EndTime = time.Now().Add(time.Duration(1) * time.Minute)
	return &boxInfo
}
func GetDepositBoxInfo() *types.BoxInfo {
	boxInfo := newBoxInfo
	boxInfo.TotalAmount.Amount = sdk.NewInt(10000)
	boxInfo.BoxType = types.Deposit
	boxInfo.Deposit = types.DepositBox{
		StartTime:     time.Now().Add(time.Duration(24) * time.Hour),
		EstablishTime: time.Now().Add(time.Duration(48) * time.Hour),
		MaturityTime:  time.Now().Add(time.Duration(96) * time.Hour),
		BottomLine:    sdk.NewInt(200),
		Price:         sdk.NewInt(100),
		Interest:      sdk.NewCoin("interest", sdk.NewInt(1000)),
		Coupon:        sdk.ZeroInt()}
	return &boxInfo
}
func GetFutureBoxInfo() *types.BoxInfo {
	boxInfo := newBoxInfo
	boxInfo.TotalAmount.Amount = sdk.NewInt(10000)
	boxInfo.BoxType = types.Future
	boxInfo.Future.TimeLine = []time.Time{time.Now().Add(time.Duration(24*30*1) * time.Hour), time.Now().Add(time.Duration(24*30*2) * time.Hour)}
	boxInfo.Future.Receivers = [][]string{{"gardvaloper1k67xljpc0lr678wyl6vld9hy3t2lc6ph2fecaf", "188", "200"},
		{"gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n", "300", "500"}}
	return &boxInfo
}

// gov and staking endblocker
func getEndBlocker(keeper keeper.Keeper) sdk.EndBlocker {
	return func(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
		tags := box.EndBlocker(ctx, keeper)
		return abci.ResponseEndBlock{
			Tags: tags,
		}
	}
}

// initialize the mock application for this module
func getMockApp(t *testing.T, numGenAccs int, genState box.GenesisState, genAccs []auth.Account) (
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
	keeper = box.NewKeeper(mapp.Cdc, keyIssue, pk, pk.Subspace("testBox"), ck, types.DefaultCodespace)

	mapp.Router().AddRoute(types.RouterKey, box.NewHandler(keeper))
	mapp.QueryRouter().AddRoute(types.QuerierRoute, box.NewQuerier(keeper))
	mapp.SetEndBlocker(getEndBlocker(keeper))
	mapp.SetInitChainer(getInitChainer(mapp, keeper, sk, genState))

	require.NoError(t, mapp.CompleteSetup(keyIssue))

	valTokens := sdk.TokensFromTendermintPower(42)
	if len(genAccs) == 0 {
		genAccs, addrs, pubKeys, privKeys = mock.CreateGenAccounts(numGenAccs,
			sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, valTokens)))
	}

	mock.SetGenesis(mapp, genAccs)

	return mapp, keeper, sk, addrs, pubKeys, privKeys
}
func getInitChainer(mapp *mock.App, keeper keeper.Keeper, stakingKeeper staking.Keeper, genState box.GenesisState) sdk.InitChainer {

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
			box.InitGenesis(ctx, keeper, box.DefaultGenesisState())
		} else {
			box.InitGenesis(ctx, keeper, genState)
		}
		return abci.ResponseInitChain{
			//Validators: validators,
		}
	}
}
