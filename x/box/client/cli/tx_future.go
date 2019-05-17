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
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
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
			issueInfo, err := issueutils.GetIssueByID(cdc, cliCtx, coin.Denom)
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
			if err = processFutureBox(coin, futureBox, issueInfo.GetDecimals()); err != nil {
				return err
			}
			coin.Amount = issueutils.MulDecimals(coin.Amount, issueInfo.GetDecimals())
			box := params.BoxFutureParams{}
			box.Sender = account.GetAddress()
			box.Name = args[0]
			box.BoxType = types.Future
			box.TotalAmount = types.BoxToken{Token: coin, Decimals: issueInfo.GetDecimals()}
			box.TradeDisabled = viper.GetBool(flagTradeDisabled)
			box.Future = futureBox
			box.Future.MiniMultiple = uint(viper.GetInt(flagMiniMultiple))
			msg := msgs.NewMsgFutureBox(&box)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	cmd.Flags().Bool(flagTradeDisabled, true, "Disable the box trade")
	cmd.Flags().Uint(flagMiniMultiple, 1, "Trade mini multiple")
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
			total = total.Add(amount)
			futureBox.Receivers[i][j] = issueutils.MulDecimals(amount, decimals).String()
		}
	}
	if !total.Equal(totalAmount.Amount) {
		return errors.ErrAmountNotValid("Receivers")
	}
	return nil
}
