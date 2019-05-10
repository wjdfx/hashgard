package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestGetBoxByAddress(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	boxInfo := GetLockBoxInfo()
	boxInfo.Owner, _ = sdk.AccAddressFromBech32("TestGetBoxByAddress")
	cap := 10
	for i := 0; i < cap; i++ {
		keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.NewCoins(boxInfo.TotalAmount))
		err := keeper.CreateBox(ctx, boxInfo)
		require.Nil(t, err)
	}
	issues := keeper.GetBoxByAddress(ctx, boxInfo.BoxType, boxInfo.Owner)

	require.Len(t, issues, cap)
}
