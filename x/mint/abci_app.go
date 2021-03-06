package mint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Inflate every block, update inflation parameters once per hour
func BeginBlocker(ctx sdk.Context, k Keeper) {

	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

	// calculate annual provisions
	annualProvisions := minter.NextAnnualProvisions(params)

	// mint coins, add to collected fees, update supply
	mintedCoin := minter.BlockProvision(annualProvisions)
	k.fck.AddCollectedFees(ctx, sdk.Coins{mintedCoin})

	// first year do not inflate total supply
	// k.sk.InflateSupply(ctx, mintedCoin.Amount)
}
