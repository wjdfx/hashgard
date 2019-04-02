package tests

import (
	"bytes"
	"log"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/mock"

	"github.com/hashgard/hashgard/x/issue"
	"github.com/hashgard/hashgard/x/issue/domain"

	"github.com/hashgard/hashgard/x/issue/keepers"
)

var (
	IssuerCoinsAccAddr   = sdk.AccAddress(crypto.AddressHash([]byte("issuerCoins")))
	ReceiverCoinsAccAddr = sdk.AccAddress(crypto.AddressHash([]byte("receiverCoins")))

	CoinIssueInfo = domain.CoinIssueInfo{
		"",
		IssuerCoinsAccAddr,
		"test",
		"tst",
		sdk.NewInt(10000),
		domain.DefaultDecimals,
		false}
)

// initialize the mock application for this module
func getMockApp(t *testing.T, numGenAccs int, genState issue.GenesisState, genAccs []auth.Account) (
	mapp *mock.App, keeper keepers.Keeper, addrs []sdk.AccAddress,
	pubKeys []crypto.PubKey, privKeys []crypto.PrivKey) {
	mapp = mock.NewApp()
	issue.RegisterCodec(mapp.Cdc)
	keyIssue := sdk.NewKVStoreKey(domain.StoreKey)
	pk := mapp.ParamsKeeper
	ck := bank.NewBaseKeeper(mapp.AccountKeeper, mapp.ParamsKeeper.Subspace(bank.DefaultParamspace), bank.DefaultCodespace)
	keeper = keepers.NewKeeper(mapp.Cdc, keyIssue, pk, pk.Subspace("testissue"), ck, domain.DefaultCodespace)
	mapp.Router().AddRoute(domain.RouterKey, issue.NewHandler(keeper))
	mapp.QueryRouter().AddRoute(domain.QuerierRoute, issue.NewQuerier(keeper))
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

// TODO: Remove once address interface has been implemented (ref: #2186)
func SortValAddresses(addrs []sdk.ValAddress) {
	var byteAddrs [][]byte
	for _, addr := range addrs {
		byteAddrs = append(byteAddrs, addr.Bytes())
	}
	SortByteArrays(byteAddrs)
	for i, byteAddr := range byteAddrs {
		addrs[i] = byteAddr
	}
}

// implement `Interface` in sort package.
type sortByteArrays [][]byte

func (b sortByteArrays) Len() int {
	return len(b)
}
func (b sortByteArrays) Less(i, j int) bool {
	// bytes package already implements Comparable for []byte.
	switch bytes.Compare(b[i], b[j]) {
	case -1:
		return true
	case 0, 1:
		return false
	default:
		log.Panic("not fail-able with `bytes.Comparable` bounded [-1, 1].")
		return false
	}
}
func (b sortByteArrays) Swap(i, j int) {
	b[j], b[i] = b[i], b[j]
}

// Public
func SortByteArrays(src [][]byte) [][]byte {
	sorted := sortByteArrays(src)
	sort.Sort(sorted)
	return sorted
}