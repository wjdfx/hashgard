package app

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banksim "github.com/cosmos/cosmos-sdk/x/bank/simulation"
	distributionsim "github.com/cosmos/cosmos-sdk/x/distribution/simulation"
	"github.com/cosmos/cosmos-sdk/x/mock/simulation"
	stakingsim "github.com/cosmos/cosmos-sdk/x/staking/simulation"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (app *HashgardApp) runtimeInvariants() []simulation.Invariant {
	return []simulation.Invariant{
		banksim.NonnegativeBalanceInvariant(app.accountKeeper),
		distributionsim.NonNegativeOutstandingInvariant(app.distributionKeeper),
		stakingsim.SupplyInvariants(app.bankKeeper, app.stakingKeeper,
			app.feeCollectionKeeper, app.distributionKeeper, app.accountKeeper),
		stakingsim.NonNegativePowerInvariant(app.stakingKeeper),
	}
}

func (app *HashgardApp) assertRuntimeInvariants() {
	ctx := app.NewContext(false, abci.Header{Height: app.LastBlockHeight() + 1})
	app.assertRuntimeInvariantsOnContext(ctx)
}

func (app *HashgardApp) assertRuntimeInvariantsOnContext(ctx sdk.Context) {
	start := time.Now()
	invariants := app.runtimeInvariants()
	for _, inv := range invariants {
		if err := inv(ctx); err != nil {
			fmt.Println(err)
			panic(fmt.Errorf("invariant broken: %s", err))
		}
	}
	end := time.Now()
	diff := end.Sub(start)
	app.BaseApp.Logger().With("module", "invariants").Info("Asserted all invariants", "duration", diff)
}
