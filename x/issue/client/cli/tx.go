package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
	"strings"

	issueutils "github.com/hashgard/hashgard/x/issue/utils"

	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/types"
)

// GetCmdIssue implements issue a coin transaction command.
func GetCmdIssueAdd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name] [symbol] [total-supply]",
		Args:  cobra.ExactArgs(3),
		Short: "issue a coin",
		Long: strings.TrimSpace(`
issue a coin. For example:
$ hashgardcli issue add mytestcoin test 1000000000000000
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			totalSupply, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("Total supply %s not a valid int, please input a valid total supply", args[2])
			}
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)
			from := cliCtx.GetFromAddress()

			msg := msgs.NewMsgIssue(&types.CoinIssueInfo{
				Issuer:          from,
				Name:            args[0],
				Symbol:          args[1],
				MintingFinished: false,
				TotalSupply:     totalSupply,
				Decimals:        types.DefaultDecimals,
			})
			if viper.IsSet(flagMintingFinished) {
				msg.MintingFinished = viper.GetBool(flagMintingFinished)
			}
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	cmd.Flags().Bool(flagMintingFinished, false, "can minting of coin")
	return cmd
}

// GetCmdIssueMint implements mint a coinIssue transaction command.
func GetCmdIssueMint(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [issue-id] [amount] [to]",
		Args:  cobra.ExactArgs(3),
		Short: "mint a coin",
		Long: strings.TrimSpace(`
mint a coin. For example:
$ hashgardcli issue mint gardh1c7d59vebq 88888
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("Amount %s not a valid int, please input a valid amount", args[2])
			}
			to, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)
			from := cliCtx.GetFromAddress()
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr,
				[]sdk.Msg{msgs.MsgIssueMint{IssueId: issueID, From: from, Amount: amount, To: to}}, false)
		},
	}
	return cmd
}

// GetCmdIssueBurn implements burn a coinIssue transaction command.
func GetCmdIssueBurn(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [issue-id] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "burn a coin",
		Long: strings.TrimSpace(`
mint a coin. For example:
$ hashgardcli issue burn gardh1c7d59vebq 88888
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			num, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("amount %s not a valid int, please input a valid amount", args[1])
			}
			amount := sdk.NewInt(num)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)
			from := cliCtx.GetFromAddress()
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr,
				[]sdk.Msg{msgs.MsgIssueBurn{IssueId: issueID, From: from, Amount: amount}}, false)
		},
	}
	return cmd
}

// GetCmdIssueFinishMinting implements finishMinting a coinIssue transaction command.
func GetCmdIssueFinishMinting(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "finish-minting [issue-id]",
		Args:  cobra.ExactArgs(1),
		Short: "finish-minting a coin",
		Long: strings.TrimSpace(`
mint a coin. For example:
$ hashgardcli issue finish-minting gardh1c7d59vebq
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)
			from := cliCtx.GetFromAddress()
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr,
				[]sdk.Msg{msgs.MsgIssueFinishMinting{IssueId: issueID, From: from}}, false)
		},
	}
	return cmd
}
