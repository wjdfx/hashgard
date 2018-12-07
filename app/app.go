package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/stake"
)

const (
	appName = "HashgardApp"
	FlagReplay = "replay"
	DefaultKeyPass = "12345678"		// DefaultKeyPass contains key password for genesis transactions
)

// default home directories for expected binaries
var (
	DefaultCLIHome  = os.ExpandEnv("$HOME/.hashgardcli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.hashgard")
)

// Extended ABCI application
type HashgardApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	// keys to access the multistore
	keyMain				*sdk.KVStoreKey
	keyAccount			*sdk.KVStoreKey
	keyStake			*sdk.KVStoreKey
	tkeyStake			*sdk.TransientStoreKey
	keySlashing			*sdk.KVStoreKey
	keyMint				*sdk.KVStoreKey
	keyDistribution		*sdk.KVStoreKey
	tkeyDistribution	*sdk.TransientStoreKey
	keyFeeCollection	*sdk.KVStoreKey
	keyParams			*sdk.KVStoreKey
	tkeyParams			*sdk.TransientStoreKey

	// manage getting and setting accounts
	accountKeeper       auth.AccountKeeper
	feeCollectionKeeper auth.FeeCollectionKeeper
	bankKeeper          bank.Keeper
	stakeKeeper         stake.Keeper
	slashingKeeper      slashing.Keeper
	mintKeeper          mint.Keeper
	distributionKeeper  distribution.Keeper
	paramsKeeper        params.Keeper
}

// NewHashgardApp returns a reference to an initialized HashgardApp.
func NewHashgardApp(logger log.Logger, db dbm.DB, traceStore io.Writer, baseAppOptions ...func(*bam.BaseApp)) *HashgardApp {
	// create and register app-level codec for TXs and accounts
	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)

	// create your application type
	var app = &HashgardApp{
		BaseApp:    		bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...),
		cdc:        		cdc,
		keyMain:    		sdk.NewKVStoreKey("main"),
		keyAccount: 		sdk.NewKVStoreKey("acc"),
		keyStake:			sdk.NewKVStoreKey("stake"),
		tkeyStake:			sdk.NewTransientStoreKey("transient_stake"),
		keyMint:			sdk.NewKVStoreKey("mint"),
		keyDistribution:	sdk.NewKVStoreKey("distribution"),
		tkeyDistribution:	sdk.NewTransientStoreKey("transient_distribution"),
		keySlashing:		sdk.NewKVStoreKey("slashing"),
		keyFeeCollection:	sdk.NewKVStoreKey("fee"),
		keyParams:			sdk.NewKVStoreKey("params"),
		tkeyParams:			sdk.NewTransientStoreKey("transient_params"),
	}

	// define the accountKeeper
	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		app.keyAccount, // target store
		auth.ProtoBaseAccount,
	)

	// add handlers
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper)

	app.feeCollectionKeeper = auth.NewFeeCollectionKeeper(
		app.cdc,
		app.keyFeeCollection,
	)

	app.paramsKeeper = params.NewKeeper(
		app.cdc,
		app.keyParams,
		app.tkeyParams,
	)

	stakeKeeper := stake.NewKeeper(
		app.cdc,
		app.keyStake,
		app.tkeyStake,
		app.bankKeeper,
		app.paramsKeeper.Subspace(stake.DefaultParamspace),
		app.RegisterCodespace(stake.DefaultCodespace),
	)

	app.mintKeeper = mint.NewKeeper(
		app.cdc,
		app.keyMint,
		app.paramsKeeper.Subspace(mint.DefaultParamspace),
		&stakeKeeper,
		app.feeCollectionKeeper,
	)

	app.distributionKeeper = distribution.NewKeeper(
		app.cdc,
		app.keyDistribution,
		app.paramsKeeper.Subspace(distribution.DefaultParamspace),
		app.bankKeeper,
		&stakeKeeper,
		app.feeCollectionKeeper,
		app.RegisterCodespace(stake.DefaultCodespace),
	)

	app.slashingKeeper = slashing.NewKeeper(
		app.cdc,
		app.keySlashing,
		&stakeKeeper,
		app.paramsKeeper.Subspace(slashing.DefaultParamspace),
		app.RegisterCodespace(slashing.DefaultCodespace),
	)

	// register the staking hooks
	// NOTE: stakeKeeper above are passed by reference,
	// so that it can be modified like below:
	app.stakeKeeper = *stakeKeeper.SetHooks(NewHooks(app.distributionKeeper.Hooks(), app.slashingKeeper.Hooks()))


	// register message routes
	app.Router().
		AddRoute("bank", bank.NewHandler(app.bankKeeper)).
		AddRoute("stake", stake.NewHandler(app.stakeKeeper)).
		AddRoute("distribution", distribution.NewHandler(app.distributionKeeper)).
		AddRoute("slashing", slashing.NewHandler(app.slashingKeeper))

	app.QueryRouter().
		AddRoute("stake", stake.NewQuerier(app.stakeKeeper, app.cdc))


	// initialize BaseApp
	app.MountStoresIAVL(
		app.keyMain,
		app.keyAccount,
		app.keyStake,
		app.keyMint,
		app.keyDistribution,
		app.keySlashing,
		app.keyFeeCollection,
		app.keyParams,
	)
	app.SetInitChainer(app.initChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.feeCollectionKeeper))
	app.MountStoresTransient(
		app.tkeyParams,
		app.tkeyStake,
		app.tkeyDistribution,
	)
	app.SetEndBlocker(app.EndBlocker)

	err := app.LoadLatestVersion(app.keyMain)
	if err != nil {
		cmn.Exit(err.Error())
	}

	return app
}

