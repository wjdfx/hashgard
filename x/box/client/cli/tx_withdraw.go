package cli

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientutils "github.com/hashgard/hashgard/x/box/client/utils"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/types"
	boxutils "github.com/hashgard/hashgard/x/box/utils"
	"github.com/spf13/cobra"
)

// GetCmdBoxWithdraw implements withdraw a box transaction command.
func GetCmdBoxWithdraw(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "withdraw [box-id]",
		Args:    cobra.ExactArgs(1),
		Short:   "Withdraw from a deposit box or future box",
		Long:    "Box holder withdraw from a deposit box or future bo when the box can be withdraw",
		Example: fmt.Sprintf("$ hashgardcli box disable boxab3jlxpt2ps %s --from foo", types.Trade),
		RunE: func(cmd *cobra.Command, args []string) error {
			boxID := args[0]
			if err := boxutils.CheckBoxId(boxID); err != nil {
				return errors.Errorf(err)
			}
			txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
			if err != nil {
				return err
			}
			if account.GetCoins().AmountOf(boxID).IsZero() {
				return errors.Errorf(errors.ErrNotEnoughAmount())
			}
			boxInfo, err := boxutils.GetBoxByID(cdc, cliCtx, boxID)
			if err != nil {
				return err
			}
			switch boxInfo.GetBoxType() {
			case types.Deposit:
				if types.BoxFinished != boxInfo.GetBoxStatus() {
					return errors.Errorf(errors.ErrNotAllowedOperation(boxInfo.GetBoxStatus()))
				}
			case types.Future:
				if types.BoxCreated == boxInfo.GetBoxStatus() {
					return errors.Errorf(errors.ErrNotAllowedOperation(boxInfo.GetBoxStatus()))
				}
				seq := boxutils.GetSeqFromFutureBoxSeq(boxID)
				if boxInfo.GetFuture().TimeLine[seq-1] > time.Now().Unix() {
					return errors.Errorf(errors.ErrNotAllowedOperation(types.BoxUndue))
				}
			default:
				return errors.Errorf(errors.ErrUnknownBox(boxID))
			}
			msg := msgs.NewMsgBoxWithdraw(boxID, account.GetAddress())
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	return cmd
}
