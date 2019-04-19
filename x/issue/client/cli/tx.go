package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/types"
)

// GetCmdIssue implements issue a coin transaction command.
func GetCmdIssueCreate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create [name] [symbol] [total-supply]",
		Args:    cobra.ExactArgs(3),
		Short:   "Issue a new coin",
		Long:    "Issue a new coin",
		Example: "$ hashgardcli issue create foocoin FOO 100000000 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			totalSupply, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("Total supply %s not a valid int, please input a valid total supply", args[2])
			}

			txBldr, cliCtx, account, err := issueutils.GetCliContext(cdc)
			if err != nil {
				return err
			}

			coinIssueInfo := types.CoinIssueInfo{
				Issuer:          account.GetAddress(),
				Owner:           account.GetAddress(),
				Name:            args[0],
				Symbol:          strings.ToUpper(args[1]),
				IssueTime:       time.Now(),
				BurnOff:         viper.GetBool(flagBurnOff),
				BurnFromOff:     viper.GetBool(flagBurnFromOff),
				BurnAnyOff:      viper.GetBool(flagBurnAnyOff),
				MintingFinished: viper.GetBool(flagMintingFinished),
				TotalSupply:     totalSupply,
				Decimals:        uint(viper.GetInt(flagDecimals)),
			}
			coinIssueInfo.SetTotalSupply(issueutils.MulDecimals(coinIssueInfo.TotalSupply, coinIssueInfo.Decimals))
			msg := msgs.CreateMsgIssue(&coinIssueInfo)

			validateErr := msg.ValidateBasic()

			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd.Flags().Uint(flagDecimals, types.CoinDecimalsMaxValue, "Decimals of coin")
	cmd.Flags().Bool(flagBurnOff, false, "can burning of coin")
	cmd.Flags().Bool(flagBurnFromOff, false, "can burning of coin from account")
	cmd.Flags().Bool(flagBurnAnyOff, false, "can burning of coin from any account by owner")
	cmd.Flags().Bool(flagMintingFinished, false, "can minting of coin")

	return cmd
}

// GetCmdIssueTransferOwnership implements transfer a coin owner ship transaction command.
func GetCmdIssueTransferOwnership(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "transfer-ownership [issue-id] [to_address]",
		Args:    cobra.ExactArgs(2),
		Short:   "Transfer-ownership a coin",
		Long:    "Transfer-ownership a coin",
		Example: "$ hashgardcli issue transfer-ownership coin155547350020 gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			txBldr, cliCtx, account, err := issueutils.GetCliContext(cdc)
			if err != nil {
				return err
			}
			to, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			_, err = issueutils.IssueOwnerCheck(cdc, cliCtx, account, issueID)
			if err != nil {
				return err
			}
			msg := msgs.NewMsgIssueTransferOwnership(issueID, account.GetAddress(), to)

			validateErr := msg.ValidateBasic()

			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	return cmd
}

// GetCmdIssue implements issue a coin transaction command.
func GetCmdIssueDescription(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "describe [issue-id] [description-file]",
		Args:    cobra.ExactArgs(2),
		Short:   "Describe a coin",
		Long:    "Describe a coin",
		Example: "$ hashgardcli issue describe coin155547350020 path/description.json --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			txBldr, cliCtx, account, err := issueutils.GetCliContext(cdc)
			if err != nil {
				return err
			}
			contents, err := ioutil.ReadFile(args[1])
			if err != nil {
				return err
			}
			buffer := bytes.Buffer{}
			err = json.Compact(&buffer, contents)
			if err != nil {
				return errors.ErrCoinDescriptionNotValid()
			}
			contents = buffer.Bytes()

			_, err = issueutils.IssueOwnerCheck(cdc, cliCtx, account, issueID)
			if err != nil {
				return err
			}
			msg := msgs.NewMsgIssueDescription(issueID, account.GetAddress(), contents)

			validateErr := msg.ValidateBasic()

			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	return cmd
}

// GetCmdIssueMint implements mint a coinIssue transaction command.
func GetCmdIssueMint(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [issue-id] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "mint a coin",
		Long:  "mint a coin to a address",
		Example: "$ hashgardcli issue mint coin155547350020 88888 --from foo\n" +
			"$ hashgardcli issue mint coin155547350020 88888 --to=gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("Amount %s not a valid int, please input a valid amount", args[2])
			}

			txBldr, cliCtx, account, err := issueutils.GetCliContext(cdc)
			if err != nil {
				return err
			}
			to := account.GetAddress()
			flagTo := viper.GetString(flagMintTo)
			if len(flagTo) > 0 {
				to, err = sdk.AccAddressFromBech32(flagTo)
				if err != nil {
					return err
				}
			}

			issueInfo, err := issueutils.IssueOwnerCheck(cdc, cliCtx, account, issueID)
			if err != nil {
				return err
			}

			amount = issueutils.MulDecimals(amount, issueInfo.GetDecimals())

			msg := msgs.MsgIssueMint{IssueId: issueID, Operator: account.GetAddress(), Amount: amount, Decimals: issueInfo.GetDecimals(), To: to}
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	cmd.Flags().String(flagMintTo, "", "Mint to account address")
	return cmd
}

