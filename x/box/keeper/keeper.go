package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/hashgard/hashgard/x/box/errors"
	boxparams "github.com/hashgard/hashgard/x/box/params"
	"github.com/hashgard/hashgard/x/box/types"
	issueerr "github.com/hashgard/hashgard/x/issue/errors"
)

// Parameter store key
var (
	ParamStoreKeyBoxParams = []byte("boxparams")
)

// Key declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamStoreKeyBoxParams, boxparams.BoxConfigParams{},
	)
}

// Box Keeper
type Keeper struct {
	// The reference to the Param Keeper to get and set Global Params
	paramsKeeper params.Keeper
	// The reference to the Paramstore to get and set box specific params
	paramSpace params.Subspace
	// The (unexposed) keys used to access the stores from the Context.
	storeKey sdk.StoreKey
	// The reference to the CoinKeeper to modify balances
	ck BankKeeper
	// The reference to the IssueKeeper to get issue info
	ik IssueKeeper
	// The codec codec for binary encoding/decoding.
	cdc *codec.Codec
	// Reserved codespace
	codespace sdk.CodespaceType
}

//Get box codec
func (keeper Keeper) Getcdc() *codec.Codec {
	return keeper.cdc
}

//Get box bankKeeper
func (keeper Keeper) GetBankKeeper() BankKeeper {
	return keeper.ck
}

//Get box issueKeeper
func (keeper Keeper) GetIssueKeeper() IssueKeeper {
	return keeper.ik
}

//New box keeper Instance
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramsKeeper params.Keeper,
	paramSpace params.Subspace, ck BankKeeper, ik IssueKeeper, codespace sdk.CodespaceType) Keeper {
	return Keeper{
		storeKey:     key,
		paramsKeeper: paramsKeeper,
		paramSpace:   paramSpace.WithKeyTable(ParamKeyTable()),
		ck:           ck,
		ik:           ik,
		cdc:          cdc,
		codespace:    codespace,
	}
}

func (keeper Keeper) getDepositedCoinsAddress(boxID string) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(fmt.Sprintf("boxDepositedCoins:%s", boxID))))

}
func (keeper Keeper) SendDepositedCoin(ctx sdk.Context, fromAddr sdk.AccAddress, amt sdk.Coins, boxID string) sdk.Error {
	toAddr := keeper.getDepositedCoinsAddress(boxID)
	return keeper.GetBankKeeper().SendCoins(ctx, fromAddr, toAddr, amt)
}
func (keeper Keeper) FetchDepositedCoin(ctx sdk.Context, toAddr sdk.AccAddress, amt sdk.Coins, boxID string) sdk.Error {
	fromAddr := keeper.getDepositedCoinsAddress(boxID)
	return keeper.GetBankKeeper().SendCoins(ctx, fromAddr, toAddr, amt)
}
func (keeper Keeper) SubDepositedCoin(ctx sdk.Context, amt sdk.Coins, boxID string) sdk.Error {
	_, err := keeper.GetBankKeeper().SubtractCoins(ctx, keeper.getDepositedCoinsAddress(boxID), amt)
	return err
}
func (keeper Keeper) GetDepositedCoins(ctx sdk.Context, boxID string) sdk.Coins {
	return keeper.GetBankKeeper().GetCoins(ctx, keeper.getDepositedCoinsAddress(boxID))
}

//Keys set
//Set box
func (keeper Keeper) setBox(ctx sdk.Context, box *types.BoxInfo) {
	store := ctx.KVStore(keeper.storeKey)
	store.Set(KeyBox(box.BoxId), keeper.cdc.MustMarshalBinaryLengthPrefixed(box))
}

//Set address
func (keeper Keeper) setAddress(ctx sdk.Context, boxType string, accAddress sdk.AccAddress, boxIDs []string) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(boxIDs)
	store.Set(KeyAddress(boxType, accAddress), bz)
}

//Set name
func (keeper Keeper) setName(ctx sdk.Context, boxType string, name string, boxIDs []string) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(boxIDs)
	store.Set(KeyName(boxType, name), bz)
}

func (keeper Keeper) setAddressDeposit(ctx sdk.Context, boxID string, accAddress sdk.AccAddress, boxDeposit *types.BoxDeposit) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(boxDeposit)
	store.Set(KeyAddressDeposit(boxID, accAddress), bz)
}

