package app

import (
	"fmt"
	"io"
	"os"
	"sort"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/stake"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
)

const (
	appName = "HashgardApp"
	FlagReplay = "replay"
	// DefaultKeyPass contains key password for genesis transactions
	DefaultKeyPass = "12345678"
)

// default home directories for expected binaries
var (
	DefaultNodeHome = os.ExpandEnv("$HOME/.hashgard")
	DefaultCLIHome  = os.ExpandEnv("$HOME/.hashgardcli")
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
	keyGov           	*sdk.KVStoreKey
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
	govKeeper           gov.Keeper
	paramsKeeper        params.Keeper
}

// NewHashgardApp returns a reference to an initialized HashgardApp.
func NewHashgardApp(logger log.Logger, db dbm.DB, traceStore io.Writer, baseAppOptions ...func(*bam.BaseApp)) *HashgardApp {

	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)

	// create your application type
	var app = &HashgardApp{
		BaseApp:    		bApp,
		cdc:        		cdc,
		keyMain:    		sdk.NewKVStoreKey("main"),
		keyAccount: 		sdk.NewKVStoreKey("acc"),
		keyStake:			sdk.NewKVStoreKey("stake"),
		tkeyStake:			sdk.NewTransientStoreKey("transient_stake"),
		keyMint:			sdk.NewKVStoreKey("mint"),
		keyDistribution:	sdk.NewKVStoreKey("distribution"),
		tkeyDistribution:	sdk.NewTransientStoreKey("transient_distribution"),
		keySlashing:		sdk.NewKVStoreKey("slashing"),
		keyGov:				sdk.NewKVStoreKey("gov"),
		keyFeeCollection:	sdk.NewKVStoreKey("fee"),
		keyParams:			sdk.NewKVStoreKey("params"),
		tkeyParams:			sdk.NewTransientStoreKey("transient_params"),
	}

	// define the accountKeeper
	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		app.keyAccount,			// target store
		auth.ProtoBaseAccount,	// prototype
	)

	app.feeCollectionKeeper = auth.NewFeeCollectionKeeper(
		app.cdc,
		app.keyFeeCollection,
	)

	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper)

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
		stake.DefaultCodespace,
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
		distribution.DefaultCodespace,
	)

	app.slashingKeeper = slashing.NewKeeper(
		app.cdc,
		app.keySlashing,
		&stakeKeeper,
		app.paramsKeeper.Subspace(slashing.DefaultParamspace),
		slashing.DefaultCodespace,
	)

	app.govKeeper = gov.NewKeeper(
		app.cdc,
		app.keyGov,
		app.paramsKeeper,
		app.paramsKeeper.Subspace(gov.DefaultParamspace),
		app.bankKeeper,
		&stakeKeeper,
		gov.DefaultCodespace,
	)

	// register the staking hooks
	// NOTE: stakeKeeper above are passed by reference,
	// so that it can be modified like below:
	app.stakeKeeper = *stakeKeeper.SetHooks(
		NewStakingHooks(app.distributionKeeper.Hooks(), app.slashingKeeper.Hooks()),
	)


	// register message routes
	app.Router().
		AddRoute("bank", bank.NewHandler(app.bankKeeper)).
		AddRoute("stake", stake.NewHandler(app.stakeKeeper)).
		AddRoute("distribution", distribution.NewHandler(app.distributionKeeper)).
		AddRoute("slashing", slashing.NewHandler(app.slashingKeeper)).
		AddRoute("gov", gov.NewHandler(app.govKeeper))

	app.QueryRouter().
		AddRoute("stake", stake.NewQuerier(app.stakeKeeper, app.cdc)).
		AddRoute("gov", gov.NewQuerier(app.govKeeper))


	// initialize BaseApp
	app.MountStores(
		app.keyMain,
		app.keyAccount,
		app.keyStake,
		app.keyMint,
		app.keyDistribution,
		app.keySlashing,
		app.keyGov,
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

// MakeCodec creates a new codec and registers all the necessary types with the codec.
func MakeCodec() *codec.Codec {
	cdc := codec.New()

	codec.RegisterCrypto(cdc)
	sdk.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	stake.RegisterCodec(cdc)
	distribution.RegisterCodec(cdc)
	slashing.RegisterCodec(cdc)
	gov.RegisterCodec(cdc)

	return cdc
}

// BeginBlocker reflects logic to run before any TXs application are processed
// by the application.
func (app *HashgardApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {

	// mint new tokens for this new block
	mint.BeginBlocker(ctx, app.mintKeeper)

	// distribute rewards from previous block
	distribution.BeginBlocker(ctx, req, app.distributionKeeper)

	// slash anyone who double signed.
	// NOTE: This should happen after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool,
	// so as to keep the CanWithdrawInvariant invariant.
	// TODO: This should really happen at EndBlocker.
	tags := slashing.BeginBlocker(ctx, req, app.slashingKeeper)

	return abci.ResponseBeginBlock{
		Tags: tags.ToKVPairs(),
	}
}

// EndBlocker reflects logic to run after all TXs are processed by the application.
// Application updates every end block.
// nolint: unparam
func (app *HashgardApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {

	tags := gov.EndBlocker(ctx, app.govKeeper)
	validatorUpdates, endBlockerTags := stake.EndBlocker(ctx, app.stakeKeeper)
	tags = append(tags, endBlockerTags...)

	app.assertRuntimeInvariants()

	return abci.ResponseEndBlock{
		ValidatorUpdates: validatorUpdates,
		Tags:             tags,
	}
}

// initialize store from a genesis state
func (app *HashgardApp) initFromGenesisState(ctx sdk.Context, genesisState GenesisState) []abci.ValidatorUpdate {
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
	gov.InitGenesis(ctx, app.govKeeper, genesisState.GovData)
	mint.InitGenesis(ctx, app.mintKeeper, genesisState.MintData)
	distribution.InitGenesis(ctx, app.distributionKeeper, genesisState.DistributionData)

	// validate genesis state
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

	return validators
}

// custom logic for hashgard initialization
func (app *HashgardApp) initChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	stateJSON := req.AppStateBytes
	// TODO is this now the whole genesis file?

	var genesisState GenesisState
	err := app.cdc.UnmarshalJSON(stateJSON, &genesisState)
	if err != nil {
		panic(err) // TODO https://github.com/cosmos/cosmos-sdk/issues/468
		// return sdk.ErrGenesisParse("").TraceCause(err, "")
	}

	validators := app.initFromGenesisState(ctx, genesisState)

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

	// assert runtime invariants
	app.assertRuntimeInvariants()

	return abci.ResponseInitChain{
		Validators: validators,
	}
}

// load a particular height
func (app *HashgardApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keyMain)
}

//______________________________________________________________________________________________

var _ sdk.StakingHooks = StakingHooks{}

// StakingHooks contains combined distribution and slashing hooks needed for the
// staking module.
type StakingHooks struct {
	dh distribution.Hooks
	sh slashing.Hooks
}

func NewStakingHooks(dh distribution.Hooks, sh slashing.Hooks) StakingHooks {
	return StakingHooks{dh, sh}
}

// nolint
func (h StakingHooks) OnValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) {
	h.dh.OnValidatorCreated(ctx, valAddr)
	h.sh.OnValidatorCreated(ctx, valAddr)
}
func (h StakingHooks) OnValidatorModified(ctx sdk.Context, valAddr sdk.ValAddress) {
	h.dh.OnValidatorModified(ctx, valAddr)
	h.sh.OnValidatorModified(ctx, valAddr)
}
func (h StakingHooks) OnValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	h.dh.OnValidatorRemoved(ctx, consAddr, valAddr)
	h.sh.OnValidatorRemoved(ctx, consAddr, valAddr)
}
func (h StakingHooks) OnValidatorBonded(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	h.dh.OnValidatorBonded(ctx, consAddr, valAddr)
	h.sh.OnValidatorBonded(ctx, consAddr, valAddr)
}
func (h StakingHooks) OnValidatorPowerDidChange(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	h.dh.OnValidatorPowerDidChange(ctx, consAddr, valAddr)
	h.sh.OnValidatorPowerDidChange(ctx, consAddr, valAddr)
}
func (h StakingHooks) OnValidatorBeginUnbonding(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	h.dh.OnValidatorBeginUnbonding(ctx, consAddr, valAddr)
	h.sh.OnValidatorBeginUnbonding(ctx, consAddr, valAddr)
}
func (h StakingHooks) OnDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.dh.OnDelegationCreated(ctx, delAddr, valAddr)
	h.sh.OnDelegationCreated(ctx, delAddr, valAddr)
}
func (h StakingHooks) OnDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.dh.OnDelegationSharesModified(ctx, delAddr, valAddr)
	h.sh.OnDelegationSharesModified(ctx, delAddr, valAddr)
}
func (h StakingHooks) OnDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.dh.OnDelegationRemoved(ctx, delAddr, valAddr)
	h.sh.OnDelegationRemoved(ctx, delAddr, valAddr)
}