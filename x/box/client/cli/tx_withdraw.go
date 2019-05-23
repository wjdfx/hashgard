package cli

import (
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientutils "github.com/hashgard/hashgard/x/box/client/utils"
	"github.com/spf13/cobra"
)

// GetCmdBoxWithdraw implements withdraw a box transaction command.
func GetCmdBoxWithdraw(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "withdraw [box-id]",
		Args:    cobra.ExactArgs(1),
		Short:   "Withdraw from a deposit box or future box",
		Long:    "Box holder withdraw from a deposit box or future bo when the box can be withdraw",
		Example: "$ hashgardcli box withdraw boxab3jlxpt2ps --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
			if err != nil {
				return err
			}
			msg, err := clientutils.GetWithdrawMsg(cdc, cliCtx, account, args[0])
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	return cmd
}
