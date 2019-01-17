package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/stake"
	tmtypes "github.com/tendermint/tendermint/types"
)

var (
	// bonded tokens given to genesis validators/accounts
	FreeFermionVal  = int64(100)
	FreeFermionsAcc = sdk.NewInt(150)
)

const (
	defaultUnbondingTime time.Duration = 60 * 10 * time.Second
	GasDenom	= "apple"
	StakeDenom       = "tgard"
)

// State to Unmarshal
type GenesisState struct {
	Accounts     		[]GenesisAccount			`json:"accounts"`
	AuthData     		auth.GenesisState			`json:"auth"`
	StakeData    		stake.GenesisState			`json:"stake"`
	MintData     		mint.GenesisState			`json:"mint"`
	DistributionData    distribution.GenesisState	`json:"distribution"`
	SlashingData 		slashing.GenesisState		`json:"slashing"`
	GovData				gov.GenesisState			`json:"gov"`
	GenTxs       		[]json.RawMessage			`json:"gentxs"`
}

func NewGenesisState(
	accounts []GenesisAccount,
	authData auth.GenesisState,
	stakeData stake.GenesisState,
	mintData mint.GenesisState,
	distrData distribution.GenesisState,
	slashingData slashing.GenesisState,
	govData gov.GenesisState,
) GenesisState {

	return GenesisState{
		Accounts:     accounts,
		AuthData:     authData,
		StakeData:    stakeData,
		MintData:     mintData,
		DistributionData:    distrData,
		SlashingData: slashingData,
		GovData:      govData,
	}
}

// NewDefaultGenesisState generates the default state for hashgard.
func NewDefaultGenesisState() GenesisState {
	return GenesisState{
		Accounts:			nil,
		StakeData:    		createStakeGenesisState(),
		MintData:			createMintGenesisState(),
		DistributionData:	distribution.DefaultGenesisState(),
		SlashingData:		slashing.DefaultGenesisState(),
		GovData:			createGovGenesisState(),
		GenTxs:				nil,
	}
}

func createStakeGenesisState() stake.GenesisState {
	return stake.GenesisState{
		Pool: stake.Pool{
			LooseTokens:  sdk.ZeroDec(),
			BondedTokens: sdk.ZeroDec(),
		},
		Params: stake.Params{
			UnbondingTime: defaultUnbondingTime,
			MaxValidators: 100,
			BondDenom:     StakeDenom,
		},
	}
}

func createMintGenesisState() mint.GenesisState {
	return mint.GenesisState{
		Minter: mint.InitialMinter(sdk.NewDecWithPrec(13, 2)),
		Params: mint.Params{
			MintDenom:           StakeDenom,
			InflationRateChange: sdk.NewDecWithPrec(13, 2),
			InflationMax:        sdk.NewDecWithPrec(20, 2),
			InflationMin:        sdk.NewDecWithPrec(7, 2),
			GoalBonded:          sdk.NewDecWithPrec(67, 2),
			BlocksPerYear:       uint64(60 * 60 * 8766 / 5), // assuming 5 second block times
		},
	}
}

func createGovGenesisState() gov.GenesisState {
	return gov.GenesisState{
		StartingProposalID: 1,
		DepositParams: gov.DepositParams{
			MinDeposit:       sdk.Coins{sdk.NewInt64Coin(StakeDenom, 10)},
			MaxDepositPeriod: time.Duration(172800) * time.Second,
		},
		VotingParams: gov.VotingParams{
			VotingPeriod: time.Duration(172800) * time.Second,
		},
		TallyParams: gov.TallyParams{
			Quorum:            sdk.NewDecWithPrec(334, 3),
			Threshold:         sdk.NewDecWithPrec(5, 1),
			Veto:              sdk.NewDecWithPrec(334, 3),
			GovernancePenalty: sdk.NewDecWithPrec(1, 2),
		},
	}
}

// nolint
type GenesisAccount struct {
	Address       sdk.AccAddress `json:"address"`
	Coins         sdk.Coins      `json:"coins"`
	Sequence      uint64          `json:"sequence_number"`
	AccountNumber uint64          `json:"account_number"`
}

func NewGenesisAccount(acc *auth.BaseAccount) GenesisAccount {
	return GenesisAccount{
		Address:       acc.Address,
		Coins:         acc.Coins,
		AccountNumber: acc.AccountNumber,
		Sequence:      acc.Sequence,
	}
}

