package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	boxcli "github.com/hashgard/hashgard/x/box/client/cli"
	clientutils "github.com/hashgard/hashgard/x/box/client/utils"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/types"
	boxutils "github.com/hashgard/hashgard/x/box/utils"
	"github.com/spf13/cobra"
)

// GetInterestInjectionCmd implements interest injection a deposit box transaction command.
func GetInterestInjectionCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "interest-injection [id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Injection interest to the deposit box",
		Long:    "Injection interest to the deposit box",
		Example: "$ hashgardcli deposit interest-injection box174876e800 88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return interest(cdc, args[0], args[1], types.Injection)
		},
	}
	return cmd
}

// GetInterestFetchCmd implements fetch interest from a deposit box transaction command.
func GetInterestFetchCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "interest-fetch [id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Fetch interest from a deposit box",
		Long:    "Fetch interest from a deposit box",
		Example: "$ hashgardcli deposit interest-fetch box174876e800 88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return interest(cdc, args[0], args[1], types.Fetch)
		},
	}
	return cmd
}

func interest(cdc *codec.Codec, id string, amountStr string, operation string) error {
	if boxutils.GetBoxTypeByValue(id) != types.Deposit {
		return errors.Errorf(errors.ErrNotSupportOperation())
	}
	txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
	if err != nil {
		return err
	}
	msg, err := clientutils.GetInterestMsg(cdc, cliCtx, account, id, amountStr, operation, true)
	if err != nil {
		return err
	}

	return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
}

func GetDepositToCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deposit-to [id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Deposit to the deposit box",
		Long:    "Deposit to the deposit box",
		Example: "$ hashgardcli deposit deposit-to box174876e800 88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return boxcli.ProcessBoxDeposit(cdc, args[0], args[1], types.DepositTo)
		},
	}
	return cmd
}

func GetFetchDepositCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "fetch [id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Fetch deposit from a deposit box",
		Long:    "Fetch deposit from a deposit box",
		Example: "$ hashgardcli deposit fetch box174876e800 88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return boxcli.ProcessBoxDeposit(cdc, args[0], args[1], types.Fetch)
		},
	}
	return cmd
}

func GetWithdrawCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "withdraw [id]",
		Args:    cobra.ExactArgs(1),
		Short:   "Holder withdraw from a deposit box",
		Long:    "Holder withdraw from a deposit box when the deposit box can be withdraw",
		Example: "$ hashgardcli deposit withdraw boxab3jlxpt2ps --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return boxcli.ProcessBoxWithdrawCmd(cdc, types.Deposit, args[0])
		},
	}
	return cmd
}

func GetDescriptionCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "describe [id] [description-file]",
		Args:    cobra.ExactArgs(2),
		Short:   "Describe a deposit box",
		Long:    "Box owner can set description of the deposit box, and the description need to be in json format. You can customize preferences or use recommended templates.",
		Example: "$ hashgardcli deposit describe boxab3jlxpt2ps path/description.json --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return boxcli.ProcessBoxDescriptionCmd(cdc, types.Deposit, args[0], args[1])
		},
	}
	return cmd
}

// GetDisableFeatureCmd implements disable feature a box transaction command.
func GetDisableFeatureCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disable [id] [feature]",
		Args:  cobra.ExactArgs(2),
		Short: "Disable feature from a deposit box",
		Long: fmt.Sprintf("Box Owner disabled the features:\n"+
			"%s:Box holder can transfer", types.Transfer),
		Example: fmt.Sprintf("$ hashgardcli deposit disable boxab3jlxpt2ps %s --from foo", types.Transfer),
		RunE: func(cmd *cobra.Command, args []string) error {
			return boxcli.ProcessBoxDisableFeatureCmd(cdc, types.Deposit, args[0], args[1])
		},
	}
	return cmd
}
