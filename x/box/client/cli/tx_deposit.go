package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/x/auth"

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

// GetCmdDepositToBox implements deposit to a box transaction command.
func GetCmdDepositToBox(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deposit-to [box-id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Deposit to the box",
		Long:    "Deposit to the box",
		Example: "$ hashgardcli box deposit-to box174876e800 88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deposit(cdc, args, types.DepositTo)
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
			return deposit(cdc, args, types.Fetch)
		},
	}
	return cmd
}
func deposit(cdc *codec.Codec, args []string, operation string) error {

	boxID := args[0]
	if err := boxutils.CheckBoxId(boxID); err != nil {
		return errors.Errorf(err)
	}
	amountArg, ok := sdk.NewIntFromString(args[1])
	if !ok {
		return fmt.Errorf("Amount %s not a valid int, please input a valid amount", args[2])
	}
	txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
	if err != nil {
		return err
	}
	boxInfo, err := boxutils.GetBoxByID(cdc, cliCtx, boxID)
	if err != nil {
		return err
	}

	decimal, err := clientutils.GetCoinDecimal(cdc, cliCtx, boxInfo.GetTotalAmount().Token)
	if err != nil {
		return err
	}
	amount := boxutils.MulDecimals(boxutils.ParseCoin(boxInfo.GetTotalAmount().Token.Denom, amountArg), decimal)

	switch operation {
	case types.DepositTo:
		if err = checkAmountByDepositTo(amount, boxInfo); err != nil {
			return err
		}
	case types.Fetch:
		if err = checkAmountByFetch(amount, boxInfo, account); err != nil {
			return err
		}
	default:
		return errors.ErrNotSupportOperation()
	}
	msg := msgs.NewMsgBoxDeposit(boxID, account.GetAddress(), sdk.NewCoin(boxInfo.GetTotalAmount().Token.Denom, amount), operation)

	validateErr := msg.ValidateBasic()
	if validateErr != nil {
		return errors.Errorf(validateErr)
	}
	return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)

}

func checkAmountByFetch(amount sdk.Int, boxInfo types.Box, account auth.Account) error {
	switch boxInfo.GetBoxType() {
	case types.Deposit:
		if !amount.Mod(boxInfo.GetDeposit().Price).IsZero() {
			return errors.ErrAmountNotValid(amount.String())
		}
		if account.GetCoins().AmountOf(boxInfo.GetBoxId()).LT(amount.Quo(boxInfo.GetDeposit().Price)) {
			return errors.Errorf(errors.ErrNotEnoughAmount())
		}
	case types.Future:
		if boxInfo.GetFuture().Deposits == nil {
			return errors.Errorf(errors.ErrNotEnoughAmount())
		}
		for _, v := range boxInfo.GetFuture().Deposits {
			if v.Address.Equals(account.GetAddress()) {
				if v.Amount.GTE(amount) {
					return nil
				}
			}
		}
		return errors.Errorf(errors.ErrNotEnoughAmount())
	default:
		return errors.Errorf(errors.ErrNotSupportOperation())
	}
	return nil
}
func checkAmountByDepositTo(amount sdk.Int, boxInfo types.Box) error {
	switch boxInfo.GetBoxType() {
	case types.Deposit:
		if !amount.Mod(boxInfo.GetDeposit().Price).IsZero() {
			return errors.ErrAmountNotValid(amount.String())
		}
		if amount.Add(boxInfo.GetDeposit().TotalDeposit).GT(boxInfo.GetTotalAmount().Token.Amount) {
			return errors.Errorf(errors.ErrNotEnoughAmount())
		}
	case types.Future:
		total := sdk.ZeroInt()
		if boxInfo.GetFuture().Deposits != nil {
			for _, v := range boxInfo.GetFuture().Deposits {
				total = total.Add(v.Amount)
			}
		}
		if amount.Add(total).GT(boxInfo.GetTotalAmount().Token.Amount) {
			return errors.Errorf(errors.ErrNotEnoughAmount())
		}
	default:
		return errors.Errorf(errors.ErrNotSupportOperation())
	}
	return nil
}
