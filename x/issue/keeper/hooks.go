package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Wrapper struct
type Hooks struct {
	keeper Keeper
}

// Create new issue hooks
func (keeper Keeper) Hooks() Hooks { return Hooks{keeper} }

func (hooks Hooks) CanSend(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) (bool, sdk.Error) {
	return true, nil
}