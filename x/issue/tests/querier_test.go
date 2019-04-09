package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue"
	queriers2 "github.com/hashgard/hashgard/x/issue/client/queriers"
	"github.com/hashgard/hashgard/x/issue/domain"
	"github.com/hashgard/hashgard/x/issue/msgs"
)

func TestQueries(t *testing.T) {
	mapp, keeper, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)
	mapp.BeginBlock(abci.RequestBeginBlock{})
	ctx := mapp.NewContext(false, abci.Header{})
	mapp.InitChainer(ctx, abci.RequestInitChain{})

	querier := issue.NewQuerier(keeper)
	handler := issue.NewHandler(keeper)

	res := handler(ctx, msgs.NewMsgIssue(&CoinIssueInfo))
	var issueID string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &issueID)

	issueInfo := getQueriedIssue(t, ctx, keeper.Getcdc(), querier, issueID)
	require.Equal(t, issueInfo.GetIssueId(), issueID)
	require.Equal(t, issueInfo.GetName(), CoinIssueInfo.GetName())

}
func getQueriedIssue(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, issueID string) domain.CoinIssueInfo {
	query := abci.RequestQuery{
		Path: queriers2.GetQueryIssuePath(issueID, domain.StoreKey),
		Data: nil,
	}
	bz, err := querier(ctx, []string{domain.StoreKey, issueID}, query)
	require.Nil(t, err)
	require.NotNil(t, bz)

	var issueInfo domain.CoinIssueInfo
	require.Nil(t, cdc.UnmarshalJSON(bz, &issueInfo))

	return issueInfo
}
