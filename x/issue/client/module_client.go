package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"

	issueCli "github.com/hashgard/hashgard/x/issue/client/cli"
	"github.com/hashgard/hashgard/x/issue/types"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	cdc *amino.Codec
}

//New ModuleClient Instance
func NewModuleClient(cdc *amino.Codec) ModuleClient {
	return ModuleClient{cdc}
}

// GetIssueCmd returns the issue commands for this module
func (mc ModuleClient) GetIssueCmd() *cobra.Command {
	issueCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "Issue coin subcommands",
	}
	issueCmd.AddCommand(
		client.GetCommands(
			issueCli.GetCmdQueryIssue(mc.cdc),
			issueCli.GetCmdQueryIssues(mc.cdc),
		)...)
	issueCmd.AddCommand(client.LineBreak)

	txCmd := client.PostCommands(
		issueCli.GetCmdIssueCreate(mc.cdc),
		issueCli.GetCmdIssueDescription(mc.cdc),
		issueCli.GetCmdIssueMint(mc.cdc),
		issueCli.GetCmdIssueBurn(mc.cdc),
		issueCli.GetCmdIssueFinishMinting(mc.cdc),
	)
	for _, cmd := range txCmd {
		_ = cmd.MarkFlagRequired(client.FlagFrom)
		issueCmd.AddCommand(cmd)
	}

	return issueCmd
}
