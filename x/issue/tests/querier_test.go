package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue"
	queriers2 "github.com/hashgard/hashgard/x/issue/client/queriers"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/types"
)

func TestQueryIssue(t *testing.T) {
	mapp, keeper, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)
	mapp.BeginBlock(abci.RequestBeginBlock{})
	ctx := mapp.NewContext(false, abci.Header{})
	mapp.InitChainer(ctx, abci.RequestInitChain{})

	querier := issue.NewQuerier(keeper)
	handler := issue.NewHandler(keeper)

	res := handler(ctx, msgs.CreateMsgIssue(&CoinIssueInfo))
	var issueID string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &issueID)

	bz := getQueried(t, ctx, querier, types.QueryIssue, issueID)

	var issueInfo types.CoinIssueInfo
	keeper.Getcdc().MustUnmarshalJSON(bz, &issueInfo)

	require.Equal(t, issueInfo.GetIssueId(), issueID)
	require.Equal(t, issueInfo.GetName(), CoinIssueInfo.GetName())

}

//func TestQueryIssues(t *testing.T) {
//	mapp, keeper, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)
//	mapp.BeginBlock(abci.RequestBeginBlock{})
//	ctx := mapp.NewContext(false, abci.Header{})
//	mapp.InitChainer(ctx, abci.RequestInitChain{})
//
//	querier := issue.NewQuerier(keeper)
//	handler := issue.NewHandler(keeper)
//
//	cap := 10
//	for i := 0; i < cap; i++ {
//		handler(ctx, msgs.CreateMsgIssue(&CoinIssueInfo))
//	}
//
//	bz := getQueried(t, ctx, querier, types.QueryIssues, CoinIssueInfo.Owner.String())
//
//	var coinIssues types.CoinIssues
//	keeper.Getcdc().MustUnmarshalJSON(bz, &coinIssues)
//
//	require.Equal(t, issueInfo.GetName(), CoinIssueInfo.GetName())
//
//}

//func getQueriedIssue(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, querierRoute string, issueID string) types.CoinIssueInfo {
//	query := abci.RequestQuery{
//		Path: queriers2.GetQueryIssuePath(issueID),
//		Data: nil,
//	}
//	bz, err := querier(ctx, []string{querierRoute, issueID}, query)
//	require.Nil(t, err)
//	require.NotNil(t, bz)
//
//	var issueInfo types.CoinIssueInfo
//	cdc.MustUnmarshalJSON(bz, &issueInfo)
//
//	return issueInfo
//}
func getQueried(t *testing.T, ctx sdk.Context, querier sdk.Querier, querierRoute string, queryPathParam string) (res []byte) {
	query := abci.RequestQuery{
		Path: queriers2.GetQueryIssuePath(queryPathParam),
		Data: nil,
	}
	bz, err := querier(ctx, []string{querierRoute, queryPathParam}, query)
	require.Nil(t, err)
	require.NotNil(t, bz)

	return bz
}
