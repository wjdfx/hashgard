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

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	tmtypes "github.com/tendermint/tendermint/types"
)


var _ auth.Account = (*AppAccount)(nil)

// AppAccount is a custom extension for this application. It is an example of
// extending auth.BaseAccount with custom fields. It is compatible with the
// stock auth.AccountKeeper, since auth.AccountKeeper uses the flexible go-amino
// library.
type AppAccount struct {
	auth.BaseAccount

	Name string `json:"name"`
}

// nolint
func (acc AppAccount) GetName() string      { return acc.Name }
func (acc *AppAccount) SetName(name string) { acc.Name = name }

// NewAppAccount returns a reference to a new AppAccount given a name and an
// auth.BaseAccount.
func NewAppAccount(name string, baseAcct auth.BaseAccount) *AppAccount {
	return &AppAccount{BaseAccount: baseAcct, Name: name}
}

// GetAccountDecoder returns the AccountDecoder function for the custom
// AppAccount.
func GetAccountDecoder(cdc *codec.Codec) auth.AccountDecoder {
	return func(accBytes []byte) (auth.Account, error) {
		if len(accBytes) == 0 {
			return nil, sdk.ErrTxDecode("accBytes are empty")
		}

		acct := new(AppAccount)
		err := cdc.UnmarshalBinaryBare(accBytes, &acct)
		if err != nil {
			panic(err)
		}

		return acct, err
	}
}

// GenesisState reflects the genesis state of the application.
type GenesisState struct {
	Accounts []GenesisAccount `json:"accounts"`
	GenTxs       []json.RawMessage     `json:"gentxs"`
}

// GenesisAccount reflects a genesis account the application expects in it's
// genesis state.
type GenesisAccount struct {
	Name    string         `json:"name"`
	Address sdk.AccAddress `json:"address"`
	Coins   sdk.Coins      `json:"coins"`
}

// NewDefaultGenesisState generates the default state for hashgard.
func NewDefaultGenesisState() GenesisState {
	return GenesisState{
		Accounts:     nil,
		GenTxs:       nil,
	}
}

// NewGenesisAccount returns a reference to a new GenesisAccount given an
// AppAccount.
func NewGenesisAccount(aa *AppAccount) *GenesisAccount {
	return &GenesisAccount{
		Name:    aa.Name,
		Address: aa.Address,
		Coins:   aa.Coins.Sort(),
	}
}

// ToAppAccount converts a GenesisAccount to an AppAccount.
func (ga *GenesisAccount) ToAppAccount() (acc *AppAccount, err error) {
	return &AppAccount{
		Name: ga.Name,
		BaseAccount: auth.BaseAccount{
			Address: ga.Address,
			Coins:   ga.Coins.Sort(),
		},
	}, nil
}

// get app init parameters for server init command
func HashgardAppInit() server.AppInit {
	return server.AppInit{
		AppGenState: HashgardAppGenStateJSON,
	}
}

// HashgardAppGenState but with JSON
func HashgardAppGenStateJSON(cdc *codec.Codec, genDoc tmtypes.GenesisDoc, appGenTxs []json.RawMessage) (appState json.RawMessage, err error) {

	// create the final app state
	genesisState, err := HashgardAppGenState(cdc, genDoc, appGenTxs)
	if err != nil {
		return nil, err
	}
	appState, err = codec.MarshalJSONIndent(cdc, genesisState)
	return
}

// Create the core parameters for genesis initialization for hashgard
// note that the pubkey input is this machines pubkey
func HashgardAppGenState(cdc *codec.Codec, genDoc tmtypes.GenesisDoc, appGenTxs []json.RawMessage) (genesisState GenesisState, err error) {


	//if err = cdc.UnmarshalJSON(genDoc.AppState, &genesisState); err != nil {
	//	return genesisState, err
	//}

	// if there are no gen txs to be processed, return the default empty state
	if len(appGenTxs) == 0 {
		return genesisState, errors.New("there must be at least one genesis tx")
	}

	var accounts = make([]GenesisAccount, 0)

	for _, genTx := range appGenTxs {
		var tx server.SimpleGenTx
		err = cdc.UnmarshalJSON(genTx, &tx)
		if err != nil {
			return
		}

		var account = GenesisAccount{Address:tx.Addr, Coins:[]sdk.Coin{{Denom:"gard", Amount: sdk.NewInt(100000000)}}}
		accounts = append(accounts, account)
	}

	genesisState.Accounts = accounts

	genesisState.GenTxs = appGenTxs

	return genesisState, nil
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
		strAddr := string(acc.Address)
		addrMap[strAddr] = acc
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

		//// validate the validator address and funds against the accounts in the state
		//msg := msgs[0].(stake.MsgCreateValidator)
		//addr := string(sdk.AccAddress(msg.ValidatorAddr))
		//acc, ok := addrMap[addr]
		//if !ok {
		//	return appGenTxs, persistentPeers, fmt.Errorf(
		//		"account %v not in genesis.json: %+v", addr, addrMap)
		//}
		//if acc.Coins.AmountOf(msg.Delegation.Denom).LT(msg.Delegation.Amount) {
		//	err = fmt.Errorf("insufficient fund for the delegation: %s < %s",
		//		acc.Coins.AmountOf(msg.Delegation.Denom), msg.Delegation.Amount)
		//}
		//
		//// exclude itself from persistent peers
		//if msg.Description.Moniker != moniker {
		//	addressesIPs = append(addressesIPs, nodeAddrIP)
		//}
	}

	sort.Strings(addressesIPs)
	persistentPeers = strings.Join(addressesIPs, ",")

	return appGenTxs, persistentPeers, nil
}

