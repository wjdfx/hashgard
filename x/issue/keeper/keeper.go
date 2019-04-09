package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"time"

	issueparams "github.com/hashgard/hashgard/x/issue/params"
	"github.com/hashgard/hashgard/x/issue/types"
	"github.com/hashgard/hashgard/x/issue/utils"
)

// Parameter store key
var (
	ParamStoreKeyIssueParams = []byte("issueparams")
)

// Key declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamStoreKeyIssueParams, issueparams.IssueConfigParams{},
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

//Get issue codec
func (keeper Keeper) Getcdc() *codec.Codec {
	return keeper.cdc
}

//New issue keeper Instance
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

//Returns issue by issueID
func (keeper Keeper) GetIssue(ctx sdk.Context, issueID string) *types.CoinIssueInfo {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyIssuer(issueID))
	if len(bz) == 0 {
		return nil
	}
	var coinIssueInfo types.CoinIssueInfo
	keeper.cdc.UnmarshalBinaryLengthPrefixed(bz, &coinIssueInfo)
	coinIssueInfo.IssueId = issueID
	return &coinIssueInfo
}

//Returns issues by accAddress
func (keeper Keeper) GetIssues(ctx sdk.Context, accAddress string) []*types.CoinIssueInfo {

	idAdders := keeper.GetAddressIssues(ctx, accAddress)
	length := len(idAdders)
	if length == 0 {
		return []*types.CoinIssueInfo{}
	}
	issues := make([]*types.CoinIssueInfo, 0, length)

	for _, v := range idAdders {
		issues = append(issues, keeper.GetIssue(ctx, v))
	}

	return issues
}

//Add a issue
func (keeper Keeper) AddIssue(ctx sdk.Context, coinIssueInfo *types.CoinIssueInfo) (sdk.Coins, sdk.Tags, sdk.Error) {
	store := ctx.KVStore(keeper.storeKey)
	issueID := ""
	for {
		issueID = utils.GetIssueID()
		if !store.Has(KeyIssuer(issueID)) {
			break
		}
	}

	coinIssueInfo.IssueTime = time.Now()
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(coinIssueInfo)
	store.Set(KeyIssuer(issueID), bz)

	idAdders := keeper.GetAddressIssues(ctx, coinIssueInfo.GetIssuer().String())
	idAdders = append(idAdders, issueID)
	keeper.setAddressIssues(ctx, coinIssueInfo.GetIssuer().String(), idAdders)

	coin := sdk.Coin{Denom: issueID, Amount: coinIssueInfo.TotalSupply}
	coins, tags, err := keeper.ck.AddCoins(ctx, coinIssueInfo.Issuer, sdk.Coins{coin})
	coinIssueInfo.IssueId = issueID
	return coins, tags, err
}

//Finished Minting a issue
func (keeper Keeper) FinishMinting(ctx sdk.Context, issueID string) *types.CoinIssueInfo {
	coinIssueInfo := keeper.GetIssue(ctx, issueID)
	coinIssueInfo.MintingFinished = true
	store := ctx.KVStore(keeper.storeKey)
	store.Set(KeyIssuer(issueID), keeper.cdc.MustMarshalBinaryLengthPrefixed(coinIssueInfo))
	return coinIssueInfo
}

//Can mint a issue
func (keeper Keeper) CanMint(ctx sdk.Context, issueID string) bool {
	coinIssueInfo := keeper.GetIssue(ctx, issueID)
	return !coinIssueInfo.MintingFinished
}

//Mint a issue
func (keeper Keeper) Mint(ctx sdk.Context, coinIssueInfo *types.CoinIssueInfo, amount sdk.Int, to sdk.AccAddress) (sdk.Coins, sdk.Tags, sdk.Error) {
	coin := sdk.Coin{Denom: coinIssueInfo.IssueId, Amount: amount}
	coins, tags, err := keeper.ck.AddCoins(ctx, to, sdk.Coins{coin})
	if err != nil {
		return coins, tags, err
	}
	coinIssueInfo.TotalSupply = coinIssueInfo.TotalSupply.Add(amount)
	store := ctx.KVStore(keeper.storeKey)
	store.Set(KeyIssuer(coinIssueInfo.IssueId), keeper.cdc.MustMarshalBinaryLengthPrefixed(coinIssueInfo))

	return coins, tags, err
}

//Burn a issue
func (keeper Keeper) Burn(ctx sdk.Context, coinIssueInfo *types.CoinIssueInfo, amount sdk.Int, who sdk.AccAddress) (sdk.Coins, sdk.Tags, sdk.Error) {
	coin := sdk.Coin{Denom: coinIssueInfo.IssueId, Amount: amount}
	coins, tags, err := keeper.ck.SubtractCoins(ctx, who, sdk.Coins{coin})
	if err != nil {
		return coins, tags, err
	}
	coinIssueInfo.TotalSupply = coinIssueInfo.TotalSupply.Sub(amount)
	store := ctx.KVStore(keeper.storeKey)
	store.Set(KeyIssuer(coinIssueInfo.IssueId), keeper.cdc.MustMarshalBinaryLengthPrefixed(coinIssueInfo))
	return coins, tags, err
}

//Send coins
func (keeper Keeper) SendCoins(ctx sdk.Context,
	fromAddr sdk.AccAddress, toAddr sdk.AccAddress,
	amt sdk.Coins) (sdk.Tags, sdk.Error) {
	return keeper.ck.SendCoins(ctx, fromAddr, toAddr, amt)
}

//Get address from a issue
func (keeper Keeper) GetAddressIssues(ctx sdk.Context, accAddress string) (idAdders []string) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyAddressIssues(accAddress))
	if bz == nil {
		return []string{}
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &idAdders)
	return idAdders
}

//Set address to a issue
func (keeper Keeper) setAddressIssues(ctx sdk.Context, accAddress string, idAdders []string) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(idAdders)
	store.Set(KeyAddressIssues(accAddress), bz)
}

// Params
// Returns the current issueConfigParams from the global param store
func (keeper Keeper) GetIssueConfigParams(ctx sdk.Context) issueparams.IssueConfigParams {
	var issueConfigParams issueparams.IssueConfigParams
	keeper.paramSpace.Get(ctx, ParamStoreKeyIssueParams, &issueConfigParams)
	return issueConfigParams
}

//Set issueConfigParams
func (keeper Keeper) SetIssueConfigParams(ctx sdk.Context, issueConfigParams issueparams.IssueConfigParams) {
	keeper.paramSpace.Set(ctx, ParamStoreKeyIssueParams, &issueConfigParams)
}
