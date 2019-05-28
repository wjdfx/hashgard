package clientpackage

import (
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"

	"github.com/hashgard/hashgard/x/mint"
	"github.com/hashgard/hashgard/x/mint/client/cli"
)

// ModuleClient exports all CLI client functionality from the minting module.
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for the minting module.
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	mintingQueryCmd := &cobra.Command{
		Use:   mint.ModuleName,
		Short: "Querying commands for the minting module",
	}

	mintingQueryCmd.AddCommand(
		client.GetCommands(
			cli.GetCmdQueryParams(mc.cdc),
			cli.GetCmdQueryMinter(mc.cdc),
		)...,
	)

	return mintingQueryCmd
}

// GetTxCmd returns the transaction commands for the minting module.
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	mintTxCmd := &cobra.Command{
		Use:   mint.ModuleName,
		Short: "Minting transaction subcommands",
	}

	return mintTxCmd
}
