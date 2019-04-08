package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/hashgard/hashgard/x/issue"
	issueCli "github.com/hashgard/hashgard/x/issue/client/cli"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	govTxCmd := &cobra.Command{
		Use:   issue.ModuleName,
		Short: "Issue transactions subcommands",
	}

	govTxCmd.AddCommand(client.PostCommands(
		issueCli.GetCmdIssue(mc.cdc),
	)...)

	return govTxCmd
}
