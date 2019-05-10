package cli

import (
	"fmt"
	"time"

	"github.com/hashgard/hashgard/x/box/client/queriers"

	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientutils "github.com/hashgard/hashgard/x/box/client/utils"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/types"
	boxutils "github.com/hashgard/hashgard/x/box/utils"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetCmdBox implements create deposit box transaction command.
func GetCmdDepositBoxCreate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-deposit [name] [total-amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Create a new deposit box",
		Long:    "Create a new deposit box",
		Example: "$ hashgardcli box create-deposit foocoin 100000000coin174876e800 --from foo",
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

			boxInfo := types.BoxInfo{}

			boxInfo.Owner = account.GetAddress()
			boxInfo.Name = args[0]
			boxInfo.CreatedTime = time.Now()
			boxInfo.TotalAmount = coin
			boxInfo.BoxType = types.Deposit
			boxInfo.TradeDisabled = viper.GetBool(flagTradeDisabled)
			boxInfo.Deposit = types.DepositBox{
				Share:         sdk.ZeroInt(),
				TotalDeposit:  sdk.ZeroInt(),
				Coupon:        sdk.ZeroInt(),
				StartTime:     time.Unix(viper.GetInt64(flagStartTime), 0),
				EstablishTime: time.Unix(viper.GetInt64(flagEstablishTime), 0),
				MaturityTime:  time.Unix(viper.GetInt64(flagMaturityTime), 0)}

			num, ok := sdk.NewIntFromString(viper.GetString(flagBottomLine))
			if !ok {
				return errors.Errorf(errors.ErrAmountNotValid(flagBottomLine))
			}
			boxInfo.Deposit.BottomLine = num
			num, ok = sdk.NewIntFromString(viper.GetString(flagPrice))
			if !ok {
				return errors.Errorf(errors.ErrAmountNotValid(flagPrice))
			}
			boxInfo.Deposit.Price = num
			boxInfo.Deposit.Interest, err = sdk.ParseCoin(viper.GetString(flagInterest))
			if err != nil {
				return err
			}
			boxInfo.Deposit.Price = issueutils.MulDecimals(boxInfo.GetDeposit().Price, issueInfo.GetDecimals())
			boxInfo.Deposit.BottomLine = issueutils.MulDecimals(boxInfo.GetDeposit().BottomLine, issueInfo.GetDecimals())
			boxInfo.TotalAmount.Amount = issueutils.MulDecimals(boxInfo.TotalAmount.Amount, issueInfo.GetDecimals())

			issueInfo, err = issueutils.GetIssueByID(cdc, cliCtx, boxInfo.Deposit.Interest.Denom)
			if err == nil {
				boxInfo.Deposit.Interest.Amount = issueutils.MulDecimals(boxInfo.Deposit.Interest.Amount, issueInfo.GetDecimals())
			}

			msg := msgs.NewMsgBox(&boxInfo)
			validateErr := msg.ValidateBasic()

			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	cmd.Flags().Bool(flagTradeDisabled, true, "Disable the box trade")
	cmd.Flags().String(flagBottomLine, "", "Box bottom line")
	cmd.Flags().String(flagPrice, "", "Box unit price")
	cmd.Flags().String(flagInterest, "", "Box interest")
	cmd.Flags().Int64(flagStartTime, 0, "Box start time")
	cmd.Flags().Int64(flagEstablishTime, 0, "Box establish time")
	cmd.Flags().Int64(flagMaturityTime, 0, "Box maturity time")

	return cmd
}

// GetCmdDepositBoxInterestInjection implements interest injection a deposit box transaction command.
func GetCmdDepositBoxInterestInjection(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "interest-injection [box-id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Injection interest to the deposit box",
		Long:    "Injection interest to the deposit box",
		Example: "$ hashgardcli box interest-injection box174876e800 88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return interest(cdc, args, types.Injection)
		},
	}
	return cmd
}

// GetCmdDepositBoxInterestFetch implements fetch interest from a deposit box transaction command.
func GetCmdDepositBoxInterestFetch(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "interest-fetch [box-id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Fetch interest from a deposit box",
		Long:    "Fetch interest from a deposit box",
		Example: "$ hashgardcli box interest-fetch box174876e800 88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return interest(cdc, args, types.Fetch)
		},
	}
	return cmd
}

