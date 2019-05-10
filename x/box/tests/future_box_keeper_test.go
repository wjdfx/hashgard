package tests

import (
	"testing"

	"github.com/hashgard/hashgard/x/box"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestFutureBoxAdd(t *testing.T) {
	//TODO
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	boxInfo := GetFutureBoxInfo()

	err := keeper.CreateBox(ctx, boxInfo)
	require.Nil(t, err)
	box := keeper.GetBox(ctx, boxInfo.BoxId)
	require.Equal(t, boxInfo.Name, box.Name)
}
