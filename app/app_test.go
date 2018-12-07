package app

import (
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/stake"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	abci "github.com/tendermint/tendermint/abci/types"
)

func setGenesis(happ *HashgardApp, accs ...*auth.BaseAccount) error {
	genaccs := make([]GenesisAccount, len(accs))
	for i, acc := range accs {
		genaccs[i] = NewGenesisAccount(acc)
	}

	genesisState := GenesisState{
		Accounts:     genaccs,
		StakeData:    stake.DefaultGenesisState(),
		DistributionData:    distribution.DefaultGenesisState(),
		SlashingData: slashing.DefaultGenesisState(),
	}

	stateBytes, err := codec.MarshalJSONIndent(happ.cdc, genesisState)
	if err != nil {
		return err
	}

	// Initialize the chain
	vals := []abci.ValidatorUpdate{}
	happ.InitChain(abci.RequestInitChain{Validators: vals, AppStateBytes: stateBytes})
	happ.Commit()

	return nil
}

func TestHashgardExport(t *testing.T) {
	db := db.NewMemDB()
	gapp := NewHashgardApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil)
	setGenesis(gapp)

	// Making a new app object with the db, so that initchain hasn't been called
	newGapp := NewHashgardApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil)
	_, _, err := newGapp.ExportAppStateAndValidators()
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}
