package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

// type declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamStoreKeyCommunityTax, sdk.Dec{},
		ParamStoreKeyBaseProposerReward, sdk.Dec{},
		ParamStoreKeyBonusProposerReward, sdk.Dec{},
		ParamStoreKeyWithdrawAddrEnabled, false,
		ParamStoreKeyFoundationAddress, sdk.AccAddress{},
	)
}

// returns the current CommunityTax rate from the global param store
// nolint: errcheck
func (k Keeper) GetCommunityTax(ctx sdk.Context) sdk.Dec {
	var percent sdk.Dec
	k.paramSpace.Get(ctx, ParamStoreKeyCommunityTax, &percent)
	return percent
}

// nolint: errcheck
func (k Keeper) SetCommunityTax(ctx sdk.Context, percent sdk.Dec) {
	k.paramSpace.Set(ctx, ParamStoreKeyCommunityTax, &percent)
}

// returns the current BaseProposerReward rate from the global param store
// nolint: errcheck
func (k Keeper) GetBaseProposerReward(ctx sdk.Context) sdk.Dec {
	var percent sdk.Dec
	k.paramSpace.Get(ctx, ParamStoreKeyBaseProposerReward, &percent)
	return percent
}

// nolint: errcheck
func (k Keeper) SetBaseProposerReward(ctx sdk.Context, percent sdk.Dec) {
	k.paramSpace.Set(ctx, ParamStoreKeyBaseProposerReward, &percent)
}

// returns the current BaseProposerReward rate from the global param store
// nolint: errcheck
func (k Keeper) GetBonusProposerReward(ctx sdk.Context) sdk.Dec {
	var percent sdk.Dec
	k.paramSpace.Get(ctx, ParamStoreKeyBonusProposerReward, &percent)
	return percent
}

// nolint: errcheck
func (k Keeper) SetBonusProposerReward(ctx sdk.Context, percent sdk.Dec) {
	k.paramSpace.Set(ctx, ParamStoreKeyBonusProposerReward, &percent)
}

// returns the current WithdrawAddrEnabled
// nolint: errcheck
func (k Keeper) GetWithdrawAddrEnabled(ctx sdk.Context) bool {
	var enabled bool
	k.paramSpace.Get(ctx, ParamStoreKeyWithdrawAddrEnabled, &enabled)
	return enabled
}

// nolint: errcheck
func (k Keeper) SetWithdrawAddrEnabled(ctx sdk.Context, enabled bool) {
	k.paramSpace.Set(ctx, ParamStoreKeyWithdrawAddrEnabled, &enabled)
}

// get the foundation address
// nolint: errcheck
func (k Keeper) GetFoundationAddress(ctx sdk.Context) sdk.AccAddress {
	var foundationAddress sdk.AccAddress
	k.paramSpace.Get(ctx, ParamStoreKeyFoundationAddress, &foundationAddress)
	return foundationAddress
}

// set the foundation address
// nolint: errcheck
func (k Keeper) SetFoundationAddress(ctx sdk.Context, foundationAddress sdk.AccAddress) {
	k.paramSpace.Set(ctx, ParamStoreKeyFoundationAddress, &foundationAddress)
}