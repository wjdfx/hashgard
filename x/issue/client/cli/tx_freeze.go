package cli

import (
	"fmt"
	"strconv"

	"github.com/hashgard/hashgard/x/issue/types"

	"github.com/hashgard/hashgard/x/issue/msgs"

	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientutils "github.com/hashgard/hashgard/x/issue/client/utils"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
	"github.com/spf13/cobra"

	"github.com/hashgard/hashgard/x/issue/errors"
)

// GetCmdIssueUnFreeze implements freeze a token transaction command.
func GetCmdIssueFreeze(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "freeze [freeze-type] [issue-id] [acc-address] [end-time]",
		Args:  cobra.ExactArgs(4),
		Short: "Freeze the transfer from a address",
		Long: "Token owner freeze the transfer from a address:\n\n" +
			types.FreezeIn + ":The address can not transfer in\n" +
			types.FreezeOut + ":The address can not transfer out\n" +
			types.FreezeInAndOut + ":The address not can transfer in and out\n\n" +
			"Note:The end-time you can use bash to get: date -d \"2020-01-01 10:30:00\" +%s",
		Example: "$ hashgardcli issue freeze in coin174876e800 gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n 1577845800 --from foo\n" +
			"$ hashgardcli issue freeze out coin174876e800 gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n 1577845800 --from foo\n" +
			"$ hashgardcli issue freeze in-out coin174876e800 gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n 1577845800 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {

			return issueFreeze(cdc, args, true)
		},
	}
	return cmd
}

// GetCmdIssueUnFreeze implements un freeze  a token transaction command.
func GetCmdIssueUnFreeze(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unfreeze [freeze-type] [issue-id] [acc-address]",
		Args:  cobra.ExactArgs(3),
		Short: "UnFreeze the transfer from a address",
		Long: "Token owner unFreeze the transfer from a address:\n\n" +
			types.FreezeIn + ":The address can transfer in\n" +
			types.FreezeOut + ":The address can transfer out\n" +
			types.FreezeInAndOut + ":The address can transfer in and out",
		Example: "$ hashgardcli issue unfreeze in coin174876e800 gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n --from foo\n" +
			"$ hashgardcli issue unfreeze out coin174876e800 gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n --from foo\n" +
			"$ hashgardcli issue unfreeze in-out coin174876e800 gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {

			return issueFreeze(cdc, args, false)
		},
	}
	return cmd
}

func issueFreeze(cdc *codec.Codec, args []string, freeze bool) error {

	freezeType := args[0]

	_, ok := types.FreezeType[freezeType]
	if !ok {
		return errors.ErrUnknownFreezeType()
	}

	issueID := args[1]
	if err := issueutils.CheckIssueId(issueID); err != nil {
		return errors.Errorf(err)
	}
	accAddress, err := sdk.AccAddressFromBech32(args[2])
	if err != nil {
		return err
	}

	txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
	if err != nil {
		return err
	}

	_, err = issueutils.IssueOwnerCheck(cdc, cliCtx, account, issueID)
	if err != nil {
		return err
	}

	var msg sdk.Msg

	if freeze {
		endTime, err := strconv.ParseInt(args[3], 10, 64)

		if err != nil {
			return fmt.Errorf("EndTime %s not a valid int, please input a valid EndTime", args[2])
		}
		msg = msgs.NewMsgIssueFreeze(issueID, account.GetAddress(), accAddress, freezeType, endTime)
	} else {
		msg = msgs.NewMsgIssueUnFreeze(issueID, account.GetAddress(), accAddress, freezeType)
	}

	validateErr := msg.ValidateBasic()
	if validateErr != nil {
		return errors.Errorf(validateErr)
	}
	return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
}
