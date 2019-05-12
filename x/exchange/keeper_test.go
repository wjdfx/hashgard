package exchange

import (
	"testing"

	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

func TestGetSetOrder(t *testing.T) {
	mapp, keeper, _, addrs, _, _ := getMockApp(t, 1, GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	order, err := keeper.CreateOrder(ctx, addrs[0], sdk.NewInt64Coin(sdk.DefaultBondDenom, 100), sdk.NewInt64Coin("foocoin", 200))
	require.NoError(t, err)
	orderId := order.OrderId

	_, ok := keeper.GetOrder(ctx, orderId)
	require.True(t, ok)
}

func TestIncrementOrderNumber(t *testing.T) {
	mapp, keeper, _, addrs, _, _ := getMockApp(t, 1, GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	keeper.CreateOrder(ctx, addrs[0], sdk.NewInt64Coin(sdk.DefaultBondDenom, 10), sdk.NewInt64Coin("foocoin", 200))
	keeper.CreateOrder(ctx, addrs[0], sdk.NewInt64Coin(sdk.DefaultBondDenom, 10), sdk.NewInt64Coin("foocoin", 200))
	order3, err := keeper.CreateOrder(ctx, addrs[0], sdk.NewInt64Coin(sdk.DefaultBondDenom, 10), sdk.NewInt64Coin("foocoin", 200))
	require.NoError(t, err)

	require.Equal(t, uint64(3), order3.OrderId)
}

func TestWithrawalOrder(t *testing.T) {
	mapp, keeper, _, addrs, _, _ := getMockApp(t, 1, GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	order, err := keeper.CreateOrder(ctx, addrs[0], sdk.NewInt64Coin(sdk.DefaultBondDenom, 100), sdk.NewInt64Coin("foocoin", 200))
	require.NoError(t, err)

	_, err = keeper.WithdrawalOrder(ctx, order.OrderId, order.Seller)
	require.NoError(t, err)
}

func TestTakeOrder(t *testing.T) {
	acc1 := auth.NewBaseAccountWithAddress(Addrs[0])
	acc1.SetCoins(sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000)))
	acc2 := auth.NewBaseAccountWithAddress(Addrs[1])
	acc2.SetCoins(sdk.NewCoins(sdk.NewInt64Coin("foocoin", 500)))

	mapp, keeper, _, _, _, _ := getMockApp(t, 0, GenesisState{}, []auth.Account{&acc1, &acc2})

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	order, err := keeper.CreateOrder(ctx, acc1.Address, sdk.NewInt64Coin(sdk.DefaultBondDenom, 100), sdk.NewInt64Coin("foocoin", 200))
	require.NoError(t, err)

	_, _, soldOut, err := keeper.TakeOrder(ctx, order.OrderId, acc2.Address, sdk.NewInt64Coin("foocoin", 100))
	require.NoError(t, err)
	require.False(t, soldOut)
}