package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/tendermint/tendermint/crypto"

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
	orderIdArr := keeper.GetAddressOrders(ctx, seller)
	orderIdArr = append(orderIdArr, orderId)
	keeper.setAddressOrders(ctx, seller, orderIdArr)

	return order, nil
}

func (keeper Keeper) WithdrawalOrder(ctx sdk.Context, orderId uint64, addr sdk.AccAddress) (amt sdk.Coin, err sdk.Error) {
	order, ok := keeper.GetOrder(ctx, orderId)
	if !ok {
		return amt, sdk.NewError(keeper.codespace, types.CodeOrderNotExist, fmt.Sprintf("this orderId is invalid : %d", orderId))
	}

	if !order.Seller.Equals(addr) {
		return amt, sdk.NewError(keeper.codespace, types.CodeNoPermission, fmt.Sprintf("order %d isn't belong to %s", orderId, addr))
	}

	amt = order.Remains
	_, err = keeper.bankKeeper.SendCoins(ctx, FrozenCoinsAccAddr, addr, []sdk.Coin{amt})
	if err != nil {
		return
	}

	keeper.deleteOrder(ctx, orderId)

	orderIdArr := keeper.GetAddressOrders(ctx, addr)
	for index := 0; index < len(orderIdArr); {
		if orderIdArr[index] == orderId {
			orderIdArr = append(orderIdArr[:index], orderIdArr[index+1:]...)
			break
		}
		index++
	}
	keeper.setAddressOrders(ctx, addr, orderIdArr)

	return amt, nil
}

func (keeper Keeper) TakeOrder(ctx sdk.Context, orderId uint64, buyer sdk.AccAddress, val sdk.Coin) (supplyTurnover sdk.Coin,
	targetTurnover sdk.Coin, soldOut bool, err sdk.Error) {
	order, ok := keeper.GetOrder(ctx, orderId)
	if !ok {
		return supplyTurnover, targetTurnover, soldOut, sdk.NewError(keeper.codespace, types.CodeOrderNotExist, fmt.Sprintf("this orderId is invalid : %d", orderId))
	}

	if val.Denom != order.Target.Denom {
		return supplyTurnover, targetTurnover, soldOut, sdk.NewError(keeper.codespace, types.CodeNotMatchTarget, fmt.Sprintf("%s doesn't match order's target(%s)", val.Denom, order.Target.Denom))
	}

	divisor := GetGratestDivisor(order.Supply.Amount, order.Target.Amount)
	remainShares := order.Remains.Amount.Quo(divisor)
	sharePrice := order.Target.Amount.Quo(divisor)

	if val.Amount.LT(sharePrice) {
		return supplyTurnover, targetTurnover, soldOut, sdk.NewError(keeper.codespace, types.CodeTooLess, fmt.Sprintf("minimum purchase threshold is %s%s", sharePrice.String(), order.Target.Denom))
	}

	shares := val.Amount.Quo(sharePrice)

	if shares.GTE(remainShares) {
		shares = remainShares
		soldOut = true
	}

	supplyTurnover = sdk.NewCoin(order.Supply.Denom, order.Supply.Amount.Quo(divisor).Mul(shares))
	targetTurnover = sdk.NewCoin(order.Target.Denom, sharePrice.Mul(shares))

	_, err = keeper.bankKeeper.SendCoins(ctx, buyer, order.Seller, []sdk.Coin{targetTurnover})
	if err != nil {
		return
	}
	_, err = keeper.bankKeeper.SendCoins(ctx, FrozenCoinsAccAddr, buyer, []sdk.Coin{supplyTurnover})
	if err != nil {
		return
	}

	if soldOut {
		keeper.deleteOrder(ctx, orderId)
		orderIdArr := keeper.GetAddressOrders(ctx, order.Seller)
		for index := 0; index < len(orderIdArr); {
			if orderIdArr[index] == orderId {
				orderIdArr = append(orderIdArr[:index], orderIdArr[index+1:]...)
				break
			}
			index++
		}
		keeper.setAddressOrders(ctx, order.Seller, orderIdArr)
	} else {
		remains := order.Remains.Sub(supplyTurnover)
		newOrder := types.Order{
			OrderId:	orderId,
			Seller:		order.Seller,
			Supply:		order.Supply,
			Target:		order.Target,
			Remains:	remains,
			CreateTime:	order.CreateTime,
		}
		keeper.setOrder(ctx, newOrder)
	}

	return supplyTurnover, targetTurnover, soldOut, nil
}

func (keeper Keeper) GetOrdersByAddr(ctx sdk.Context, addr sdk.AccAddress) (orders types.Orders, err sdk.Error) {
	orderIdArr := keeper.GetAddressOrders(ctx, addr)
	for _, orderId := range orderIdArr {
		order, ok := keeper.GetOrder(ctx, orderId)
		if !ok {
			return types.Orders{}, sdk.NewError(keeper.codespace, types.CodeOrderNotExist, fmt.Sprintf("this orderId is invalid : %d", orderId))
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (keeper Keeper) GetFrozenFundByAddr(ctx sdk.Context, addr sdk.AccAddress) (fund sdk.Coins, err sdk.Error) {
	orderIdArr := keeper.GetAddressOrders(ctx, addr)
	for _, orderId := range orderIdArr {
		order, ok := keeper.GetOrder(ctx, orderId)
		if !ok {
			return sdk.Coins{}, sdk.NewError(keeper.codespace, types.CodeOrderNotExist, fmt.Sprintf("this orderId is invalid : %d", orderId))
		}
		fund = fund.Add([]sdk.Coin{order.Remains})
	}

	return fund, nil
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

func (keeper Keeper) HasOrder(ctx sdk.Context, orderId uint64) bool {
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

// Peeks the next available orderId without incrementing it
func (keeper Keeper) PeekCurrentOrderId(ctx sdk.Context) (orderId uint64, err sdk.Error) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyNextOrderId)
	if bz == nil {
		return 0, sdk.NewError(keeper.codespace, types.CodeInvalidGenesis, "InitialOrderId never set")
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &orderId)
	return orderId, nil
}


// Set the initial order ID
func (keeper Keeper) SetInitialOrderId(ctx sdk.Context, orderId uint64) sdk.Error {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyNextOrderId)
	if bz != nil {
		return sdk.NewError(keeper.codespace, types.CodeInvalidGenesis, "Initial ProposalID already set")
	}
	bz = keeper.cdc.MustMarshalBinaryLengthPrefixed(orderId)
	store.Set(KeyNextOrderId, bz)
	return nil
}

func (keeper Keeper) GetAddressOrders(ctx sdk.Context, addr sdk.AccAddress) (orderIdArr []uint64) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyAddressOrders(addr))
	if bz == nil {
		return []uint64{}
	}

	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &orderIdArr)
	return orderIdArr
}

func (keeper Keeper) setAddressOrders(ctx sdk.Context, addr sdk.AccAddress, orderIdArr []uint64) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(orderIdArr)
	store.Set(KeyAddressOrders(addr), bz)
}