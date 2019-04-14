package faucet

import (
	codec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// expected bank keeper
type BankKeeper interface {
	GetCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

	// TODO: remove once exchange doesn't require use of accounts
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) (sdk.Tags, sdk.Error)
}

type Keeper struct {
	cdc          *codec.Codec
	paramSpace   params.Subspace
	bankKeeper   BankKeeper
	codespace    sdk.CodespaceType
}

func NewKeeper(cdc *codec.Codec, paramSpace params.Subspace, ck BankKeeper, codespace sdk.CodespaceType) Keeper {
	return Keeper {
		cdc:		cdc,
		paramSpace:	paramSpace,
		bankKeeper:	ck,
		codespace:	codespace,
	}
}

func (keeper Keeper) GetFaucetOrigin(ctx sdk.Context) sdk.AccAddress {
	var faucetOrigin sdk.AccAddress
	keeper.paramSpace.Get(ctx, ParamStoreKeyFaucetOrigin, &faucetOrigin)
	return faucetOrigin
}

func (keeper Keeper) GetSendCoins(ctx sdk.Context) sdk.Coins {
	var sendCoins sdk.Coins
	keeper.paramSpace.Get(ctx, ParamStoreKeySendCoins, &sendCoins)
	return sendCoins
}

func (keeper Keeper) GetLimitCoins(ctx sdk.Context) sdk.Coins {
	var limitCoins sdk.Coins
	keeper.paramSpace.Get(ctx, ParamStoreKeySendCoins, &limitCoins)
	return limitCoins
}

func (keeper Keeper) setFaucetOrigin(ctx sdk.Context, faucetOrigin sdk.AccAddress) {
	keeper.paramSpace.Set(ctx, ParamStoreKeyFaucetOrigin, &faucetOrigin)
}

func (keeper Keeper) setSendCoins(ctx sdk.Context, sendCoins sdk.Coins) {
	keeper.paramSpace.Set(ctx, ParamStoreKeySendCoins, &sendCoins)
}

func (keeper Keeper) setLimitCoins(ctx sdk.Context, limitCoins sdk.Coins) {
	keeper.paramSpace.Set(ctx, ParamStoreKeyLimitCoins, &limitCoins)
}