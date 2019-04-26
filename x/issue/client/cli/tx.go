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
		Short:   "Issue a new token",
		Long:    "Issue a new token",
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
				Issuer:             account.GetAddress(),
				Owner:              account.GetAddress(),
				Name:               args[0],
				Symbol:             strings.ToUpper(args[1]),
				IssueTime:          time.Now(),
				BurnOwnerDisabled:  viper.GetBool(flagBurnOwnerDisabled),
				BurnHolderDisabled: viper.GetBool(flagBurnHolderDisabled),
				BurnFromDisabled:   viper.GetBool(flagBurnFromDisabled),
				MintingFinished:    viper.GetBool(flagMintingFinished),
				TotalSupply:        totalSupply,
				Decimals:           uint(viper.GetInt(flagDecimals)),
			}
			coinIssueInfo.SetTotalSupply(issueutils.MulDecimals(coinIssueInfo.TotalSupply, coinIssueInfo.Decimals))
			msg := msgs.NewMsgIssue(&coinIssueInfo)

			validateErr := msg.ValidateBasic()

			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd.Flags().Uint(flagDecimals, types.CoinDecimalsMaxValue, "Decimals of the token")
	cmd.Flags().Bool(flagBurnOwnerDisabled, false, "Disable token owner burn the token")
	cmd.Flags().Bool(flagBurnHolderDisabled, false, "Disable token holder burn the token")
	cmd.Flags().Bool(flagBurnFromDisabled, false, "Disable token owner burn the token from any holder")
	cmd.Flags().Bool(flagMintingFinished, false, "Token owner can not minting the token")

	return cmd
}

// GetCmdIssueTransferOwnership implements transfer a coin owner ship transaction command.
func GetCmdIssueTransferOwnership(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "transfer-ownership [issue-id] [to_address]",
		Args:    cobra.ExactArgs(2),
		Short:   "Transfer ownership a token",
		Long:    "Token owner transfer the ownership to new account",
		Example: "$ hashgardcli issue transfer-ownership coin174876e800 gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc --from foo",
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
		Short:   "Describe a token",
		Long:    "Owner can add description of the token issued by owner, and the description need to be in json format. You can customize preferences or use recommended templates.",
		Example: "$ hashgardcli issue describe coin174876e800 path/description.json --from foo",
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
		Short: "Mint a token",
		Long:  "Token owner mint the token to a address",
		Example: "$ hashgardcli issue mint coin174876e800 88888 --from foo\n" +
			"$ hashgardcli issue mint coin174876e800 88888 --to=gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc --from foo",
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

			if issueInfo.IsMintingFinished() {
				return errors.Errorf(errors.ErrCanNotMint(issueID))
			}

			amount = issueutils.MulDecimals(amount, issueInfo.GetDecimals())

			msg := msgs.MsgIssueMint{IssueId: issueID, Sender: account.GetAddress(), Amount: amount, Decimals: issueInfo.GetDecimals(), To: to}
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

// GetCmdIssueBurnFrom implements burn a coinIssue transaction command.
func GetCmdIssueBurn(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "burn [issue-id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Token holder burn the token",
		Long:    "Token holder or the Owner burn the token he holds (the Owner can burn if 'burning_owner_disabled' is false, the holder can burn if 'burning_holder_disabled' is false)",
		Example: "$ hashgardcli issue burn coin174876e800 88888 --from foo",
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

			issueInfo, err := issueutils.GetIssueByID(cdc, cliCtx, issueID)
			if err != nil {
				return err
			}

			burnType := types.BurnHolder

			if issueInfo.GetOwner().Equals(account.GetAddress()) {
				burnType = types.BurnOwner
			}
			amount, err = issueutils.BurnCheck(account, account.GetAddress(), issueInfo, amount, burnType)
			if err != nil {
				return err
			}

			var msg sdk.Msg

			if types.BurnOwner == burnType {
				msg = msgs.NewMsgIssueBurnOwner(issueID, account.GetAddress(), amount)
			} else {
				msg = msgs.NewMsgIssueBurnHolder(issueID, account.GetAddress(), amount)
			}

			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return errors.Errorf(validateErr)
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	return cmd
}

// GetCmdIssueBurnFrom implements burn a coinIssue transaction command.
func GetCmdIssueBurnFrom(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "burn-from [issue-id] [address] [amount]",
		Args:    cobra.ExactArgs(3),
		Short:   "Token owner burn the token",
		Long:    "Token Owner burn the token from any holder (the Owner can burn if 'burning_any_disabled' is false)",
		Example: "$ hashgardcli issue burn-from coin174876e800 gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n 88888 --from foo",
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

			issueInfo, err := issueutils.GetIssueByID(cdc, cliCtx, issueID)
			if err != nil {
				return err
			}

			amount, err = issueutils.BurnCheck(account, accAddress, issueInfo, amount, types.BurnFrom)
			if err != nil {
				return err
			}
			msg := msgs.NewMsgIssueBurnFrom(issueID, account.GetAddress(), accAddress, amount)

			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return errors.Errorf(validateErr)
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	return cmd
}

// GetCmdIssueDisableFeature implements disable feature a coinIssue transaction command.
func GetCmdIssueDisableFeature(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disable [issue-id] [feature]",
		Args:  cobra.ExactArgs(2),
		Short: "Disable feature from a token",
		Long: "Owner disabled the features:\n" +
			types.BurnOwner + ":Token owner can burn the token\n" +
			types.BurnHolder + ":Token holder can burn the token\n" +
			types.BurnFrom + ":Token owner can burn the token from any holder\n" +
			types.Minting + ":Token owner can mint the token",
		Example: "$ hashgardcli issue disable coin174876e800 " + types.BurnOwner + " --from foo\n" +
			"$ hashgardcli issue disable coin174876e800 " + types.BurnHolder + " --from foo\n" +
			"$ hashgardcli issue disable coin174876e800 " + types.BurnFrom + " --from foo\n" +
			"$ hashgardcli issue disable coin174876e800 " + types.Minting + " --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			feature := args[1]

			_, ok := types.Features[feature]
			if !ok {
				return errors.ErrUnknownFeatures()
			}

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

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msgs.NewMsgIssueDisableFeature(issueID, account.GetAddress(), feature)}, false)
		},
	}
	return cmd
}

