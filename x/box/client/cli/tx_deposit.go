package cli

import (
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientutils "github.com/hashgard/hashgard/x/box/client/utils"
	"github.com/hashgard/hashgard/x/box/types"
	"github.com/spf13/cobra"
)

// GetCmdDepositToBox implements deposit to a box transaction command.
func GetCmdDepositToBox(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deposit-to [box-id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Deposit to the box",
		Long:    "Deposit to the box",
		Example: "$ hashgardcli box deposit-to box174876e800 88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deposit(cdc, args[0], args[1], types.DepositTo)
		},
	}
	return cmd
}

// GetCmdFetchDepositFromBox implements fetch deposit from a box transaction command.
func GetCmdFetchDepositFromBox(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deposit-fetch [box-id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Fetch deposit from a deposit box",
		Long:    "Fetch deposit from a deposit box",
		Example: "$ hashgardcli box deposit-fetch box174876e800 88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deposit(cdc, args[0], args[1], types.Fetch)
		},
	}
	return cmd
}
func deposit(cdc *codec.Codec, boxID string, amountStr string, operation string) error {
	txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
	if err != nil {
		return err
	}
	msg, err := clientutils.GetDepositMsg(cdc, cliCtx, account, boxID, amountStr, operation, true)
	if err != nil {
		return err
	}
	return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
}