// GetCmdIssueBurn implements burn a coinIssue transaction command.
func GetCmdIssueBurn(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "burn [issue-id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "burn a coin",
		Long:    "burn a coin",
		Example: "$ hashgardcli issue burn coin155547350020 88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("Amount %s not a valid int, please input a valid amount", args[1])
			}

			txBldr, cliCtx, account, err := issueutils.GetCliContext(cdc)
			if err != nil {
				return err
			}
			amount, err = issueutils.BurnCheck(cdc, cliCtx, account, nil, issueID, amount, types.BurnOwner)
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr,
				[]sdk.Msg{msgs.MsgIssueBurn{IssueId: issueID, Operator: account.GetAddress(), Amount: amount}}, false)
		},
	}
	return cmd
}

// GetCmdIssueBurnFrom implements burn a coinIssue transaction command.
func GetCmdIssueBurnFrom(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "burn-from [issue-id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "burn a coin from my account",
		Long:    "burn a coin from my account",
		Example: "$ hashgardcli issue burn-from gardh1c7d59vebq 88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("Amount %s not a valid int, please input a valid amount", args[1])
			}

			txBldr, cliCtx, account, err := issueutils.GetCliContext(cdc)
			if err != nil {
				return err
			}

			amount, err = issueutils.BurnCheck(cdc, cliCtx, account, account.GetAddress(), issueID, amount, types.BurnFrom)
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr,
				[]sdk.Msg{msgs.MsgIssueBurnFrom{IssueId: issueID, Operator: account.GetAddress(), From: account.GetAddress(), Amount: amount}}, false)
		},
	}
	return cmd
}

// GetCmdIssueBurnAny implements burn a coinIssue transaction command.
func GetCmdIssueBurnAny(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "burn-any [issue-id] [address] [amount]",
		Args:    cobra.ExactArgs(3),
		Short:   "burn a coin from any address",
		Long:    "burn a coin from any address",
		Example: "$ hashgardcli issue burn-any gardh1c7d59vebq gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n 88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			accAddress, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("Amount %s not a valid int, please input a valid amount", args[1])
			}

			txBldr, cliCtx, account, err := issueutils.GetCliContext(cdc)
			if err != nil {
				return err
			}

			amount, err = issueutils.BurnCheck(cdc, cliCtx, account, accAddress, issueID, amount, types.BurnAny)
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr,
				[]sdk.Msg{msgs.MsgIssueBurnAny{IssueId: issueID, Operator: account.GetAddress(), From: accAddress, Amount: amount}}, false)
		},
	}
	return cmd
}

// GetCmdIssueBurnOff implements burnOff a coinIssue transaction command.
func GetCmdIssueBurnOff(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "burn-off [issue-id]",
		Args:    cobra.ExactArgs(1),
		Short:   "burn-off a coin",
		Long:    "burn-off a coin",
		Example: "$ hashgardcli issue burn-off gardh1c7d59vebq --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getIssueFlagCmd(cdc, cmd, args, msgs.MsgIssueBurnOff{})
		},
	}
	return cmd
}

// GetCmdIssueBurnFromOff implements burnFromOff a coinIssue transaction command.
func GetCmdIssueBurnFromOff(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "burn-from-off [issue-id]",
		Args:    cobra.ExactArgs(1),
		Short:   "burn-from-off a coin",
		Long:    "burn-from-off a coin",
		Example: "$ hashgardcli issue burn-from-off gardh1c7d59vebq --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getIssueFlagCmd(cdc, cmd, args, msgs.MsgIssueBurnFromOff{})
		},
	}
	return cmd
}

// GetCmdIssueBurnAnyOff implements burnAnyOff a coinIssue transaction command.
func GetCmdIssueBurnAnyOff(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "burn-any-off [issue-id]",
		Args:    cobra.ExactArgs(1),
		Short:   "burn-any-off a coin",
		Long:    "burn-any-off a coin",
		Example: "$ hashgardcli issue burn-any-off gardh1c7d59vebq --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getIssueFlagCmd(cdc, cmd, args, msgs.MsgIssueBurnAnyOff{})
		},
	}
	return cmd
}

// GetCmdIssueFinishMinting implements finishMinting a coinIssue transaction command.
func GetCmdIssueFinishMinting(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "finish-minting [issue-id]",
		Args:    cobra.ExactArgs(1),
		Short:   "finish-minting a coin",
		Long:    "finish-minting a coin",
		Example: "$ hashgardcli issue finish-minting gardh1c7d59vebq --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getIssueFlagCmd(cdc, cmd, args, msgs.MsgIssueFinishMinting{})
		},
	}
	return cmd
}

func getIssueFlagCmd(cdc *codec.Codec, cmd *cobra.Command, args []string, msg msgs.MsgFlag) error {
	issueID := args[0]
	if err := issueutils.CheckIssueId(issueID); err != nil {
		return errors.Errorf(err)
	}
	txBldr, cliCtx, account, err := issueutils.GetCliContext(cdc)
	if err != nil {
		return err
	}
	_, err = issueutils.IssueOwnerCheck(cdc, cliCtx, account, issueID)
	if err != nil {
		return err
	}

	switch msg.(type) {
	case msgs.MsgIssueBurnOff:
		msg = msgs.NewMsgIssueBurnOff(issueID, account.GetAddress())
	case msgs.MsgIssueBurnFromOff:
		msg = msgs.NewMsgIssueBurnFromOff(issueID, account.GetAddress())
	case msgs.MsgIssueBurnAnyOff:
		msg = msgs.NewMsgIssueBurnAnyOff(issueID, account.GetAddress())
	case msgs.MsgIssueFinishMinting:
		msg = msgs.NewMsgIssueFinishMinting(issueID, account.GetAddress())
	default:
		return nil
	}

	return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
}