func NewGenesisAccountI(acc auth.Account) GenesisAccount {
	return GenesisAccount{
		Address:       acc.GetAddress(),
		Coins:         acc.GetCoins(),
		AccountNumber: acc.GetAccountNumber(),
		Sequence:      acc.GetSequence(),
	}
}

// convert GenesisAccount to auth.BaseAccount
func (ga *GenesisAccount) ToAccount() (acc *auth.BaseAccount) {
	return &auth.BaseAccount{
		Address:       ga.Address,
		Coins:         ga.Coins.Sort(),
		AccountNumber: ga.AccountNumber,
		Sequence:      ga.Sequence,
	}
}

func NewDefaultGenesisAccount(addr sdk.AccAddress) GenesisAccount {
	accAuth := auth.NewBaseAccountWithAddress(addr)
	coins := sdk.Coins{
		sdk.NewCoin(GasDenom, sdk.NewInt(1000)),
		sdk.NewCoin(StakeDenom, FreeFermionsAcc),
	}

	coins.Sort()

	accAuth.Coins = coins
	return NewGenesisAccount(&accAuth)
}


// HashgardAppGenState but with JSON
func HashgardAppGenStateJSON(cdc *codec.Codec, genDoc tmtypes.GenesisDoc, appGenTxs []json.RawMessage) (appState json.RawMessage, err error) {

	// create the final app state
	genesisState, err := HashgardAppGenState(cdc, genDoc, appGenTxs)
	if err != nil {
		return nil, err
	}
	return codec.MarshalJSONIndent(cdc, genesisState)
}

// Create the core parameters for genesis initialization for hashgard
// note that the pubkey input is this machines pubkey
func HashgardAppGenState(cdc *codec.Codec, genDoc tmtypes.GenesisDoc, appGenTxs []json.RawMessage) (genesisState GenesisState, err error) {


	if err = cdc.UnmarshalJSON(genDoc.AppState, &genesisState); err != nil {
		return genesisState, err
	}

	// if there are no gen txs to be processed, return the default empty state
	if len(appGenTxs) == 0 {
		return genesisState, errors.New("there must be at least one genesis tx")
	}

	stakeData := genesisState.StakeData
	for i, genTx := range appGenTxs {
		var tx auth.StdTx
		if err := cdc.UnmarshalJSON(genTx, &tx); err != nil {
			return genesisState, err
		}
		msgs := tx.GetMsgs()
		if len(msgs) != 1 {
			return genesisState, errors.New(
				"must provide genesis StdTx with exactly 1 CreateValidator message")
		}
		if _, ok := msgs[0].(stake.MsgCreateValidator); !ok {
			return genesisState, fmt.Errorf(
				"Genesis transaction %v does not contain a MsgCreateValidator", i)
		}
	}

	for _, acc := range genesisState.Accounts {
		// create the genesis account, give'm few steaks and a buncha token with there name
		for _, coin := range acc.Coins {
			if coin.Denom == StakeDenom {
				stakeData.Pool.LooseTokens = stakeData.Pool.LooseTokens.
					Add(sdk.NewDecFromInt(coin.Amount)) // increase the supply
			}
		}
	}

	genesisState.StakeData = stakeData
	genesisState.GenTxs = appGenTxs
	return genesisState, nil
}

// HashgardValidateGenesisState ensures that the genesis state obeys the expected invariants
// TODO: No validators are both bonded and jailed (#2088)
// TODO: Error if there is a duplicate validator (#1708)
// TODO: Ensure all state machine parameters are in genesis (#1704)
func HashgardValidateGenesisState(genesisState GenesisState) error {
	err := validateGenesisStateAccounts(genesisState.Accounts)
	if err != nil {
		return err
	}
	// skip stakeData validation as genesis is created from txs
	if len(genesisState.GenTxs) > 0 {
		return nil
	}

	err = stake.ValidateGenesis(genesisState.StakeData)
	if err != nil {
		return err
	}
	err = mint.ValidateGenesis(genesisState.MintData)
	if err != nil {
		return err
	}
	err = distribution.ValidateGenesis(genesisState.DistributionData)
	if err != nil {
		return err
	}
	err = gov.ValidateGenesis(genesisState.GovData)
	if err != nil {
		return err
	}
	err = slashing.ValidateGenesis(genesisState.SlashingData)
	if err != nil {
		return err
	}

	return nil
}

