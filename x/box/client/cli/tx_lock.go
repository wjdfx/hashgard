package cli

import (
	"strconv"

	"github.com/hashgard/hashgard/x/box/params"

	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientutils "github.com/hashgard/hashgard/x/box/client/utils"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
	"github.com/spf13/cobra"

	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/types"
)

// GetCmdBox implements create lock box transaction command.
func GetCmdLockBoxCreate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-lock [name] [total-amount] [end-time]",
		Args:    cobra.ExactArgs(3),
		Short:   "Create a new lock box",
		Long:    "Create a new lock box",
		Example: "$ hashgardcli box create-lock foocoin 100000000coin174876e800 2557223200 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			// parse coins trying to be sent
			coin, err := sdk.ParseCoin(args[1])
			if err != nil {
				return err
			}

			txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
			if err != nil {
				return err
			}

			issueInfo, err := issueutils.GetIssueByID(cdc, cliCtx, coin.Denom)
			if err != nil {
				return err
			}

			endTime, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return err
			}

			coin.Amount = issueutils.MulDecimals(coin.Amount, issueInfo.GetDecimals())

			box := &params.BoxLockParams{}
			box.Sender = account.GetAddress()
			box.Name = args[0]
			box.BoxType = types.Lock
			box.TotalAmount = types.BoxToken{Token: coin, Decimals: issueInfo.GetDecimals()}
			box.Lock = types.LockBox{EndTime: endTime}

			msg := msgs.NewMsgLockBox(box)

			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	return cmd
}