//Add address deposit
func (keeper Keeper) addAddressDeposit(ctx sdk.Context, boxID string, accAddress sdk.AccAddress, boxDeposit *types.BoxDeposit) {
	boxDeposit.Amount = boxDeposit.Amount.Add(keeper.GetDepositByAddress(ctx, boxID, accAddress).Amount)
	keeper.setAddressDeposit(ctx, boxID, accAddress, boxDeposit)
}

//Keys remove
//Remove box
func (keeper Keeper) RemoveBox(ctx sdk.Context, box *types.BoxInfo) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(KeyBox(box.BoxId))
}

//Remove address deposit
func (keeper Keeper) removeAddressDeposit(ctx sdk.Context, boxID string, accAddress sdk.AccAddress) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(KeyAddressDeposit(boxID, accAddress))
}

//Return deposit amount by accAddress
func (keeper Keeper) GetDepositByAddress(ctx sdk.Context, boxID string, accAddress sdk.AccAddress) *types.BoxDeposit {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyAddressDeposit(boxID, accAddress))
	if bz == nil {
		return types.NewZeroBoxDeposit()
	}
	var boxDeposit types.BoxDeposit
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &boxDeposit)
	return &boxDeposit
}

func (keeper Keeper) GetCoinDecimal(ctx sdk.Context, coin sdk.Coin) (uint, sdk.Error) {
	if coin.Denom == types.Agard {
		return types.AgardDecimal, nil
	}
	coinIssueInfo := keeper.GetIssueKeeper().GetIssue(ctx, coin.Denom)
	if coinIssueInfo == nil {
		return 0, issueerr.ErrUnknownIssue(coin.Denom)
	}
	return coinIssueInfo.Decimals, nil
}

//Keys return
//Return box by boxID
func (keeper Keeper) GetBox(ctx sdk.Context, boxID string) *types.BoxInfo {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyBox(boxID))
	if len(bz) == 0 {
		return nil
	}
	var box types.BoxInfo
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &box)
	return &box
}

//Return box by boxID and and check owner
func (keeper Keeper) GetBoxByOwner(ctx sdk.Context, sender sdk.AccAddress, boxID string) (*types.BoxInfo, sdk.Error) {
	box := keeper.GetBox(ctx, boxID)
	if box == nil {
		return nil, errors.ErrUnknownBox(boxID)
	}
	if !box.Owner.Equals(sender) {
		return nil, errors.ErrOwnerMismatch(boxID)
	}
	return box, nil
}

//Returns box list by type and accAddress
func (keeper Keeper) GetBoxByAddress(ctx sdk.Context, boxType string, accAddress sdk.AccAddress) []*types.BoxInfo {
	boxIDs := keeper.GetBoxIdsByAddress(ctx, boxType, accAddress)
	length := len(boxIDs)
	if length == 0 {
		return []*types.BoxInfo{}
	}
	boxs := make([]*types.BoxInfo, 0, length)

	for _, v := range boxIDs {
		boxs = append(boxs, keeper.GetBox(ctx, v))
	}
	return boxs
}

func (keeper Keeper) CheckDepositByAddress(ctx sdk.Context, boxID string, accAddress sdk.AccAddress) bool {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyAddressDeposit(boxID, accAddress))
	if bz == nil {
		return false
	}
	return true
}

//Queries

//Search box by name
func (keeper Keeper) SearchBox(ctx sdk.Context, boxType string, name string) []*types.BoxInfo {
	store := ctx.KVStore(keeper.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, KeyName(boxType, name))
	defer iterator.Close()
	list := make([]*types.BoxInfo, 0, 1)
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		if len(bz) == 0 {
			continue
		}
		boxIDs := make([]string, 0, 1)
		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &boxIDs)

		for _, v := range boxIDs {
			list = append(list, keeper.GetBox(ctx, v))
		}
	}
	return list
}

