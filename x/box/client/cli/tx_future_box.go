package cli

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hashgard/hashgard/x/box/params"

	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientutils "github.com/hashgard/hashgard/x/box/client/utils"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/types"
	boxutils "github.com/hashgard/hashgard/x/box/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetCmdFutureBoxCreate implements create Future box transaction command.
func GetCmdFutureBoxCreate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-future [name] [total-amount] [mini-multiple] [distribute-file]",
		Args:    cobra.ExactArgs(4),
		Short:   "Create a new future box",
		Long:    "Create a new future box",
		Example: "$ hashgardcli box create-future foocoin 100000000coin174876e800 1 path/distribute.json --from foo",
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
			decimal, err := clientutils.GetCoinDecimal(cdc, cliCtx, coin)
			if err != nil {
				return err
			}
			contents, err := ioutil.ReadFile(args[3])
			if err != nil {
				return err
			}

			futureBox := types.FutureBox{}
			err = json.Unmarshal(contents, &futureBox)
			if err != nil {
				return err
			}
			coin.Amount = boxutils.MulDecimals(coin, decimal)
			if err = processFutureBox(coin, futureBox, decimal); err != nil {
				return err
			}
			box := params.BoxFutureParams{}
			box.Sender = account.GetAddress()
			box.Name = args[0]
			box.BoxType = types.Future
			box.TotalAmount = types.BoxToken{Token: coin, Decimals: decimal}
			box.TransferDisabled = viper.GetBool(flagTransferDisabled)
			box.Future = futureBox
			msg := msgs.NewMsgFutureBox(&box)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	cmd.Flags().Bool(flagTransferDisabled, true, "Disable transfer the box")
	return cmd
}

func processFutureBox(totalAmount sdk.Coin, futureBox types.FutureBox, decimals uint) sdk.Error {
	if futureBox.Receivers == nil {
		return errors.ErrNotSupportOperation()
	}
	total := sdk.ZeroInt()
	for i, items := range futureBox.Receivers {
		for j, rec := range items {
			if j == 0 {
				_, err := sdk.AccAddressFromBech32(rec)
				if err != nil {
					return sdk.ErrInvalidAddress(rec)
				}
				continue
			}
			amount, ok := sdk.NewIntFromString(rec)
			if !ok {
				return errors.ErrAmountNotValid(rec)
			}
			amount = boxutils.MulDecimals(boxutils.ParseCoin(totalAmount.Denom, amount), decimals)
			total = total.Add(amount)
			futureBox.Receivers[i][j] = amount.String()
		}
	}
	if !total.Equal(totalAmount.Amount) {
		return errors.ErrAmountNotValid("Receivers")
	}
	return nil
}
