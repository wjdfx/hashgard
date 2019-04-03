package keepers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"

	issueparams "github.com/hashgard/hashgard/x/issue/params"
	"github.com/hashgard/hashgard/x/issue/utils"

	"github.com/hashgard/hashgard/x/issue/domain"
)

// Parameter store key
var (
	ParamStoreKeyIssueParams = []byte("issueparams")
)

// Key declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamStoreKeyIssueParams, issueparams.IssueParams{},
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

func (keeper Keeper) Getcdc() *codec.Codec {
	return keeper.cdc
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
func (keeper Keeper) GetIssue(ctx sdk.Context, issueID string) *domain.CoinIssueInfo {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyIssuer(issueID))
	if len(bz) == 0 {
		return nil
	}
	var coinIssueInfo domain.CoinIssueInfo
	keeper.cdc.UnmarshalBinaryLengthPrefixed(bz, &coinIssueInfo)
	coinIssueInfo.IssueId = issueID
	return &coinIssueInfo
}

func (keeper Keeper) AddIssue(ctx sdk.Context, coinIssueInfo *domain.CoinIssueInfo) (string, sdk.Coins, sdk.Tags, sdk.Error) {
	store := ctx.KVStore(keeper.storeKey)
	issueID := ""
	for {
		issueID = utils.GetIssueID()
		if !store.Has(KeyIssuer(issueID)) {
			break
		}
	}
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(coinIssueInfo)
	store.Set(KeyIssuer(issueID), bz)
	coin := sdk.Coin{Denom: issueID, Amount: coinIssueInfo.TotalSupply}
	coins, tags, error := keeper.ck.AddCoins(ctx, coinIssueInfo.Issuer, sdk.Coins{coin})
	return issueID, coins, tags, error
}

func (keeper Keeper) FinishMinting(ctx sdk.Context, issueID string) *domain.CoinIssueInfo {
	coinIssueInfo := keeper.GetIssue(ctx, issueID)
	coinIssueInfo.MintingFinished = true
	store := ctx.KVStore(keeper.storeKey)
	store.Set(KeyIssuer(issueID), keeper.cdc.MustMarshalBinaryLengthPrefixed(coinIssueInfo))
	return coinIssueInfo
}

func (keeper Keeper) CanMint(ctx sdk.Context, issueID string) bool {
	coinIssueInfo := keeper.GetIssue(ctx, issueID)
	return !coinIssueInfo.MintingFinished
}

func (keeper Keeper) Mint(ctx sdk.Context, coinIssueInfo *domain.CoinIssueInfo, amount sdk.Int, to sdk.AccAddress) (sdk.Coins, sdk.Tags, sdk.Error) {
	coin := sdk.Coin{Denom: coinIssueInfo.IssueId, Amount: amount}
	coins, tags, error := keeper.ck.AddCoins(ctx, to, sdk.Coins{coin})
	if error != nil {
		return coins, tags, error
	}
	coinIssueInfo.TotalSupply = coinIssueInfo.TotalSupply.Add(amount)
	store := ctx.KVStore(keeper.storeKey)
	store.Set(KeyIssuer(coinIssueInfo.IssueId), keeper.cdc.MustMarshalBinaryLengthPrefixed(coinIssueInfo))

	return coins, tags, error
}
func (keeper Keeper) Burn(ctx sdk.Context, coinIssueInfo *domain.CoinIssueInfo, amount sdk.Int, who sdk.AccAddress) (sdk.Coins, sdk.Tags, sdk.Error) {
	coin := sdk.Coin{Denom: coinIssueInfo.IssueId, Amount: amount}
	coins, tags, error := keeper.ck.SubtractCoins(ctx, who, sdk.Coins{coin})
	if error != nil {
		return coins, tags, error
	}
	coinIssueInfo.TotalSupply = coinIssueInfo.TotalSupply.Sub(amount)
	store := ctx.KVStore(keeper.storeKey)
	store.Set(KeyIssuer(coinIssueInfo.IssueId), keeper.cdc.MustMarshalBinaryLengthPrefixed(coinIssueInfo))
	return coins, tags, error
}
func (keeper Keeper) SendCoins(ctx sdk.Context,
	fromAddr sdk.AccAddress, toAddr sdk.AccAddress,
	amt sdk.Coins) (sdk.Tags, sdk.Error) {
	return keeper.ck.SendCoins(ctx, fromAddr, toAddr, amt)
}

// Params
// Returns the current IssueParams from the global param store
func (keeper Keeper) GetIssueConfigParams(ctx sdk.Context) issueparams.IssueConfigParams {
	var issueConfigParams issueparams.IssueConfigParams
	keeper.paramSpace.Get(ctx, ParamStoreKeyIssueParams, &issueConfigParams)
	return issueConfigParams
}
func (keeper Keeper) SetIssueConfigParams(ctx sdk.Context, issueConfigParams issueparams.IssueConfigParams) {
	keeper.paramSpace.Set(ctx, ParamStoreKeyIssueParams, &issueConfigParams)
}