//List
func (keeper Keeper) List(ctx sdk.Context, params boxparams.BoxQueryParams) []*types.BoxInfo {
	if params.Owner != nil && !params.Owner.Empty() {
		return keeper.GetBoxByAddress(ctx, params.BoxType, params.Owner)
	}
	store := ctx.KVStore(keeper.storeKey)
	startBoxId := params.StartBoxId
	endBoxId := startBoxId
	if len(startBoxId) == 0 {
		endBoxId = KeyBoxIdStr(params.BoxType, types.BoxMaxId)
		startBoxId = KeyBoxIdStr(params.BoxType, types.BoxMinId-1)
	} else {
		startBoxId = KeyBoxIdStr(params.BoxType, types.BoxMinId-1)
	}
	iterator := store.ReverseIterator(KeyBox(startBoxId), KeyBox(endBoxId))
	defer iterator.Close()
	list := make([]*types.BoxInfo, 0, params.Limit)
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		if len(bz) == 0 {
			continue
		}
		var info types.BoxInfo
		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &info)
		list = append(list, &info)
		if len(list) >= params.Limit {
			break
		}
	}
	return list
}

//Create a box
func (keeper Keeper) CreateBox(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	decimal, err := keeper.GetCoinDecimal(ctx, box.TotalAmount.Token)
	if err != nil {
		return err
	}
	if box.TotalAmount.Decimals != decimal {
		return errors.ErrDecimalsNotValid(box.TotalAmount.Decimals)
	}
	store := ctx.KVStore(keeper.storeKey)
	id, err := keeper.getNewBoxID(store, box.BoxType)
	if err != nil {
		return err
	}
	box.BoxId = KeyBoxIdStr(box.BoxType, id)
	box.CreatedTime = ctx.BlockHeader().Time.Unix()

	switch box.BoxType {
	case types.Lock:
		err = keeper.ProcessLockBoxCreate(ctx, box)
	case types.Deposit:
		err = keeper.ProcessDepositBoxCreate(ctx, box)
	case types.Future:
		err = keeper.ProcessFutureBoxCreate(ctx, box)
	default:
		return errors.ErrUnknownBoxType()
	}
	if err != nil {
		return err
	}
	boxIDs := keeper.GetBoxIdsByAddress(ctx, box.BoxType, box.Owner)
	boxIDs = append(boxIDs, box.BoxId)
	keeper.setAddress(ctx, box.BoxType, box.Owner, boxIDs)

	boxIDs = keeper.GetBoxIdsByName(ctx, box.BoxType, box.Name)
	boxIDs = append(boxIDs, box.BoxId)
	keeper.setName(ctx, box.BoxType, box.Name, boxIDs)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(box)
	store.Set(KeyBox(box.BoxId), bz)
	return nil
}

func (keeper Keeper) ProcessDepositToBox(ctx sdk.Context, boxID string, sender sdk.AccAddress, deposit sdk.Coin, operation string) (*types.BoxInfo, sdk.Error) {
	box := keeper.GetBox(ctx, boxID)
	if box == nil {
		return nil, errors.ErrUnknownBox(boxID)
	}
	if types.BoxDepositing != box.BoxStatus {
		return nil, errors.ErrNotAllowedOperation(box.BoxStatus)
	}
	switch box.BoxType {
	case types.Deposit:
		return box, keeper.processDepositBoxDeposit(ctx, box, sender, deposit, operation)
	case types.Future:
		return box, keeper.processFutureBoxDeposit(ctx, box, sender, deposit, operation)
	}
	return nil, errors.ErrUnknownBoxType()
}

func (keeper Keeper) SetBoxDescription(ctx sdk.Context, boxID string, sender sdk.AccAddress, description []byte) (*types.BoxInfo, sdk.Error) {
	box, err := keeper.GetBoxByOwner(ctx, sender, boxID)
	if err != nil {
		return box, err
	}
	box.Description = string(description)
	keeper.setBox(ctx, box)
	return box, nil
}
func (keeper Keeper) DisableFeature(ctx sdk.Context, sender sdk.AccAddress, boxID string, feature string) (*types.BoxInfo, sdk.Error) {
	boxInfo, err := keeper.GetBoxByOwner(ctx, sender, boxID)
	if err != nil {
		return nil, err
	}
	switch feature {
	case types.Trade:
		return boxInfo, keeper.disableTrade(ctx, sender, boxInfo)
	default:
		return nil, errors.ErrUnknownFeatures()
	}
}
func (keeper Keeper) disableTrade(ctx sdk.Context, sender sdk.AccAddress, boxInfo *types.BoxInfo) sdk.Error {
	if boxInfo.GetBoxType() == types.Lock {
		return errors.ErrNotSupportOperation()
	}
	if !boxInfo.IsTradeDisabled() {
		return nil
	}
	boxInfo.TradeDisabled = false
	keeper.setBox(ctx, boxInfo)
	return nil
}

