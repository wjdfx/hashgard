package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"

	boxCli "github.com/hashgard/hashgard/x/box/client/cli"
	"github.com/hashgard/hashgard/x/box/types"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	cdc *amino.Codec
}

//New ModuleClient Instance
func NewModuleClient(cdc *amino.Codec) ModuleClient {
	return ModuleClient{cdc}
}

// GetBoxCmd returns the box commands for this module
func (mc ModuleClient) GetBoxCmd() *cobra.Command {
	boxCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "Box token subcommands",
	}
	boxCmd.AddCommand(
		client.GetCommands(
			boxCli.GetCmdQueryBoxs(mc.cdc),
			boxCli.GetCmdQueryBox(mc.cdc),
			boxCli.GetCmdSearchBoxs(mc.cdc),
			boxCli.GetCmdQueryDepositBoxDeposit(mc.cdc),
		)...)
	boxCmd.AddCommand(client.LineBreak)

	cmdDepositBox := boxCli.GetCmdDepositBoxCreate(mc.cdc)
	boxCli.MarkCmdDepositBoxCreateFlagRequired(cmdDepositBox)

	txCmd := client.PostCommands(
		boxCli.GetCmdLockBoxCreate(mc.cdc),
		cmdDepositBox,
		boxCli.GetCmdFutureBoxCreate(mc.cdc),
		boxCli.GetCmdDepositBoxInterestInjection(mc.cdc),
		boxCli.GetCmdDepositBoxInterestFetch(mc.cdc),
		boxCli.GetCmdDepositToBox(mc.cdc),
		boxCli.GetCmdFetchDepositFromBox(mc.cdc),
		boxCli.GetCmdBoxDescription(mc.cdc),
		boxCli.GetCmdBoxDisableFeature(mc.cdc),
	)

	for _, cmd := range txCmd {
		_ = cmd.MarkFlagRequired(client.FlagFrom)
		boxCmd.AddCommand(cmd)
	}

	return boxCmd
}
