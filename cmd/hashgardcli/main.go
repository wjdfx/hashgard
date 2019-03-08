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
	"github.com/cosmos/cosmos-sdk/x/auth"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	bankcmd "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	distributioncmd "github.com/cosmos/cosmos-sdk/x/distribution/client/cli"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govcmd "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingcmd "github.com/cosmos/cosmos-sdk/x/slashing/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakecmd "github.com/cosmos/cosmos-sdk/x/staking/client/cli"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/hashgard/hashgard/app"
	hashgardInit "github.com/hashgard/hashgard/init"
	"github.com/hashgard/hashgard/version"
)

// rootCmd is the entry point for this binary
var (
	rootCmd = &cobra.Command{
		Use:   "hashgardcli",
		Short: "Command line interface for interacting with hashgard",
	}
)

func main() {
	// get the codec
	cdc := app.MakeCodec()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(hashgardInit.Bech32PrefixAccAddr, hashgardInit.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(hashgardInit.Bech32PrefixValAddr, hashgardInit.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(hashgardInit.Bech32PrefixConsAddr, hashgardInit.Bech32PrefixConsPub)
	config.Seal()

	// disable sorting
	cobra.EnableCommandSorting = false

	// TODO: setup keybase, viper object, etc. to be passed into
	// the below functions and eliminate global vars, like we do
	// with the cdc

	// Add --chain-id to persistent flags and mark it required
	rootCmd.PersistentFlags().String(client.FlagChainID, "", "Chain ID of tendermint node")
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}


	// Add tendermint subcommands
	tendermintCmd := &cobra.Command{
		Use:     "tendermint",
		Short:   "Tendermint state querying subcommands",
	}
	tendermintCmd.AddCommand(
		rpc.BlockCommand(),
		rpc.ValidatorCommand(cdc),
		tx.SearchTxCmd(cdc),
		tx.QueryTxCmd(cdc),
	)

	// Add bank subcommands
	bankCmd := &cobra.Command{
		Use:	"bank",
		Short:	"Bank subcommands",
	}
	bankCmd.AddCommand(
		authcmd.GetAccountCmd(auth.StoreKey, cdc),
		client.LineBreak,
	)
	bankCmd.AddCommand(
		bankcmd.SendTxCmd(cdc),
		authcmd.GetSignCommand(cdc),
		authcmd.GetMultiSignCommand(cdc),
		authcmd.GetBroadcastCommand(cdc),
		authcmd.GetEncodeCommand(cdc),
	)

	// Add stake subcommands
	stakeCmd := &cobra.Command{
		Use:	"stake",
		Short:	"Stake and validation subcommands",
	}
	stakeCmd.AddCommand(
		client.GetCommands(
			stakecmd.GetCmdQueryDelegation(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryDelegations(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryUnbondingDelegation(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryUnbondingDelegations(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryRedelegation(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryRedelegations(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryValidator(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryValidators(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryValidatorDelegations(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryValidatorUnbondingDelegations(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryValidatorRedelegations(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryParams(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryPool(staking.StoreKey, cdc),
		)...)
	stakeCmd.AddCommand(client.LineBreak)
	stakeCmd.AddCommand(
		client.PostCommands(
			stakecmd.GetCmdCreateValidator(cdc),
			stakecmd.GetCmdEditValidator(cdc),
			stakecmd.GetCmdDelegate(cdc),
			stakecmd.GetCmdRedelegate(staking.StoreKey, cdc),
			stakecmd.GetCmdUnbond(staking.StoreKey, cdc),
		)...)

	// Add slashing subcommands
	slashingCmd := &cobra.Command{
		Use:	"slashing",
		Short:	"Slashing subcommands",
	}

	slashingCmd.AddCommand(
		client.GetCommands(
			slashingcmd.GetCmdQuerySigningInfo(slashing.StoreKey, cdc),
			slashingcmd.GetCmdQueryParams(cdc),
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
		client.GetCommands(
			distributioncmd.GetCmdQueryParams(distribution.StoreKey, cdc),
			distributioncmd.GetCmdQueryOutstandingRewards(distribution.StoreKey, cdc),
			distributioncmd.GetCmdQueryValidatorCommission(distribution.StoreKey, cdc),
			distributioncmd.GetCmdQueryValidatorSlashes(distribution.StoreKey, cdc),
			distributioncmd.GetCmdQueryDelegatorRewards(distribution.StoreKey, cdc),
		)...)
	distributionCmd.AddCommand(
		client.PostCommands(
			distributioncmd.GetCmdWithdrawRewards(cdc),
			distributioncmd.GetCmdSetWithdrawAddr(cdc),
			distributioncmd.GetCmdWithdrawAllRewards(cdc, distribution.StoreKey),
		)...)

	// Add gov subcommands
	govCmd := &cobra.Command{
		Use:	"gov",
		Short:	"Governance subcommands",
	}
	govCmd.AddCommand(
		client.GetCommands(
			govcmd.GetCmdQueryProposal(gov.StoreKey, cdc),
			govcmd.GetCmdQueryProposals(gov.StoreKey, cdc),
			govcmd.GetCmdQueryVote(gov.StoreKey, cdc),
			govcmd.GetCmdQueryVotes(gov.StoreKey, cdc),
			govcmd.GetCmdQueryParam(gov.StoreKey, cdc),
			govcmd.GetCmdQueryParams(gov.StoreKey, cdc),
			govcmd.GetCmdQueryProposer(gov.StoreKey, cdc),
			govcmd.GetCmdQueryDeposit(gov.StoreKey, cdc),
			govcmd.GetCmdQueryDeposits(gov.StoreKey, cdc),
			govcmd.GetCmdQueryTally(gov.StoreKey, cdc),
		)...)
	govCmd.AddCommand(client.LineBreak)
	govCmd.AddCommand(
		client.PostCommands(
			govcmd.GetCmdDeposit(gov.StoreKey, cdc),
			govcmd.GetCmdVote(gov.StoreKey, cdc),
			govcmd.GetCmdSubmitProposal(cdc),
		)...)

	rootCmd.AddCommand(
		client.ConfigCmd(app.DefaultCLIHome),
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
		version.VersionCmd,
	)

	// prepare and add flags
	executor := cli.PrepareMainCmd(rootCmd, "BC", app.DefaultCLIHome)
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