// GetCmdIssueApprove implements approve a token transaction command.
func GetCmdIssueApprove(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "approve [issue-id] [address] [amount]",
		Args:    cobra.ExactArgs(3),
		Short:   "Approve spend tokens on behalf of sender",
		Long:    "Approve the passed address to spend the specified amount of tokens on behalf of sender",
		Example: "$ hashgardcli issue approve coin174876e800 gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n 88888 --from foo",
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

			issueInfo, err := issueutils.GetIssueByID(cdc, cliCtx, issueID)
			if err != nil {
				return err
			}
			amount = issueutils.MulDecimals(amount, issueInfo.GetDecimals())

			msg := msgs.NewMsgIssueApprove(issueID, account.GetAddress(), accAddress, amount)

			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return errors.Errorf(validateErr)
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	return cmd
}

// GetCmdIssueSendFrom implements send from a token transaction command.
func GetCmdIssueSendFrom(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-from [issue-id] [from_address] [to_address] [amount]",
		Args:  cobra.ExactArgs(4),
		Short: "Send tokens from one address to another",
		Long:  "Send tokens from one address to another by allowance",
		Example: "$ hashgardcli issue send-from coin174876e800 gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n gard1vud9ptwagudgq7yht53cwuf8qfmgkd0qcej0ah " +
			"88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			fromAddress, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}
			toAddress, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[3])
			if !ok {
				return fmt.Errorf("Amount %s not a valid int, please input a valid amount", args[1])
			}

			txBldr, cliCtx, account, err := issueutils.GetCliContext(cdc)
			if err != nil {
				return err
			}

			if err := issueutils.CheckAllowance(cdc, cliCtx, issueID, fromAddress, account.GetAddress(), amount); err != nil {
				return err
			}

			issueInfo, err := issueutils.GetIssueByID(cdc, cliCtx, issueID)
			if err != nil {
				return err
			}
			amount = issueutils.MulDecimals(amount, issueInfo.GetDecimals())

			msg := msgs.NewMsgIssueSendFrom(issueID, account.GetAddress(), fromAddress, toAddress, amount)

			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return errors.Errorf(validateErr)
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	return cmd
}

// GetCmdIssueIncreaseApproval implements increase approval a token transaction command.
func GetCmdIssueIncreaseApproval(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "increase-approval [issue-id] [address] [amount]",
		Args:    cobra.ExactArgs(3),
		Short:   "Increase approve spend tokens on behalf of sender",
		Long:    "Increase approve the passed address to spend the specified amount of tokens on behalf of sender",
		Example: "$ hashgardcli issue increase-approval coin174876e800 gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n 88888 --from foo",
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

			issueInfo, err := issueutils.GetIssueByID(cdc, cliCtx, issueID)
			if err != nil {
				return err
			}
			amount = issueutils.MulDecimals(amount, issueInfo.GetDecimals())

			msg := msgs.NewMsgIssueIncreaseApproval(issueID, account.GetAddress(), accAddress, amount)

			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return errors.Errorf(validateErr)
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	return cmd
}

// GetCmdIssueDecreaseApproval implements decrease approval a token transaction command.
func GetCmdIssueDecreaseApproval(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "decrease-approval [issue-id] [address] [amount]",
		Args:    cobra.ExactArgs(3),
		Short:   "Decrease approve spend tokens on behalf of sender",
		Long:    "Decrease approve the passed address to spend the specified amount of tokens on behalf of sender",
		Example: "$ hashgardcli issue increase-approval coin174876e800 gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n 88888 --from foo",
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

			issueInfo, err := issueutils.GetIssueByID(cdc, cliCtx, issueID)
			if err != nil {
				return err
			}
			amount = issueutils.MulDecimals(amount, issueInfo.GetDecimals())

			msg := msgs.NewMsgIssueDecreaseApproval(issueID, account.GetAddress(), accAddress, amount)

			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return errors.Errorf(validateErr)
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	return cmd
}