// MakeCodec creates a new codec codec and registers all the necessary types
// with the codec.
func MakeCodec() *codec.Codec {
	cdc := codec.New()

	codec.RegisterCrypto(cdc)
	sdk.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	stake.RegisterCodec(cdc)
	distribution.RegisterCodec(cdc)
	slashing.RegisterCodec(cdc)

	return cdc
}

// BeginBlocker reflects logic to run before any TXs application are processed
// by the application.
func (app *HashgardApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {

	tags := slashing.BeginBlocker(ctx, req, app.slashingKeeper)

	// distribute rewards from previous block
	distribution.BeginBlocker(ctx, req, app.distributionKeeper)

	// mint new tokens for this new block
	mint.BeginBlocker(ctx, app.mintKeeper)

	return abci.ResponseBeginBlock{
		Tags: tags.ToKVPairs(),
	}
}

// EndBlocker reflects logic to run after all TXs are processed by the
// application. Application updates every end block
func (app *HashgardApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {

	validatorUpdates := stake.EndBlocker(ctx, app.stakeKeeper)

	return abci.ResponseEndBlock{
		ValidatorUpdates: validatorUpdates,
	}
}

// initChainer implements the custom application logic that the BaseApp will
// invoke upon initialization. In this case, it will take the application's
// state provided by 'req' and attempt to deserialize said state. The state
// should contain all the genesis accounts. These accounts will be added to the
// application's account mapper.
// custom logic for hashgard
func (app *HashgardApp) initChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	stateJSON := req.AppStateBytes
	// TODO is this now the whole genesis file?

	var genesisState GenesisState
	err := app.cdc.UnmarshalJSON(stateJSON, &genesisState)
	if err != nil {
		// TODO: https://github.com/cosmos/cosmos-sdk/issues/468
		panic(err)
	}

	// sort by account number to maintain consistency
	sort.Slice(genesisState.Accounts, func(i, j int) bool {
		return genesisState.Accounts[i].AccountNumber < genesisState.Accounts[j].AccountNumber
	})

	// load the accounts
	for _, gacc := range genesisState.Accounts {
		acc := gacc.ToAccount()
		acc.AccountNumber = app.accountKeeper.GetNextAccountNumber(ctx)
		app.accountKeeper.SetAccount(ctx, acc)
	}

	// load the initial stake information
	validators, err := stake.InitGenesis(ctx, app.stakeKeeper, genesisState.StakeData)
	if err != nil {
		panic(err) // TODO find a way to do this w/o panics
	}

	// initialize module-specific stores
	auth.InitGenesis(ctx, app.feeCollectionKeeper, genesisState.AuthData)
	slashing.InitGenesis(ctx, app.slashingKeeper, genesisState.SlashingData, genesisState.StakeData)
	mint.InitGenesis(ctx, app.mintKeeper, genesisState.MintData)
	distribution.InitGenesis(ctx, app.distributionKeeper, genesisState.DistributionData)
	err = HashgardValidateGenesisState(genesisState)
	if err != nil {
		panic(err) // TODO find a way to do this w/o panics
	}

	if len(genesisState.GenTxs) > 0 {
		for _, genTx := range genesisState.GenTxs {
			var tx auth.StdTx
			err = app.cdc.UnmarshalJSON(genTx, &tx)
			if err != nil {
				panic(err)
			}
			bz := app.cdc.MustMarshalBinaryLengthPrefixed(tx)
			res := app.BaseApp.DeliverTx(bz)
			if !res.IsOK() {
				panic(res.Log)
			}
		}

		validators = app.stakeKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	}

	// sanity check
	if len(req.Validators) > 0 {
		if len(req.Validators) != len(validators) {
			panic(fmt.Errorf("len(RequestInitChain.Validators) != len(validators) (%d != %d)",
				len(req.Validators), len(validators)))
		}
		sort.Sort(abci.ValidatorUpdates(req.Validators))
		sort.Sort(abci.ValidatorUpdates(validators))
		for i, val := range validators {
			if !val.Equal(req.Validators[i]) {
				panic(fmt.Errorf("validators[%d] != req.Validators[%d] ", i, i))
			}
		}
	}


	return abci.ResponseInitChain{
		Validators: validators,
	}
}

// ExportAppStateAndValidators implements custom application logic that exposes
// various parts of the application's state and set of validators. An error is
// returned if any step getting the state or set of validators fails.
func (app *HashgardApp) ExportAppStateAndValidators() (appState json.RawMessage, validators []tmtypes.GenesisValidator, err error) {
	ctx := app.NewContext(true, abci.Header{})

	// iterate to get the accounts
	accounts := []GenesisAccount{}
	appendAccountsFn := func(acc auth.Account) (stop bool) {
		account := NewGenesisAccountI(acc)
		accounts = append(accounts, account)
		return false
	}

	app.accountKeeper.IterateAccounts(ctx, appendAccountsFn)

	genState := NewGenesisState(
		accounts,
		auth.ExportGenesis(ctx, app.feeCollectionKeeper),
		stake.ExportGenesis(ctx, app.stakeKeeper),
		mint.ExportGenesis(ctx, app.mintKeeper),
		distribution.ExportGenesis(ctx, app.distributionKeeper),
		slashing.ExportGenesis(ctx, app.slashingKeeper),
	)

	appState, err = codec.MarshalJSONIndent(app.cdc, genState)
	if err != nil {
		return nil, nil, err
	}

	validators = stake.WriteValidators(ctx, app.stakeKeeper)

	return appState, validators, err
}

//______________________________________________________________________________________________

// Combined Staking Hooks
type Hooks struct {
	dh distribution.Hooks
	sh slashing.Hooks
}

func NewHooks(dh distribution.Hooks, sh slashing.Hooks) Hooks {
	return Hooks{dh, sh}
}

var _ sdk.StakingHooks = Hooks{}

// nolint
func (h Hooks) OnValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) {
	h.dh.OnValidatorCreated(ctx, valAddr)
	h.sh.OnValidatorCreated(ctx, valAddr)
}
func (h Hooks) OnValidatorModified(ctx sdk.Context, valAddr sdk.ValAddress) {
	h.dh.OnValidatorModified(ctx, valAddr)
	h.sh.OnValidatorModified(ctx, valAddr)
}
func (h Hooks) OnValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	h.dh.OnValidatorRemoved(ctx, consAddr, valAddr)
	h.sh.OnValidatorRemoved(ctx, consAddr, valAddr)
}
func (h Hooks) OnValidatorBonded(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	h.dh.OnValidatorBonded(ctx, consAddr, valAddr)
	h.sh.OnValidatorBonded(ctx, consAddr, valAddr)
}
func (h Hooks) OnValidatorPowerDidChange(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	h.dh.OnValidatorPowerDidChange(ctx, consAddr, valAddr)
	h.sh.OnValidatorPowerDidChange(ctx, consAddr, valAddr)
}
func (h Hooks) OnValidatorBeginUnbonding(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	h.dh.OnValidatorBeginUnbonding(ctx, consAddr, valAddr)
	h.sh.OnValidatorBeginUnbonding(ctx, consAddr, valAddr)
}
func (h Hooks) OnDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.dh.OnDelegationCreated(ctx, delAddr, valAddr)
	h.sh.OnDelegationCreated(ctx, delAddr, valAddr)
}
func (h Hooks) OnDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.dh.OnDelegationSharesModified(ctx, delAddr, valAddr)
	h.sh.OnDelegationSharesModified(ctx, delAddr, valAddr)
}
func (h Hooks) OnDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.dh.OnDelegationRemoved(ctx, delAddr, valAddr)
	h.sh.OnDelegationRemoved(ctx, delAddr, valAddr)
}
