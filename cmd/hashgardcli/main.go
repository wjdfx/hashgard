package main

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	_ "github.com/cosmos/cosmos-sdk/client/lcd/statik"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	bankcmd "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	govcmd "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	distrcmd "github.com/cosmos/cosmos-sdk/x/distribution/client/cli"
	slashingcmd "github.com/cosmos/cosmos-sdk/x/slashing/client/cli"
	stakecmd "github.com/cosmos/cosmos-sdk/x/stake/client/cli"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/hashgard/hashgard/app"
	hashgardInit "github.com/hashgard/hashgard/init"
	"github.com/hashgard/hashgard/version"
)

const (
	storeAcc        = "acc"
	storeGov        = "gov"
	storeSlashing   = "slashing"
	storeStake      = "stake"
	queryRouteStake = "stake"
)

// rootCmd is the entry point for this binary
var (
	rootCmd = &cobra.Command{
		Use:   "hashgardcli",
		Short: "Hashgard light-client",
	}
)

func main() {
	// disable sorting
	cobra.EnableCommandSorting = false

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(hashgardInit.Bech32PrefixAccAddr, hashgardInit.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(hashgardInit.Bech32PrefixValAddr, hashgardInit.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(hashgardInit.Bech32PrefixConsAddr, hashgardInit.Bech32PrefixConsPub)
	config.Seal()

	// get the codec
	cdc := app.MakeCodec()

	// TODO: setup keybase, viper object, etc. to be passed into
	// the below functions and eliminate global vars, like we do
	// with the cdc
	rootCmd.AddCommand(client.ConfigCmd())

	// add standard rpc commands
	rpc.AddCommands(rootCmd)

	//Add query commands
	queryCmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Querying subcommands",
	}
	queryCmd.AddCommand(
		rpc.BlockCommand(),
		rpc.ValidatorCommand(),
	)

	tx.AddCommands(queryCmd, cdc)

	queryCmd.AddCommand(client.LineBreak)
	queryCmd.AddCommand(client.GetCommands(
		authcmd.GetAccountCmd(storeAcc, cdc, authcmd.GetAccountDecoder(cdc)),
		stakecmd.GetCmdQueryDelegation(storeStake, cdc),
		stakecmd.GetCmdQueryDelegations(storeStake, cdc),
		stakecmd.GetCmdQueryUnbondingDelegation(storeStake, cdc),
		stakecmd.GetCmdQueryUnbondingDelegations(storeStake, cdc),
		stakecmd.GetCmdQueryRedelegation(storeStake, cdc),
		stakecmd.GetCmdQueryRedelegations(storeStake, cdc),
		stakecmd.GetCmdQueryValidator(storeStake, cdc),
		stakecmd.GetCmdQueryValidators(storeStake, cdc),
		stakecmd.GetCmdQueryValidatorUnbondingDelegations(queryRouteStake, cdc),
		stakecmd.GetCmdQueryValidatorRedelegations(queryRouteStake, cdc),
		stakecmd.GetCmdQueryParams(storeStake, cdc),
		stakecmd.GetCmdQueryPool(storeStake, cdc),
		govcmd.GetCmdQueryProposal(storeGov, cdc),
		govcmd.GetCmdQueryProposals(storeGov, cdc),
		govcmd.GetCmdQueryVote(storeGov, cdc),
		govcmd.GetCmdQueryVotes(storeGov, cdc),
		govcmd.GetCmdQueryDeposit(storeGov, cdc),
		govcmd.GetCmdQueryDeposits(storeGov, cdc),
		slashingcmd.GetCmdQuerySigningInfo(storeSlashing, cdc),
	)...)

	//Add query commands
	txCmd := &cobra.Command{
		Use:   "tx",
		Short: "Transactions subcommands",
	}

	//Add auth and bank commands
	txCmd.AddCommand(
		client.PostCommands(
			bankcmd.GetBroadcastCommand(cdc),
			authcmd.GetSignCommand(cdc, authcmd.GetAccountDecoder(cdc)),
		)...)
	txCmd.AddCommand(client.LineBreak)

	txCmd.AddCommand(
		client.PostCommands(
			stakecmd.GetCmdCreateValidator(cdc),
			stakecmd.GetCmdEditValidator(cdc),
			stakecmd.GetCmdDelegate(cdc),
			stakecmd.GetCmdRedelegate(storeStake, cdc),
			stakecmd.GetCmdUnbond(storeStake, cdc),
			distrcmd.GetCmdWithdrawRewards(cdc),
			distrcmd.GetCmdSetWithdrawAddr(cdc),
			govcmd.GetCmdDeposit(cdc),
			bankcmd.SendTxCmd(cdc),
			govcmd.GetCmdSubmitProposal(cdc),
			slashingcmd.GetCmdUnjail(cdc),
			govcmd.GetCmdVote(cdc),
		)...)

	rootCmd.AddCommand(
		queryCmd,
		txCmd,
		lcd.ServeCommand(cdc),
		client.LineBreak,
	)

	// add proxy, version and key info
	rootCmd.AddCommand(
		keys.Commands(),
		client.LineBreak,
		version.ServeVersionCommand(cdc),
	)

	// prepare and add flags
	executor := cli.PrepareMainCmd(rootCmd, "BC", app.DefaultCLIHome)
	err := executor.Execute()
	if err != nil {
		// Note: Handle with #870
		panic(err)
	}
}