//Send coins
func (keeper Keeper) SendCoins(ctx sdk.Context,
	fromAddr sdk.AccAddress, toAddr sdk.AccAddress,
	amt sdk.Coins) sdk.Error {
	return keeper.ck.SendCoins(ctx, fromAddr, toAddr, amt)
}

//Get name from a box
func (keeper Keeper) GetBoxIdsByName(ctx sdk.Context, boxType string, name string) (boxIDs []string) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyName(boxType, name))
	if bz == nil {
		return []string{}
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &boxIDs)
	return boxIDs
}

//Get address from a box
func (keeper Keeper) GetBoxIdsByAddress(ctx sdk.Context, boxType string, accAddress sdk.AccAddress) (boxIDs []string) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyAddress(boxType, accAddress))
	if bz == nil {
		return []string{}
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &boxIDs)
	return boxIDs
}

// BoxQueues

// Returns an iterator for all the box in the Active Queue that expire by time
func (keeper Keeper) ActiveBoxQueueIterator(ctx sdk.Context, endTime int64) sdk.Iterator {
	store := ctx.KVStore(keeper.storeKey)
	return store.Iterator(PrefixActiveQueue, sdk.PrefixEndBytes(PrefixActiveBoxQueueTime(endTime)))
}

// Inserts a boxID into the active box queue at time
func (keeper Keeper) InsertActiveBoxQueue(ctx sdk.Context, endTime int64, boxIdStr string) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(boxIdStr)
	store.Set(KeyActiveBoxQueue(endTime, boxIdStr), bz)
}

// removes a boxID from the Active box Queue
func (keeper Keeper) RemoveFromActiveBoxQueue(ctx sdk.Context, endTime int64, boxIdStr string) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(KeyActiveBoxQueue(endTime, boxIdStr))
}
func (keeper Keeper) RemoveFromActiveBoxQueueByKey(ctx sdk.Context, key []byte) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(key)
}

// Params
// Returns the current boxConfigParams from the global param store
func (keeper Keeper) GetBoxConfigParams(ctx sdk.Context) boxparams.BoxConfigParams {
	var boxConfigParams boxparams.BoxConfigParams
	keeper.paramSpace.Get(ctx, ParamStoreKeyBoxParams, &boxConfigParams)
	return boxConfigParams
}

//Set boxConfigParams
func (keeper Keeper) SetBoxConfigParams(ctx sdk.Context, boxConfigParams boxparams.BoxConfigParams) {
	keeper.paramSpace.Set(ctx, ParamStoreKeyBoxParams, &boxConfigParams)
}

//Set the initial boxCount
func (keeper Keeper) SetInitialBoxStartingBoxId(ctx sdk.Context, boxType string, boxID uint64) sdk.Error {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyNextBoxID(boxType))
	if bz != nil {
		return sdk.NewError(keeper.codespace, types.CodeInvalidGenesis, "Initial BoxId already set")
	}
	bz = keeper.cdc.MustMarshalBinaryLengthPrefixed(boxID)
	store.Set(KeyNextBoxID(boxType), bz)
	return nil
}

// Get the last used boxID
func (keeper Keeper) GetLastBoxID(ctx sdk.Context, boxType string) (boxID uint64) {
	boxID, err := keeper.PeekCurrentBoxID(ctx, boxType)
	if err != nil {
		return 0
	}
	boxID--
	return
}

// Gets the next available boxID and increments it
func (keeper Keeper) getNewBoxID(store sdk.KVStore, boxType string) (boxID uint64, err sdk.Error) {
	bz := store.Get(KeyNextBoxID(boxType))
	if bz == nil {
		return 0, sdk.NewError(keeper.codespace, types.CodeInvalidGenesis, "InitialBoxID never set")
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &boxID)
	bz = keeper.cdc.MustMarshalBinaryLengthPrefixed(boxID + 1)
	store.Set(KeyNextBoxID(boxType), bz)
	return boxID, nil
}

// Peeks the next available BoxID without incrementing it
func (keeper Keeper) PeekCurrentBoxID(ctx sdk.Context, boxType string) (boxID uint64, err sdk.Error) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyNextBoxID(boxType))
	if bz == nil {
		return 0, sdk.NewError(keeper.codespace, types.CodeInvalidGenesis, "InitialBoxID never set")
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &boxID)
	return boxID, nil
}
