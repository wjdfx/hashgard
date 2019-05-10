package cli

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	boxqueriers "github.com/hashgard/hashgard/x/box/client/queriers"
	"github.com/hashgard/hashgard/x/box/client/utils"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/params"
	"github.com/hashgard/hashgard/x/box/types"
	boxutils "github.com/hashgard/hashgard/x/box/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetCmdQueryBox implements the query box command.
func GetCmdQueryBox(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "query-box [box-id]",
		Args:    cobra.ExactArgs(1),
		Short:   "Query a single box",
		Long:    "Query details for a box. You can find the box-id by running hashgardcli box list-box",
		Example: "$ hashgardcli box query-box boxab3jlxpt2ps",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			boxID := args[0]
			if err := boxutils.CheckBoxId(boxID); err != nil {
				return errors.Errorf(err)
			}
			// Query the box
			res, err := boxqueriers.QueryBoxByID(boxID, cliCtx)
			if err != nil {
				return err
			}
			var box types.BoxInfo
			cdc.MustUnmarshalJSON(res, &box)

			return cliCtx.PrintOutput(utils.GetBoxInfo(cdc, cliCtx, box))
		},
	}
}

// GetCmdQueryDepositBoxDeposit implements the query box command.
func GetCmdQueryDepositBoxDeposit(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "query-deposit [box-id]",
		Args:    cobra.ExactArgs(1),
		Short:   "Query deposit list from deposit box",
		Long:    "Query deposit list from deposit box",
		Example: "$ hashgardcli box query-deposit boxab3jlxpt2ps",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			boxID := args[0]
			if err := boxutils.CheckBoxId(boxID); err != nil {
				return errors.Errorf(err)
			}
			address, err := sdk.AccAddressFromBech32(viper.GetString(flagAddress))
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
			if boxInfo.GetDeposit().Status == types.BoxCreated {
				return nil
			}

			boxQueryParams := params.BoxQueryDepositListParams{
				BoxId: boxID,
				Owner: address,
			}
			// Query the box
			res, err := boxqueriers.QueryDepositList(boxQueryParams, cdc, cliCtx)
			if err != nil {
				return err
			}

			var boxs types.DepositBoxDepositToList
			cdc.MustUnmarshalJSON(res, &boxs)
			for i, box := range boxs {
				if box.Amount.IsZero() {
					continue
				}
				boxs[i].Amount = boxutils.GetBoxCoinByDecimal(cdc, cliCtx, sdk.NewCoin(boxInfo.GetTotalAmount().Denom, box.Amount)).Amount
			}
			return cliCtx.PrintOutput(boxs)
		},
	}
	cmd.Flags().String(flagAddress, "", "Box owner address")
	//cmd.Flags().Int32(flagLimit, 30, "Query number of box results per page returned")
	return cmd
}

// GetCmdQueryBox implements the query box command.
func GetCmdQueryBoxs(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list-box [box-type]",
		Args:    cobra.ExactArgs(1),
		Short:   "Query box list",
		Long:    "Query all or one of the account box list, the limit default is 30",
		Example: "$ hashgardcli box list-box lock",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			_, ok := types.BoxType[args[0]]
			if !ok {
				return errors.Errorf(errors.ErrUnknownBoxType())
			}
			address, err := sdk.AccAddressFromBech32(viper.GetString(flagAddress))
			if err != nil {
				return err
			}
			boxQueryParams := params.BoxQueryParams{
				StartBoxId: viper.GetString(flagStartBoxId),
				BoxType:    args[0],
				Owner:      address,
				Limit:      viper.GetInt(flagLimit),
			}
			// Query the box
			res, err := boxqueriers.QueryBoxsList(boxQueryParams, cdc, cliCtx)
			if err != nil {
				return err
			}

			var boxs types.BoxInfos
			cdc.MustUnmarshalJSON(res, &boxs)
			return cliCtx.PrintOutput(utils.GetBoxList(cdc, cliCtx, boxs, boxQueryParams.BoxType))
		},
	}

	cmd.Flags().String(flagAddress, "", "Box owner address")
	cmd.Flags().String(flagStartBoxId, "", "Start boxId of box results")
	cmd.Flags().Int32(flagLimit, 30, "Query number of box results per page returned")

	return cmd
}

// GetCmdQueryBoxs implements the query box command.
func GetCmdSearchBoxs(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "search [box-type] [name]",
		Args:    cobra.ExactArgs(2),
		Short:   "Search boxs",
		Long:    "Search boxs based on name",
		Example: "$ hashgardcli box search fo",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, ok := types.BoxType[args[0]]
			if !ok {
				return errors.Errorf(errors.ErrUnknownBoxType())
			}
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query the box
			res, err := boxqueriers.QueryBoxByName(args[0], strings.ToLower(args[1]), cliCtx)
			if err != nil {
				return err
			}
			var boxs types.BoxInfos
			cdc.MustUnmarshalJSON(res, &boxs)
			for i, box := range boxs {
				boxs[i].TotalAmount = boxutils.GetBoxCoinByDecimal(cdc, cliCtx, box.TotalAmount)
			}
			return cliCtx.PrintOutput(utils.GetBoxList(cdc, cliCtx, boxs, args[0]))
		},
	}
	return cmd
}
