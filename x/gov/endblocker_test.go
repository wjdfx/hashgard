package gov

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/hashgard/hashgard/x/box"
	"github.com/hashgard/hashgard/x/issue"

	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestTickExpiredDepositPeriod(t *testing.T) {
	mapp, keeper, _, addrs, _, _ := getMockApp(t, 10, GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	keeper.ck.SetSendEnabled(ctx, true)
	govHandler := NewHandler(keeper)

	inactiveQueue := keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	newProposalMsg := NewMsgSubmitProposal("Test", "test", ProposalTypeText, addrs[0], sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 5)}, []ProposalParam{}, TaxUsage{})

	res := govHandler(ctx, newProposalMsg)
	require.True(t, res.IsOK())

	inactiveQueue = keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	newHeader := ctx.BlockHeader()
	newHeader.Time = ctx.BlockHeader().Time.Add(time.Duration(1) * time.Second)
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	newHeader = ctx.BlockHeader()
	newHeader.Time = ctx.BlockHeader().Time.Add(keeper.GetDepositParams(ctx).MaxDepositPeriod)
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.True(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	EndBlocker(ctx, keeper)

	inactiveQueue = keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()
}

func TestTickMultipleExpiredDepositPeriod(t *testing.T) {
	mapp, keeper, _, addrs, _, _ := getMockApp(t, 10, GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	keeper.ck.SetSendEnabled(ctx, true)
	govHandler := NewHandler(keeper)

	inactiveQueue := keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	newProposalMsg := NewMsgSubmitProposal("Test", "test", ProposalTypeText, addrs[0], sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 5)}, []ProposalParam{}, TaxUsage{})

	res := govHandler(ctx, newProposalMsg)
	require.True(t, res.IsOK())

	inactiveQueue = keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	newHeader := ctx.BlockHeader()
	newHeader.Time = ctx.BlockHeader().Time.Add(time.Duration(2) * time.Second)
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	newProposalMsg2 := NewMsgSubmitProposal("Test2", "test2", ProposalTypeText, addrs[1], sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 5)}, []ProposalParam{}, TaxUsage{})
	res = govHandler(ctx, newProposalMsg2)
	require.True(t, res.IsOK())

	newHeader = ctx.BlockHeader()
	newHeader.Time = ctx.BlockHeader().Time.Add(keeper.GetDepositParams(ctx).MaxDepositPeriod).Add(time.Duration(-1) * time.Second)
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.True(t, inactiveQueue.Valid())
	inactiveQueue.Close()
	EndBlocker(ctx, keeper)
	inactiveQueue = keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	newHeader = ctx.BlockHeader()
	newHeader.Time = ctx.BlockHeader().Time.Add(time.Duration(5) * time.Second)
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.True(t, inactiveQueue.Valid())
	inactiveQueue.Close()
	EndBlocker(ctx, keeper)
	inactiveQueue = keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()
}

func TestTickPassedDepositPeriod(t *testing.T) {
	mapp, keeper, _, addrs, _, _ := getMockApp(t, 10, GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	keeper.ck.SetSendEnabled(ctx, true)
	govHandler := NewHandler(keeper)

	inactiveQueue := keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()
	activeQueue := keeper.ActiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, activeQueue.Valid())
	activeQueue.Close()

	newProposalMsg := NewMsgSubmitProposal("Test", "test", ProposalTypeText, addrs[0], sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 5)}, []ProposalParam{}, TaxUsage{})

	res := govHandler(ctx, newProposalMsg)
	require.True(t, res.IsOK())
	var proposalID uint64
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(res.Data, &proposalID)

	inactiveQueue = keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	newHeader := ctx.BlockHeader()
	newHeader.Time = ctx.BlockHeader().Time.Add(time.Duration(1) * time.Second)
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	newDepositMsg := NewMsgDeposit(addrs[1], proposalID, sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 5)})
	res = govHandler(ctx, newDepositMsg)
	require.True(t, res.IsOK())

	activeQueue = keeper.ActiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, activeQueue.Valid())
	activeQueue.Close()
}