// Ensures that there are no duplicate accounts in the genesis state,
func validateGenesisStateAccounts(accs []GenesisAccount) (err error) {
	addrMap := make(map[string]bool, len(accs))
	for i := 0; i < len(accs); i++ {
		acc := accs[i]
		strAddr := string(acc.Address)
		if _, ok := addrMap[strAddr]; ok {
			return fmt.Errorf("Duplicate account in genesis state: Address %v", acc.Address)
		}
		addrMap[strAddr] = true
	}
	return
}

// CollectStdTxs processes and validates application's genesis StdTxs and returns
// the list of appGenTxs, and persistent peers required to generate genesis.json.
func CollectStdTxs(cdc *codec.Codec, moniker string, genTxsDir string, genDoc tmtypes.GenesisDoc) (
	appGenTxs []auth.StdTx, persistentPeers string, err error) {

	var fos []os.FileInfo
	fos, err = ioutil.ReadDir(genTxsDir)
	if err != nil {
		return appGenTxs, persistentPeers, err
	}

	// prepare a map of all accounts in genesis state to then validate
	// against the validators addresses
	var appState GenesisState
	if err := cdc.UnmarshalJSON(genDoc.AppState, &appState); err != nil {
		return appGenTxs, persistentPeers, err
	}

	addrMap := make(map[string]GenesisAccount, len(appState.Accounts))
	for i := 0; i < len(appState.Accounts); i++ {
		acc := appState.Accounts[i]
		addrMap[acc.Address.String()] = acc
	}

	// addresses and IPs (and port) validator server info
	var addressesIPs []string

	for _, fo := range fos {
		filename := filepath.Join(genTxsDir, fo.Name())
		if !fo.IsDir() && (filepath.Ext(filename) != ".json") {
			continue
		}

		// get the genStdTx
		var jsonRawTx []byte
		if jsonRawTx, err = ioutil.ReadFile(filename); err != nil {
			return appGenTxs, persistentPeers, err
		}
		var genStdTx auth.StdTx
		if err = cdc.UnmarshalJSON(jsonRawTx, &genStdTx); err != nil {
			return appGenTxs, persistentPeers, err
		}
		appGenTxs = append(appGenTxs, genStdTx)

		// the memo flag is used to store
		// the ip and node-id, for example this may be:
		// "528fd3df22b31f4969b05652bfe8f0fe921321d5@192.168.2.37:26656"
		nodeAddrIP := genStdTx.GetMemo()
		if len(nodeAddrIP) == 0 {
			return appGenTxs, persistentPeers, fmt.Errorf(
				"couldn't find node's address and IP in %s", fo.Name())
		}

		// genesis transactions must be single-message
		msgs := genStdTx.GetMsgs()
		if len(msgs) != 1 {

			return appGenTxs, persistentPeers, errors.New(
				"each genesis transaction must provide a single genesis message")
		}

		msg := msgs[0].(stake.MsgCreateValidator)
		// validate delegator and validator addresses and funds against the accounts in the state
		delAddr := msg.DelegatorAddr.String()
		valAddr := sdk.AccAddress(msg.ValidatorAddr).String()

		delAcc, delOk := addrMap[delAddr]
		_, valOk := addrMap[valAddr]

		accsNotInGenesis := []string{}
		if !delOk {
			accsNotInGenesis = append(accsNotInGenesis, delAddr)
		}
		if !valOk {
			accsNotInGenesis = append(accsNotInGenesis, valAddr)
		}
		if len(accsNotInGenesis) != 0 {
			return appGenTxs, persistentPeers, fmt.Errorf(
				"account(s) %v not in genesis.json: %+v", strings.Join(accsNotInGenesis, " "), addrMap)
		}

		if delAcc.Coins.AmountOf(msg.Delegation.Denom).LT(msg.Delegation.Amount) {
			return appGenTxs, persistentPeers, fmt.Errorf(
				"insufficient fund for delegation %v: %v < %v",
				delAcc.Address, delAcc.Coins.AmountOf(msg.Delegation.Denom), msg.Delegation.Amount,
			)
		}

		// exclude itself from persistent peers
		if msg.Description.Moniker != moniker {
			addressesIPs = append(addressesIPs, nodeAddrIP)
		}
	}

	sort.Strings(addressesIPs)
	persistentPeers = strings.Join(addressesIPs, ",")

	return appGenTxs, persistentPeers, nil
}