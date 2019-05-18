package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestGetLockBoxByAddress(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	handler := box.NewHandler(keeper)

	boxInfo := GetLockBoxInfo()
	boxInfo.Sender, _ = sdk.AccAddressFromBech32("TestGetBoxByAddress")
	cap := 10
	for i := 0; i < cap; i++ {
		keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Sender, sdk.NewCoins(boxInfo.TotalAmount.Token))

		msg := msgs.NewMsgLockBox(boxInfo)
		res := handler(ctx, msg)
		require.True(t, res.IsOK())
	}
	issues := keeper.GetBoxByAddress(ctx, boxInfo.BoxType, boxInfo.Sender)

	require.Len(t, issues, cap)
}