func TestTickPassedVotingPeriod(t *testing.T) {
	mapp, keeper, _, addrs, _, _ := getMockApp(t, 10, GenesisState{}, nil)
	SortAddresses(addrs)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	keeper.ck.SetSendEnabled(ctx, true)
	govHandler := NewHandler(keeper)

	inactiveQueue := keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()
	activeQueue := keeper.ActiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, activeQueue.Valid())
	activeQueue.Close()

	proposalCoins := sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, sdk.TokensFromTendermintPower(5))}
	newProposalMsg := NewMsgSubmitProposal("Test", "test", ProposalTypeText, addrs[0], proposalCoins, []ProposalParam{}, TaxUsage{})

	res := govHandler(ctx, newProposalMsg)
	require.True(t, res.IsOK())
	var proposalID uint64
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(res.Data, &proposalID)

	newHeader := ctx.BlockHeader()
	newHeader.Time = ctx.BlockHeader().Time.Add(time.Duration(1) * time.Second)
	ctx = ctx.WithBlockHeader(newHeader)

	newDepositMsg := NewMsgDeposit(addrs[1], proposalID, proposalCoins)
	res = govHandler(ctx, newDepositMsg)
	require.True(t, res.IsOK())

	newHeader = ctx.BlockHeader()
	newHeader.Time = ctx.BlockHeader().Time.Add(keeper.GetDepositParams(ctx).MaxDepositPeriod).Add(keeper.GetVotingParams(ctx).VotingPeriod)
	ctx = ctx.WithBlockHeader(newHeader)

	inactiveQueue = keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()

	activeQueue = keeper.ActiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.True(t, activeQueue.Valid())
	var activeProposalID uint64
	keeper.cdc.UnmarshalBinaryLengthPrefixed(activeQueue.Value(), &activeProposalID)
	proposal, ok := keeper.GetProposal(ctx, activeProposalID)
	require.True(t, ok)
	require.Equal(t, StatusVotingPeriod, proposal.Status)
	depositsIterator := keeper.GetDeposits(ctx, proposalID)
	require.True(t, depositsIterator.Valid())
	depositsIterator.Close()
	activeQueue.Close()

	EndBlocker(ctx, keeper)

	activeQueue = keeper.ActiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, activeQueue.Valid())
	activeQueue.Close()
}

func TestParameterChangePassedVotingPeriod(t *testing.T) {
	mapp, keeper, _, boxKeeper, issueKeeper, addrs, _, _ := getMockAppParams(t, 10, GenesisState{}, nil)
	SortAddresses(addrs)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})
	keeper.ck.SetSendEnabled(ctx, true)
	govHandler := NewHandler(keeper)

	inactiveQueue := keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, inactiveQueue.Valid())
	inactiveQueue.Close()
	activeQueue := keeper.ActiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	require.False(t, activeQueue.Valid())
	activeQueue.Close()

	niceVal := "9"
	niceCoin, _ := sdk.ParseCoin(niceVal + sdk.DefaultBondDenom)

	proposalParam := []ProposalParam{
		{Key: communityTax, Value: niceVal},
		{Key: minDeposit, Value: niceVal + sdk.DefaultBondDenom},
		{Key: inflation, Value: niceVal},
		{Key: inflationBase, Value: niceVal},
		{Key: signedBlocksWindow, Value: niceVal},
		{Key: minSignedPerWindow, Value: niceVal},
		{Key: downtimeJailDuration, Value: niceVal + "s"},
		{Key: slashFractionDowntime, Value: niceVal},
		{Key: unbondingTime, Value: niceVal + "s"},
		{Key: maxValidators, Value: niceVal},
		{Key: foundationAddress, Value: addrs[0].String()},
	}

	boxDefaultParams := box.DefaultParams(sdk.DefaultBondDenom)
	jsonStr, _ := json.Marshal(boxDefaultParams)
	proposalParam = appendFeeParams(jsonStr, BoxModule, proposalParam, niceCoin)

	issueDefaultParams := issue.DefaultParams(sdk.DefaultBondDenom)
	jsonStr, _ = json.Marshal(issueDefaultParams)
	proposalParam = appendFeeParams(jsonStr, IssueModule, proposalParam, niceCoin)

	proposalCoins := sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, sdk.TokensFromTendermintPower(5))}
	newProposalMsg := NewMsgSubmitProposal("Test", "test", ProposalTypeParameterChange, addrs[0], proposalCoins, proposalParam, TaxUsage{})
	fmt.Println(string(newProposalMsg.GetSignBytes()))
	res := govHandler(ctx, newProposalMsg)
	require.True(t, res.IsOK())
	var proposalID uint64
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(res.Data, &proposalID)

	proposal, _ := keeper.GetProposal(ctx, proposalID)
	proposal.Status = StatusPassed
	keeper.SetProposal(ctx, proposal)

	err := keeper.ExecuteProposal(ctx, proposal)
	require.NoError(t, err)

	boxParams := boxKeeper.GetParams(ctx)
	issueParams := issueKeeper.GetParams(ctx)

	require.Equal(t, boxParams.LockCreateFee.Amount, niceCoin.Amount)
	require.Equal(t, issueParams.IssueFee.Amount, niceCoin.Amount)
}

func appendFeeParams(json []byte, module string, proposalParam []ProposalParam, niceCoin sdk.Coin) []ProposalParam {
	content := string(json)
	content = strings.ReplaceAll(content, "{", "")
	content = strings.ReplaceAll(content, "}", "")
	content = strings.ReplaceAll(content, "\"", "")
	keys := strings.Split(content, ",")
	for _, key := range keys {
		strs := strings.Split(key, ":")
		if strings.HasSuffix(strs[0], strings.ToLower(Fee)) {
			proposalParam = append(proposalParam, ProposalParam{Key: module + strs[0], Value: niceCoin.String()})
		}
	}
	return proposalParam
}
