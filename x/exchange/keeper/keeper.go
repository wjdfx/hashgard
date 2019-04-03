package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/tendermint/tendermint/crypto"

	"github.com/hashgard/hashgard/x/exchange/msgs"
	"github.com/hashgard/hashgard/x/exchange/types"
)

var (
	ParamsStoreKeyExchangeParams = []byte("exchangeparams")

	// TODO: Find another way to implement this without using accounts, or find a cleaner way to implement it using accounts.
	FrozenCoinsAccAddr = sdk.AccAddress(crypto.AddressHash([]byte("exchangeFrozenCoins")))
)

func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamsStoreKeyExchangeParams, types.ExchangeParams{},
	)
}

type Keeper struct {
	storeKey 		sdk.StoreKey
	cdc				*codec.Codec
	paramsKeeper 	params.Keeper
	paramSpace		params.Subspace
	bankKeeper		types.BankKeeper
	codespace		sdk.CodespaceType
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramsKeeper params.Keeper,
	paramSpace params.Subspace, bankKeeper types.BankKeeper, codespace sdk.CodespaceType) Keeper {
		return Keeper{
			storeKey:		key,
			cdc:			cdc,
			paramsKeeper:	paramsKeeper,
			paramSpace:		paramSpace.WithKeyTable(ParamKeyTable()),
			bankKeeper:		bankKeeper,
			codespace:		codespace,
		}
}


func (keeper Keeper) CreateOrder(ctx sdk.Context,seller sdk.AccAddress,
	supply sdk.Coin, target sdk.Coin) (order types.Order, err sdk.Error) {
	orderId, err := keeper.getNewOrderId(ctx)
	if err != nil {
		return
	}

	createTime := ctx.BlockHeader().Time

	order = types.Order{
		OrderId:	orderId,
		Seller:		seller,
		Supply:		supply,
		Target:		target,
		Remains:	supply,
		CreateTime:	createTime,
	}

	_, err = keeper.bankKeeper.SendCoins(ctx, seller, FrozenCoinsAccAddr, []sdk.Coin{supply})
	if err != nil {
		return
	}

	keeper.setOrder(ctx, order)

	return order, nil
}


// Store level
func (keeper Keeper) GetOrder(ctx sdk.Context, orderId uint64) (types.Order, bool) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyOrder(orderId))
	if bz == nil {
		return types.Order{}, false
	}
	var order types.Order
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &order)
	return order, true
}

func (keeper Keeper) setOrder(ctx sdk.Context, order types.Order) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(order)
	store.Set(KeyOrder(order.OrderId), bz)
}

func (keeper Keeper) deleteOrder(ctx sdk.Context, orderId uint64) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(KeyOrder(orderId))
}

func (keeper Keeper) hasOrder(ctx sdk.Context, orderId uint64) bool {
	store := ctx.KVStore(keeper.storeKey)
	return store.Has(KeyOrder(orderId))
}

// Get the next available OrderId and increments it
func (keeper Keeper) getNewOrderId(ctx sdk.Context) (orderId uint64, err sdk.Error) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyNextOrderId)
	if bz == nil {
		return 0, sdk.NewError(keeper.codespace, types.CodeInvalidGenesis, "InitialOrderId never set")
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &orderId)
	bz = keeper.cdc.MustMarshalBinaryLengthPrefixed(orderId + 1)
	store.Set(KeyNextOrderId, bz)
	return orderId, nil
}



