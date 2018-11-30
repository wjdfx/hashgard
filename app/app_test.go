package app

import (
	"github.com/cosmos/cosmos-sdk/codec"
	abci "github.com/tendermint/tendermint/abci/types"
)

func setGenesis(baseApp *HashgardApp, accounts ...*AppAccount) (GenesisState, error) {
	genAccts := make([]GenesisAccount, len(accounts))
	for i, appAct := range accounts {
		genAccts[i] = *NewGenesisAccount(appAct)
	}

	genesisState := GenesisState{Accounts: genAccts}
	stateBytes, err := codec.MarshalJSONIndent(baseApp.cdc, genesisState)
	if err != nil {
		return GenesisState{}, err
	}

	// initialize and commit the chain
	baseApp.InitChain(abci.RequestInitChain{
		Validators: []abci.ValidatorUpdate{}, AppStateBytes: stateBytes,
	})
	baseApp.Commit()

	return genesisState, nil
}
