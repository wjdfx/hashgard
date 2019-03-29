package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"strconv"
	"strings"

	"github.com/hashgard/hashgard/x/issue/domain"
	"github.com/hashgard/hashgard/x/issue/msgs"
)

//TODO
// GetCmdIssue implements issue a coin transaction command.
func GetCmdIssueAdd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "issue a coin",
		Long: strings.TrimSpace(`
issue a coin. For example:
$ hashgardcli issue add --name=test --total-supply=8888888
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			coinIssueInfo, err := parseIssueFlags()
			if err != nil {
				return err
			}
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)
			from := cliCtx.GetFromAddress()
			coinIssueInfo.Issuer = from
			coinIssueInfo.Decimals = domain.DefaultDecimals
			msg := msgs.NewMsgIssue(coinIssueInfo)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	cmd.Flags().String(flagName, "", "name of coin")
	cmd.Flags().String(flagTotalSupply, "", "total supply of coin")
	return cmd
}

// GetCmdIssueMint implements mint a coinIssue transaction command.
func GetCmdIssueMint(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [issue-id] [amount] [option]",
		Args:  cobra.ExactArgs(2),
		Short: "mint a coin",
		Long: strings.TrimSpace(`
mint a coin. For example:
$ hashgardcli issue mint gardh1c7d59vebq 88888
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			num, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("amount %s not a valid int, please input a valid amount", args[1])
			}
			amount := sdk.NewInt(num)

			if !strings.HasPrefix(issueID, domain.IDPreStr) {
				return fmt.Errorf("issue-id %s not a valid issue, please input a valid issue-id", args[0])
			}
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)
			from := cliCtx.GetFromAddress()
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr,
				[]sdk.Msg{msgs.MsgIssueMint{issueID, from, amount}}, false)
		},
	}
	return cmd
}

// GetCmdIssueBurn implements burn a coinIssue transaction command.
func GetCmdIssueBurn(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [issue-id] [amount] [option]",
		Args:  cobra.ExactArgs(2),
		Short: "burn a coin",
		Long: strings.TrimSpace(`
mint a coin. For example:
$ hashgardcli issue burn gardh1c7d59vebq 88888
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			num, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("amount %s not a valid int, please input a valid amount", args[1])
			}
			amount := sdk.NewInt(num)
			if !strings.HasPrefix(issueID, domain.IDPreStr) {
				return fmt.Errorf("issue-id %s not a valid issue, please input a valid issue-id", args[0])
			}
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)
			from := cliCtx.GetFromAddress()
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr,
				[]sdk.Msg{msgs.MsgIssueBurn{issueID, from, amount}}, false)
		},
	}
	return cmd
}

// GetCmdIssueFinishMinting implements finishMinting a coinIssue transaction command.
func GetCmdIssueFinishMinting(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "finish-minting [issue-id] [option]",
		Args:  cobra.ExactArgs(1),
		Short: "finish-minting a coin",
		Long: strings.TrimSpace(`
mint a coin. For example:
$ hashgardcli issue finish-minting gardh1c7d59vebq
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)
			from := cliCtx.GetFromAddress()
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr,
				[]sdk.Msg{msgs.MsgIssueFinishMinting{issueID, from}}, false)
		},
	}
	return cmd
}
