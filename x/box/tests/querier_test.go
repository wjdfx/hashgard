package tests

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/hashgard/hashgard/x/box/params"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/types"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hashgard/hashgard/x/box"
)

func TestDepositBoxList(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.NewContext(false, abci.Header{})
	cap := 1000
	for i := 0; i < cap; i++ {

		duration, _ := time.ParseDuration(strconv.Itoa(i) + "m")
		boxInfo := GetDepositBoxInfo()
		_, err := keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.Coins{boxInfo.TotalAmount})
		require.Nil(t, err)
		boxInfo.CreatedTime = time.Now().Add(duration)

		err = keeper.CreateBox(ctx, boxInfo)
		if err != nil {
			fmt.Println(err.Error())
		}
		require.Nil(t, err)
	}

	boxID := ""
	for i := 0; i < 100; i++ {
		//fmt.Println("==================page:" + strconv.Itoa(i))
		boxs := keeper.List(ctx, params.BoxQueryParams{StartBoxId: boxID, BoxType: types.Deposit, Owner: nil, Limit: 10})
		require.Len(t, boxs, 10)
		for j, box := range boxs {

			if j > 0 {
				require.True(t, boxs[j].CreatedTime.Before(boxs[j-1].CreatedTime))
			}
			//fmt.Println(box.BoxId + "----" + box.CreatedTime.String())
			boxID = box.BoxId
		}

	}

}
func TestBoxList(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.NewContext(false, abci.Header{})

	boxTypes := []string{types.Lock, types.Deposit, types.Future}
	cap := 1000
	boxTypeCount := map[string]int{types.Lock: 0, types.Deposit: 0, types.Future: 0}
	for i := 0; i < cap; i++ {

		duration, _ := time.ParseDuration(strconv.Itoa(i) + "m")
		boxInfo := GetDepositBoxInfo()
		boxInfo.BoxType = boxTypes[rand.Intn(len(boxTypes))]
		_, err := keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.Coins{boxInfo.TotalAmount})
		require.Nil(t, err)
		boxInfo.CreatedTime = time.Now().Add(duration)

		err = keeper.CreateBox(ctx, boxInfo)
		if err != nil {
			fmt.Println(err.Error())
		}
		require.Nil(t, err)
		boxTypeCount[boxInfo.BoxType] = boxTypeCount[boxInfo.BoxType] + 1
	}
	pageSize := 10
	for key, value := range boxTypeCount {
		boxId := ""
		page := (value + pageSize - 1) / pageSize
		for i := 0; i < page; i++ {
			boxs := keeper.List(ctx, params.BoxQueryParams{StartBoxId: boxId, BoxType: key, Owner: nil, Limit: pageSize})
			if i == page-1 {
				require.Len(t, boxs, value%pageSize)
			} else {
				require.Len(t, boxs, pageSize)
			}

			for j, box := range boxs {

				if j > 0 {
					require.True(t, boxs[j].CreatedTime.Before(boxs[j-1].CreatedTime))
				}
				boxId = box.BoxId
				//fmt.Println(boxId)
			}

		}
	}

}
func TestQueryDepositListFromDepositBox(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, box.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.NewContext(false, abci.Header{})

	boxInfo := GetDepositBoxInfo()
	err := keeper.CreateBox(ctx, boxInfo)
	require.Nil(t, err)

	keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.NewCoins(boxInfo.Deposit.Interest))

	_, err = keeper.ProcessDepositBoxInterest(ctx, boxInfo.BoxId, boxInfo.Owner, boxInfo.Deposit.Interest, types.Injection)
	require.Nil(t, err)

	boxInfo = keeper.GetBox(ctx, boxInfo.BoxId)
	err = keeper.ProcessDepositBoxByEndBlocker(ctx, boxInfo)
	require.Nil(t, err)

	keeper.GetBankKeeper().AddCoins(ctx, TransferAccAddr, sdk.NewCoins(boxInfo.TotalAmount))
	keeper.GetBankKeeper().AddCoins(ctx, boxInfo.Owner, sdk.NewCoins(boxInfo.TotalAmount))

	_, err = keeper.ProcessDepositBoxDeposit(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Denom, sdk.NewInt(5000)), types.DepositTo)
	require.Nil(t, err)

	_, err = keeper.ProcessDepositBoxDeposit(ctx, boxInfo.BoxId, TransferAccAddr, sdk.NewCoin(boxInfo.TotalAmount.Denom, sdk.NewInt(1000)), types.DepositTo)
	require.Nil(t, err)

	_, err = keeper.ProcessDepositBoxDeposit(ctx, boxInfo.BoxId, boxInfo.Owner, sdk.NewCoin(boxInfo.TotalAmount.Denom, sdk.NewInt(1000)), types.DepositTo)
	require.Nil(t, err)

	list := keeper.QueryDepositListFromDepositBox(ctx, boxInfo.BoxId, nil)
	require.NotEmpty(t, list)
	require.Equal(t, list[0].Amount, sdk.NewInt(6000))
	require.Equal(t, list[1].Amount, sdk.NewInt(1000))

	list = keeper.QueryDepositListFromDepositBox(ctx, boxInfo.BoxId, boxInfo.Owner)
	require.NotEmpty(t, list)
	require.Equal(t, list[0].Amount, sdk.NewInt(1000))

}