// GetCmdDepositToBoxDeposit implements deposit to a deposit box transaction command.
func GetCmdDepositToBoxDeposit(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deposit-to [box-id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Deposit to the deposit box",
		Long:    "Deposit to the deposit box",
		Example: "$ hashgardcli box deposit-to box174876e800 88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deposit(cdc, args, types.DepositTo)
		},
	}
	return cmd
}

// GetCmdInterestFetch implements fetch interest from a box transaction command.
func GetCmdDepositBoxFetchDeposit(cdc *codec.Codec) *cobra.Command {
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
	if boxInfo.GetBoxType() != types.Deposit {
		return errors.Errorf(errors.ErrNotSupportOperation())
	}
	if boxInfo.GetDeposit().Status != types.DepositBoxDeposit {
		return errors.Errorf(errors.ErrNotAllowedOperation(boxInfo.GetDeposit().Status))
	}

	issueInfo, err := issueutils.GetIssueByID(cdc, cliCtx, boxInfo.GetDeposit().Interest.Denom)
	if err != nil {
		return err
	}

	amount := issueutils.MulDecimals(amountArg, issueInfo.GetDecimals())

	if !amount.Mod(boxInfo.GetDeposit().Price).IsZero() {
		return errors.ErrAmountNotValid(amount.String())
	}

	if types.Fetch == operation {
		res, err := queriers.QueryDepositAmountFromDepositBox(boxID, account.GetAddress(), cliCtx)
		if err == nil {
			var depositAmount sdk.Int
			cdc.MustUnmarshalJSON(res, &depositAmount)
			if depositAmount.LT(amount) {
				return errors.Errorf(errors.ErrNotEnoughAmount())
			}
		}
	} else {
		if boxInfo.GetDeposit().TotalDeposit.Add(amount).GT(boxInfo.GetTotalAmount().Amount) {
			return errors.Errorf(errors.ErrNotEnoughAmount())
		}
	}
	msg := msgs.NewMsgBoxDeposit(boxID, account.GetAddress(), sdk.NewCoin(boxInfo.GetTotalAmount().Denom, amount), operation)

	validateErr := msg.ValidateBasic()
	if validateErr != nil {
		return errors.Errorf(validateErr)
	}
	return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)

}
func interest(cdc *codec.Codec, args []string, operation string) error {

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
	if boxInfo.GetBoxType() != types.Deposit {
		return errors.Errorf(errors.ErrNotSupportOperation())
	}
	if boxInfo.GetDeposit().Status != types.BoxCreated {
		return errors.Errorf(errors.ErrNotSupportOperation())
	}

	issueInfo, err := issueutils.GetIssueByID(cdc, cliCtx, boxInfo.GetDeposit().Interest.Denom)
	if err != nil {
		return err
	}

	amount := issueutils.MulDecimals(amountArg, issueInfo.GetDecimals())

	if types.Fetch == operation {
		flag := true
		for i, v := range boxInfo.GetDeposit().InterestInjection {
			if v.Address.Equals(account.GetAddress()) {
				if boxInfo.GetDeposit().InterestInjection[i].Amount.GTE(amount) {
					flag = false
					break
				}
			}
		}
		if flag {
			return errors.ErrNotEnoughAmount()
		}
	} else {
		if boxInfo.GetDeposit().InterestInjection != nil {
			totalInterest := sdk.ZeroInt()
			for _, v := range boxInfo.GetDeposit().InterestInjection {
				if v.Address.Equals(account.GetAddress()) {
					totalInterest = totalInterest.Add(v.Amount)
				}
			}
			if totalInterest.Add(amount).GT(boxInfo.GetDeposit().Interest.Amount) {
				return errors.Errorf(errors.ErrInterestInjectionNotValid(sdk.NewCoin(boxInfo.GetDeposit().Interest.Denom, amountArg)))
			}
		}
	}
	msg := msgs.NewMsgBoxInterest(boxID, account.GetAddress(), sdk.NewCoin(boxInfo.GetDeposit().Interest.Denom, amount), operation)

	validateErr := msg.ValidateBasic()
	if validateErr != nil {
		return errors.Errorf(validateErr)
	}
	return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
}

func MarkCmdDepositBoxCreateFlagRequired(cmd *cobra.Command) {
	cmd.MarkFlagRequired(flagBottomLine)
	cmd.MarkFlagRequired(flagPrice)
	cmd.MarkFlagRequired(flagInterest)
	cmd.MarkFlagRequired(flagStartTime)
	cmd.MarkFlagRequired(flagEstablishTime)
	cmd.MarkFlagRequired(flagMaturityTime)
}
