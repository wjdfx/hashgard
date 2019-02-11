package main

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	_ "github.com/cosmos/cosmos-sdk/client/lcd/statik"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	bankcmd "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	distributioncmd "github.com/cosmos/cosmos-sdk/x/distribution/client/cli"
	govcmd "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	slashingcmd "github.com/cosmos/cosmos-sdk/x/slashing/client/cli"
	stakecmd "github.com/cosmos/cosmos-sdk/x/staking/client/cli"
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
	storeDistribution	= "distribution"
)

// rootCmd is the entry point for this binary
var (
	rootCmd = &cobra.Command{
		Use:   "hashgardcli",
		Short: "Hashgard light-client",
	}
)

func main() {

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(hashgardInit.Bech32PrefixAccAddr, hashgardInit.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(hashgardInit.Bech32PrefixValAddr, hashgardInit.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(hashgardInit.Bech32PrefixConsAddr, hashgardInit.Bech32PrefixConsPub)
	config.Seal()

	// get the codec
	cdc := app.MakeCodec()

	// disable sorting
	cobra.EnableCommandSorting = false

	// TODO: setup keybase, viper object, etc. to be passed into
	// the below functions and eliminate global vars, like we do
	// with the cdc


	// Add tendermint subcommands
	tendermintCmd := &cobra.Command{
		Use:     "tendermint",
		Short:   "Tendermint state querying subcommands",
	}
	tendermintCmd.AddCommand(
		rpc.BlockCommand(),
		rpc.ValidatorCommand(),
		tx.SearchTxCmd(cdc),
		tx.QueryTxCmd(cdc),
	)

	// Add bank subcommands
	bankCmd := &cobra.Command{
		Use:	"bank",
		Short:	"Bank subcommands",
	}
	bankCmd.AddCommand(
		authcmd.GetAccountCmd(storeAcc, cdc),
		client.LineBreak,
	)
	bankCmd.AddCommand(
		bankcmd.SendTxCmd(cdc),
		authcmd.GetSignCommand(cdc),
		authcmd.GetMultiSignCommand(cdc),
		bankcmd.GetBroadcastCommand(cdc),
	)

	// Add stake subcommands
	stakeCmd := &cobra.Command{
		Use:	"stake",
		Short:	"Stake and validation subcommands",
	}
	stakeCmd.AddCommand(
		client.GetCommands(
			stakecmd.GetCmdQueryDelegation(storeStake, cdc),
			stakecmd.GetCmdQueryDelegations(storeStake, cdc),
			stakecmd.GetCmdQueryUnbondingDelegation(storeStake, cdc),
			stakecmd.GetCmdQueryUnbondingDelegations(storeStake, cdc),
			stakecmd.GetCmdQueryRedelegation(storeStake, cdc),
			stakecmd.GetCmdQueryRedelegations(storeStake, cdc),
			stakecmd.GetCmdQueryValidator(storeStake, cdc),
			stakecmd.GetCmdQueryValidators(storeStake, cdc),
			stakecmd.GetCmdQueryValidatorDelegations(storeStake, cdc),
			stakecmd.GetCmdQueryValidatorUnbondingDelegations(storeStake, cdc),
			stakecmd.GetCmdQueryValidatorRedelegations(storeStake, cdc),
			stakecmd.GetCmdQueryParams(storeStake, cdc),
			stakecmd.GetCmdQueryPool(storeStake, cdc),
		)...)
	stakeCmd.AddCommand(client.LineBreak)
	stakeCmd.AddCommand(
		client.PostCommands(
			stakecmd.GetCmdCreateValidator(cdc),
			stakecmd.GetCmdEditValidator(cdc),
			stakecmd.GetCmdDelegate(cdc),
			stakecmd.GetCmdRedelegate(storeStake, cdc),
			stakecmd.GetCmdUnbond(storeStake, cdc),
		)...)

	// Add slashing subcommands
	slashingCmd := &cobra.Command{
		Use:	"slashing",
		Short:	"Slashing subcommands",
	}
	slashingCmd.AddCommand(
		client.GetCommands(
			slashingcmd.GetCmdQuerySigningInfo(storeSlashing, cdc),
		)...)
	slashingCmd.AddCommand(client.LineBreak)
	slashingCmd.AddCommand(
		client.PostCommands(
			slashingcmd.GetCmdUnjail(cdc),
		)...)

	// Add distribution subcommands
	distributionCmd := &cobra.Command{
		Use:	"distribution",
		Short:	"Distribution subcommands",
	}
	distributionCmd.AddCommand(
		client.PostCommands(
			distributioncmd.GetCmdWithdrawRewards(cdc),
			distributioncmd.GetCmdSetWithdrawAddr(cdc),
		)...)

	// Add gov subcommands
	govCmd := &cobra.Command{
		Use:	"gov",
		Short:	"Governance subcommands",
	}
	govCmd.AddCommand(
		client.GetCommands(
			govcmd.GetCmdQueryProposal(storeGov, cdc),
			govcmd.GetCmdQueryProposals(storeGov, cdc),
			govcmd.GetCmdQueryVote(storeGov, cdc),
			govcmd.GetCmdQueryVotes(storeGov, cdc),
			govcmd.GetCmdQueryParams(storeGov, cdc),
			govcmd.GetCmdQueryDeposit(storeGov, cdc),
			govcmd.GetCmdQueryDeposits(storeGov, cdc),
			govcmd.GetCmdQueryTally(storeGov, cdc),
		)...)
	govCmd.AddCommand(client.LineBreak)
	govCmd.AddCommand(
		client.PostCommands(
			govcmd.GetCmdDeposit(storeGov, cdc),
			govcmd.GetCmdVote(storeGov, cdc),
			govcmd.GetCmdSubmitProposal(cdc),
		)...)

	rootCmd.AddCommand(
		client.ConfigCmd(),
		rpc.StatusCommand(),
		client.LineBreak,
		keys.Commands(),
		tendermintCmd,
		client.LineBreak,
		bankCmd,
		stakeCmd,
		slashingCmd,
		distributionCmd,
		govCmd,
		client.LineBreak,
		client.LineBreak,
		version.VersionCmd,
	)

	// prepare and add flags
	executor := cli.PrepareMainCmd(rootCmd, "HG", app.DefaultCLIHome)
	err := initConfig(rootCmd)
	if err != nil {
		panic(err)
	}

	err = executor.Execute()
	if err != nil {
		fmt.Printf("Failed executing CLI command: %s, exiting...\n", err)
		os.Exit(1)
	}
}

func initConfig(cmd *cobra.Command) error {
	home, err := cmd.PersistentFlags().GetString(cli.HomeFlag)
	if err != nil {
		return err
	}

	cfgFile := path.Join(home, "config", "config.toml")
	if _, err := os.Stat(cfgFile); err == nil {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}

	if err := viper.BindPFlag(cli.EncodingFlag, cmd.PersistentFlags().Lookup(cli.EncodingFlag)); err != nil {
		return err
	}
	return viper.BindPFlag(cli.OutputFlag, cmd.PersistentFlags().Lookup(cli.OutputFlag))
}