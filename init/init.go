package init

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	cfg "github.com/tendermint/tendermint/config"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/common"
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/hashgard/hashgard/app"
)

const (
	flagOverwrite  = "overwrite"
	flagClientHome = "home-client"
	flagMoniker    = "moniker"
)

type printInfo struct {
	Moniker    string          `json:"moniker"`
	ChainID    string          `json:"chain_id"`
	NodeID     string          `json:"node_id"`
	AppMessage json.RawMessage `json:"app_message"`
}

// nolint: errcheck
func displayInfo(cdc *codec.Codec, info printInfo) error {
	out, err := codec.MarshalJSONIndent(cdc, info)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "%s\n", string(out))
	return nil
}

// get cmd to initialize all files for tendermint and application
// nolint: errcheck
func InitCmd(ctx *server.Context, cdc *codec.Codec, appInit server.AppInit) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize genesis config, priv-validator file, p2p-node file, and application configuration files",
		Long: `Initialize validators's and node's configuration files.`,
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {

			config := ctx.Config
			config.SetRoot(viper.GetString(cli.HomeFlag))
			chainID := viper.GetString(client.FlagChainID)
			if chainID == "" {
				chainID = fmt.Sprintf("test-chain-%v", common.RandStr(6))
			}

			// initialize node validator files
			nodeID, pk, err := InitializeNodeValidatorFiles(config)
			if err != nil {
				return err
			}


			// set
			genTx, _, validator, err := server.SimpleAppGenTx(cdc, pk)
			if err != nil {
				return err
			}


			if viper.GetString(flagMoniker) != "" {
				config.Moniker = viper.GetString(flagMoniker)
			}

			var appStateJSON json.RawMessage
			genFile := config.GenesisFile()

			// initialize genesis.json file
			if appStateJSON, err = initializeEmptyGenesis(cdc, genFile, chainID, viper.GetBool(flagOverwrite)); err != nil {
				return err
			}

			// add genTx to appState
			var appState app.GenesisState
			cdc.UnmarshalJSON(appStateJSON, &appState)
			appState.GenTxs = append(appState.GenTxs, genTx)

			if appStateJSON, err = codec.MarshalJSONIndent(cdc, appState); err != nil {
				return err
			}

			if err = ExportGenesisFile(genFile, chainID, []tmtypes.GenesisValidator{validator}, appStateJSON); err != nil {
				return err
			}

			toPrint := printInfo{
				ChainID:    chainID,
				Moniker:    config.Moniker,
				NodeID:     nodeID,
				AppMessage: appStateJSON,
			}

			// generate config.toml
			cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)

			return displayInfo(cdc, toPrint)
		},
	}

	cmd.Flags().String(cli.HomeFlag, app.DefaultNodeHome, "node's home directory")
	cmd.Flags().String(flagClientHome, app.DefaultCLIHome, "client's home directory")
	cmd.Flags().String(client.FlagChainID, "",
		"genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().String(client.FlagName, "", "validator's moniker")
	cmd.MarkFlagRequired(client.FlagName)
	return cmd
}
