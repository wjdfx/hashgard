package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/hashgard/hashgard/x/exchange/queriers"
	"github.com/hashgard/hashgard/x/exchange/types"
)

func GetCmdQueryOrder(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "query-order [order-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query detail info of the order with the specific order-id",
		Long: strings.TrimSpace(`
$ hashgardcli exchange query-order 1
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// validate that the proposal id is a uint
			orderId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("order-id %s not a valid uint, please input a valid order-id", args[0])
			}

			params := queriers.NewQueryOrderParams(orderId)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/order", queryRoute), bz)
			if err != nil {
				return err
			}

			var order types.Order
			cdc.MustUnmarshalJSON(res, &order)
			return cliCtx.PrintOutput(order)
		},
	}
}

func GetCmdQueryOrdersByAddr(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "query-orders [address]",
		Args:  cobra.ExactArgs(1),
		Short: "Query all orders of a specific address",
		Long: strings.TrimSpace(`
$ hashgardcli exchange query-orders gard1hf4n743fujvxrwx8af7u35anjqpdd2cx8p6cdd
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			seller, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			params := queriers.NewQueryOrdersParams(seller)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/orders", queryRoute), bz)
			if err != nil {
				return err
			}

			var orders types.Orders
			cdc.MustUnmarshalJSON(res, &orders)
			return cliCtx.PrintOutput(orders)
		},
	}
}

func GetCmdFrozenFund(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "query-frozen [address]",
		Args:  cobra.ExactArgs(1),
		Short: "Query frozen fund of a specific address",
		Long: strings.TrimSpace(`
$ hashgardcli exchange query-frozen gard1hf4n743fujvxrwx8af7u35anjqpdd2cx8p6cdd
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			seller, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			params := queriers.NewQueryFrozenFundParams(seller)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/frozen", queryRoute), bz)
			if err != nil {
				return err
			}

			var coins sdk.Coins
			cdc.MustUnmarshalJSON(res, &coins)
			return cliCtx.PrintOutput(coins)
		},
	}
}
