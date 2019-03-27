package issue

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

const (
	// ModuleKey is the name of the module
	ModuleName = "issue"
	// StoreKey is the store key string for issue
	StoreKey = ModuleName
	// RouterKey is the message route for issue
	RouterKey = ModuleName
	// QuerierRoute is the querier route for issue
	QuerierRoute = ModuleName
	// Parameter store default namestore
	DefaultParamspace = ModuleName
)

// Parameter store key
var (
	ParamStoreKeyIssueParams = []byte("issueparams")
)

// Key declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamStoreKeyIssueParams, IssueParams{},
	)
}

// Issue Keeper
type Keeper struct {
	// The reference to the Param Keeper to get and set Global Params
	paramsKeeper params.Keeper
	// The reference to the Paramstore to get and set issue specific params
	paramSpace params.Subspace
	// The (unexposed) keys used to access the stores from the Context.
	storeKey sdk.StoreKey
	// The reference to the CoinKeeper to modify balances
	ck BankKeeper
	// The codec codec for binary encoding/decoding.
	cdc *codec.Codec
	// Reserved codespace
	codespace sdk.CodespaceType
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramsKeeper params.Keeper,
	paramSpace params.Subspace, ck BankKeeper, codespace sdk.CodespaceType) Keeper {
	return Keeper{
		storeKey:     key,
		paramsKeeper: paramsKeeper,
		paramSpace:   paramSpace.WithKeyTable(ParamKeyTable()),
		ck:           ck,
		cdc:          cdc,
		codespace:    codespace,
	}
}
func (keeper Keeper) GetIssue(ctx sdk.Context, issueID string) *CoinIssueInfo {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyIssuer(issueID))
	var coinIssueInfo CoinIssueInfo
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &coinIssueInfo)
	return &coinIssueInfo
}

func (keeper Keeper) AddIssue(ctx sdk.Context, msgIssue MsgIssue) (string, sdk.Coins, sdk.Tags, sdk.Error) {
	store := ctx.KVStore(keeper.storeKey)
	issueID := ""
	for {
		issueID = GetIssueID()
		if !store.Has(KeyIssuer(issueID)) {
			break
		}
	}
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(msgIssue.CoinIssueInfo)
	store.Set(KeyIssuer(issueID), bz)
	coin := sdk.Coin{Denom: issueID, Amount: msgIssue.CoinIssueInfo.TotalSupply}
	coins, tags, error := keeper.ck.AddCoins(ctx, msgIssue.Issuer, sdk.Coins{coin})
	return issueID, coins, tags, error
}

func (keeper Keeper) FinishMinting(ctx sdk.Context, issueID string) bool {
	coinIssueInfo := keeper.GetIssue(ctx, issueID)
	coinIssueInfo.MintingFinished = true
	store := ctx.KVStore(keeper.storeKey)
	store.Set(KeyIssuer(issueID), keeper.cdc.MustMarshalBinaryLengthPrefixed(coinIssueInfo))
	return true
}

func (keeper Keeper) CanMint(ctx sdk.Context, issueID string) bool {
	coinIssueInfo := keeper.GetIssue(ctx, issueID)
	return !coinIssueInfo.MintingFinished
}

func (keeper Keeper) Mint(ctx sdk.Context, issueID string, amount sdk.Int, to sdk.AccAddress) *CoinIssueInfo {
	coinIssueInfo := keeper.GetIssue(ctx, issueID)
	coinIssueInfo.TotalSupply = coinIssueInfo.TotalSupply.Add(amount)
	store := ctx.KVStore(keeper.storeKey)
	store.Set(KeyIssuer(issueID), keeper.cdc.MustMarshalBinaryLengthPrefixed(coinIssueInfo))
	coin := sdk.Coin{Denom: issueID, Amount: amount}
	_, _, err := keeper.ck.AddCoins(ctx, to, sdk.Coins{coin})
	if err != nil {
		panic(err)
	}
	return coinIssueInfo
	//return sdk.Result{
	//	Tags: sdk.NewTags(
	//		tags.IssueID, issueID,
	//		tags.TotalSupply, coinIssueInfo.TotalSupply.String(),
	//	),
	//}
}
func (keeper Keeper) Burn(ctx sdk.Context, issueID string, amount sdk.Int, who sdk.AccAddress) bool {
	coinIssueInfo := keeper.GetIssue(ctx, issueID)
	coinIssueInfo.TotalSupply = coinIssueInfo.TotalSupply.Sub(amount)
	store := ctx.KVStore(keeper.storeKey)
	store.Set(KeyIssuer(issueID), keeper.cdc.MustMarshalBinaryLengthPrefixed(coinIssueInfo))
	coin := sdk.Coin{Denom: issueID, Amount: amount}
	_, _, err := keeper.ck.SubtractCoins(ctx, who, sdk.Coins{coin})
	if err != nil {
		panic(err)
	}
	return true
}

//func (keeper IssueKeeper) hasCoin(ctx sdk.Context, coinName string,issuerAddr sdk.AccAddress) bool{
//	store := ctx.KVStore(keeper.storeKey)
//
//	iterator:= sdk.KVStorePrefixIterator(store, KeyIssuer(coinName,issuerAddr))
//
//	defer iterator.Close()
//
//	return iterator.Valid() && len(iterator.Key())>0
//
//}
func (keeper Keeper) SendCoins(ctx sdk.Context,
	fromAddr sdk.AccAddress, toAddr sdk.AccAddress,
	amt sdk.Coins) (sdk.Tags, sdk.Error) {
	return keeper.ck.SendCoins(ctx, fromAddr, toAddr, amt)
}

// Params
// Returns the current IssueParams from the global param store
func (keeper Keeper) GetIssueParams(ctx sdk.Context) IssueParams {
	var IssueParams IssueParams
	keeper.paramSpace.Get(ctx, ParamStoreKeyIssueParams, &IssueParams)
	return IssueParams
}
func (keeper Keeper) setIssueParams(ctx sdk.Context, IssueParams IssueParams) {
	keeper.paramSpace.Set(ctx, ParamStoreKeyIssueParams, &IssueParams)
}
