package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"

	issueCli "github.com/hashgard/hashgard/x/issue/client/cli"
	"github.com/hashgard/hashgard/x/issue/domain"
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
	issueCmd := &cobra.Command{
		Use:   domain.ModuleName,
		Short: "Issue coin subcommands",
	}
	issueCmd.AddCommand(
		client.PostCommands(
			issueCli.GetCmdIssueAdd(mc.cdc),
			issueCli.GetCmdIssueMint(mc.cdc),
			issueCli.GetCmdIssueBurn(mc.cdc),
			issueCli.GetCmdIssueFinishMinting(mc.cdc),
		)...)
	issueCmd.AddCommand(client.LineBreak)
	issueCmd.AddCommand(
		client.GetCommands(
			issueCli.GetCmdQueryIssue(mc.storeKey, mc.cdc),
		)...)

	return issueCmd
}
