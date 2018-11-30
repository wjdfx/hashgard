package main

import (
	"encoding/json"
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
	"io"

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

	// tendermint subcommands
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
		hashgardInit.InitCmd(ctx, cdc, app.HashgardAppInit()),
		startCmd,
		tendermintCmd,

		server.UnsafeResetAllCmd(ctx),
		client.LineBreak,
		tendermintCmd,
		server.ExportCmd(ctx, cdc, exportAppStateAndTMValidators),
		client.LineBreak,
	)

	rootCmd.AddCommand(
		version.ServeVersionCommand(cdc),
	)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "BC", app.DefaultNodeHome)
	executor.Execute()

}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewHashgardApp(logger, db, traceStore, baseapp.SetPruning(viper.GetString("pruning")))
}

func exportAppStateAndTMValidators(logger log.Logger, db dbm.DB, traceStore io.Writer) (json.RawMessage, []tmtypes.GenesisValidator, error) {
	gApp := app.NewHashgardApp(logger, db, traceStore)
	return gApp.ExportAppStateAndValidators()
}
