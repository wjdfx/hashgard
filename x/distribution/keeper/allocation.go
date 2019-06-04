package keeper

import (
	"fmt"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/distribution/types"
)

// allocate fees handles distribution of the collected fees
func (k Keeper) AllocateTokens(ctx sdk.Context, sumPreviousPrecommitPower, totalPreviousPower int64,
	previousProposer sdk.ConsAddress, previousVotes []abci.VoteInfo) {

	logger := ctx.Logger().With("module", "x/distribution")

	// fetch and clear the collected fees for distribution, since this is
	// called in BeginBlock, collected fees will be from the previous block
	// (and distributed to the previous proposer)
	feesCollectedInt := k.feeCollectionKeeper.GetCollectedFees(ctx)
	feesCollected := sdk.NewDecCoins(feesCollectedInt)
	k.feeCollectionKeeper.ClearCollectedFees(ctx)

	// temporary workaround to keep CanWithdrawInvariant happy
	// general discussions here: https://github.com/cosmos/cosmos-sdk/issues/2906#issuecomment-441867634
	feePool := k.GetFeePool(ctx)
	if totalPreviousPower == 0 {
		feePool.CommunityPool = feePool.CommunityPool.Add(feesCollected)
		k.SetFeePool(ctx, feePool)
		return
	}

	// calculate fraction votes
	previousFractionVotes := sdk.NewDec(sumPreviousPrecommitPower).Quo(sdk.NewDec(totalPreviousPower))

	// calculate previous proposer reward
	baseProposerReward := k.GetBaseProposerReward(ctx)
	bonusProposerReward := k.GetBonusProposerReward(ctx)
	proposerMultiplier := baseProposerReward.Add(bonusProposerReward.MulTruncate(previousFractionVotes))
	proposerReward := feesCollected.MulDecTruncate(proposerMultiplier)

	// pay previous proposer
	remaining := feesCollected
	proposerValidator := k.stakingKeeper.ValidatorByConsAddr(ctx, previousProposer)

	if proposerValidator != nil {
		k.AllocateTokensToValidator(ctx, proposerValidator, proposerReward)
		remaining = remaining.Sub(proposerReward)
	} else {
		// previous proposer can be unknown if say, the unbonding period is 1 block, so
		// e.g. a validator undelegates at block X, it's removed entirely by
		// block X+1's endblock, then X+2 we need to refer to the previous
		// proposer for X+1, but we've forgotten about them.
		logger.Error(fmt.Sprintf(
			"WARNING: Attempt to allocate proposer rewards to unknown proposer %s. "+
				"This should happen only if the proposer unbonded completely within a single block, "+
				"which generally should not happen except in exceptional circumstances (or fuzz testing). "+
				"We recommend you investigate immediately.",
			previousProposer.String()))
	}

	// calculate fraction allocated to validators
	communityTax := k.GetCommunityTax(ctx)
	voteMultiplier := sdk.OneDec().Sub(proposerMultiplier).Sub(communityTax)

	// allocate tokens proportionally to voting power
	// TODO consider parallelizing later, ref https://github.com/cosmos/cosmos-sdk/pull/3099#discussion_r246276376
	for _, vote := range previousVotes {
		validator := k.stakingKeeper.ValidatorByConsAddr(ctx, vote.Validator.Address)

		// TODO consider microslashing for missing votes.
		// ref https://github.com/cosmos/cosmos-sdk/issues/2525#issuecomment-430838701
		powerFraction := sdk.NewDec(vote.Validator.Power).QuoTruncate(sdk.NewDec(totalPreviousPower))
		reward := feesCollected.MulDecTruncate(voteMultiplier).MulDecTruncate(powerFraction)
		k.AllocateTokensToValidator(ctx, validator, reward)
		remaining = remaining.Sub(reward)
	}

	// allocate community funding
	feePool.CommunityPool = feePool.CommunityPool.Add(remaining)
	k.SetFeePool(ctx, feePool)

}

// allocate tokens to a particular validator, splitting according to commission
func (k Keeper) AllocateTokensToValidator(ctx sdk.Context, val sdk.Validator, tokens sdk.DecCoins) {

	// split tokens between validator and delegators according to commission
	commission := tokens.MulDec(val.GetCommission())
	shared := tokens.Sub(commission)

	// update current commission
	currentCommission := k.GetValidatorAccumulatedCommission(ctx, val.GetOperator())
	currentCommission = currentCommission.Add(commission)
	k.SetValidatorAccumulatedCommission(ctx, val.GetOperator(), currentCommission)

	// update current rewards
	currentRewards := k.GetValidatorCurrentRewards(ctx, val.GetOperator())
	currentRewards.Rewards = currentRewards.Rewards.Add(shared)
	k.SetValidatorCurrentRewards(ctx, val.GetOperator(), currentRewards)

	// update outstanding rewards
	outstanding := k.GetValidatorOutstandingRewards(ctx, val.GetOperator())
	outstanding = outstanding.Add(tokens)
	k.SetValidatorOutstandingRewards(ctx, val.GetOperator(), outstanding)
}

// Allocate the community pool from the community fee pool, burn or send to specific account
func (k Keeper) AllocateCommunityPool(ctx sdk.Context, destAddr sdk.AccAddress, percent sdk.Dec, burn bool) sdk.Error {
	logger := ctx.Logger()
	feePool := k.GetFeePool(ctx)
	communityPool := feePool.CommunityPool
	allocateCoins, _ := communityPool.MulDec(percent).TruncateDecimal()
	feePool.CommunityPool = communityPool.Sub(sdk.NewDecCoins(allocateCoins))
	k.SetFeePool(ctx, feePool)

	logger.Info("Spend community tax fund", "total_community_tax_fund", communityPool.String(), "left_community_tax_fund", feePool.CommunityPool.String())

	if burn {
		logger.Info("Burn community tax", "burn_amount", allocateCoins.String())
		if !allocateCoins.IsZero() {
			foundationAddress := k.GetFoundationAddress(ctx)
			_, err := k.bankKeeper.SubtractCoins(ctx, foundationAddress, allocateCoins)
			if err != nil {
				logger := ctx.Logger().With("module", "x/distr")
				logger.Info(fmt.Sprintf("the fund of foundation address(%s) is insufficient", foundationAddress))
				return types.ErrFoundationDryUp(types.DefaultCodespace)
			}
		}
	} else {
		logger.Info("Grant community tax to account", "grant_amount", allocateCoins.String(), "grant_address", destAddr.String())
		if !allocateCoins.IsZero() {
			foundationAddress := k.GetFoundationAddress(ctx)
			_, err := k.bankKeeper.SubtractCoins(ctx, foundationAddress, allocateCoins)
			if err != nil {
				logger := ctx.Logger().With("module", "x/distr")
				logger.Info(fmt.Sprintf("the fund of foundation address(%s) is insufficient", foundationAddress))
				return types.ErrFoundationDryUp(types.DefaultCodespace)
			}

			_, err = k.bankKeeper.AddCoins(ctx, destAddr, allocateCoins)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
