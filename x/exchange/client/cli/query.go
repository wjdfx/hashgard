package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/exchange"
)

func GetCmdQueryOrder(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "query-order [order-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query details of a order",
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

			params := exchange.NewQueryOrderParams(orderId)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/order", queryRoute), bz)
			if err != nil {
				return err
			}

			var order exchange.Order
			cdc.MustUnmarshalJSON(res, &order)
			return cliCtx.PrintOutput(order)
		},
	}
}

func GetCmdQueryOrdersByAddr(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "query-orders [address]",
		Args:  cobra.ExactArgs(1),
		Short: "Query orders of a address",
		Long: strings.TrimSpace(`
$ hashgardcli exchange query-orders address
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			seller, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			params := exchange.NewQueryOrdersParams(seller)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/orders", queryRoute), bz)
			if err != nil {
				return err
			}

			var orders exchange.Orders
			cdc.MustUnmarshalJSON(res, &orders)
			return cliCtx.PrintOutput(orders)
		},
	}
}

func GetCmdFrozenFund(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "query-frozen [address]",
		Args:  cobra.ExactArgs(1),
		Short: "Query frozen fund of an address",
		Long: strings.TrimSpace(`
$ hashgardcli exchange query-frozen address
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			seller, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			params := exchange.NewQueryFrozenFundParams(seller)
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