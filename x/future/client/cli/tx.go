package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	boxcli "github.com/hashgard/hashgard/x/box/client/cli"
	"github.com/hashgard/hashgard/x/box/types"
	"github.com/spf13/cobra"
)

func GetDepositToCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deposit-to [id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Deposit to the future box",
		Long:    "Deposit to the future box",
		Example: "$ hashgardcli future deposit-to box174876e800 88888 --from foo",
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
		Short:   "Fetch deposit from a future box",
		Long:    "Fetch deposit from a future box",
		Example: "$ hashgardcli future fetch box174876e800 88888 --from foo",
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
		Short:   "Holder withdraw from a future box",
		Long:    "Holder withdraw from a future box when the future box can be withdraw",
		Example: "$ hashgardcli future withdraw boxab3jlxpt2ps --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return boxcli.ProcessBoxWithdrawCmd(cdc, types.Future, args[0])
		},
	}
	return cmd
}

func GetDescriptionCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "describe [id] [description-file]",
		Args:    cobra.ExactArgs(2),
		Short:   "Describe a future box",
		Long:    "Box owner can set description of the future box, and the description need to be in json format. You can customize preferences or use recommended templates.",
		Example: "$ hashgardcli future describe boxab3jlxpt2ps path/description.json --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return boxcli.ProcessBoxDescriptionCmd(cdc, types.Future, args[0], args[1])
		},
	}
	return cmd
}

// GetDisableFeatureCmd implements disable feature a box transaction command.
func GetDisableFeatureCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disable [id] [feature]",
		Args:  cobra.ExactArgs(2),
		Short: "Disable feature from a future box",
		Long: fmt.Sprintf("Box Owner disabled the features:\n"+
			"%s:Box holder can transfer", types.Transfer),
		Example: fmt.Sprintf("$ hashgardcli future disable boxab3jlxpt2ps %s --from foo", types.Transfer),
		RunE: func(cmd *cobra.Command, args []string) error {
			return boxcli.ProcessBoxDisableFeatureCmd(cdc, types.Future, args[0], args[1])
		},
	}
	return cmd
}
