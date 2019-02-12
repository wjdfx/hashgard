package main

import (
	"encoding/json"
	"io"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"

	"github.com/hashgard/hashgard/app"
	"github.com/hashgard/hashgard/version"
	hashgardInit "github.com/hashgard/hashgard/init"
)

func main() {

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(hashgardInit.Bech32PrefixAccAddr, hashgardInit.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(hashgardInit.Bech32PrefixValAddr, hashgardInit.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(hashgardInit.Bech32PrefixConsAddr, hashgardInit.Bech32PrefixConsPub)
	config.Seal()

	cdc := app.MakeCodec()
	ctx := server.NewDefaultContext()
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:               "hashgard",
		Short:             "Hashgard Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}
	tendermintCmd := &cobra.Command{
		Use:   "tendermint",
		Short: "Tendermint subcommands",
	}

	tendermintCmd.AddCommand(
		server.ShowNodeIDCmd(ctx),
		server.ShowValidatorCmd(ctx),
		server.ShowAddressCmd(ctx),
	)

	startCmd := server.StartCmd(ctx, newApp)
	startCmd.Flags().Bool(app.FlagReplay, false, "Replay the last block")

	rootCmd.AddCommand(
		startCmd,
		hashgardInit.InitCmd(ctx, cdc),
		hashgardInit.CollectGenTxsCmd(ctx, cdc),
		hashgardInit.TestnetFilesCmd(ctx, cdc),
		hashgardInit.GenTxCmd(ctx, cdc),
		hashgardInit.AddGenesisAccountCmd(ctx, cdc),
		server.UnsafeResetAllCmd(ctx),
		client.LineBreak,
		tendermintCmd,
		client.LineBreak,
		server.ExportCmd(ctx, cdc, exportAppStateAndTMValidators),
		client.LineBreak,
		version.VersionCmd,
	)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "BC", app.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		// handle with #870
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewHashgardApp(
		logger,
		db,
		traceStore,
		true,
		baseapp.SetPruning(store.NewPruningOptions(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
	)
}

func exportAppStateAndTMValidators(logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool) (json.RawMessage, []tmtypes.GenesisValidator, error) {

	if height != -1 {
		hApp := app.NewHashgardApp(logger, db, traceStore, false)
		err := hApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return hApp.ExportAppStateAndValidators(forZeroHeight)
	}

	hApp := app.NewHashgardApp(logger, db, traceStore, false)
	return hApp.ExportAppStateAndValidators(forZeroHeight)
}
