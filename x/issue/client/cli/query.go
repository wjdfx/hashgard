package cli

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	issuequeriers "github.com/hashgard/hashgard/x/issue/client/queriers"
	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/params"
	"github.com/hashgard/hashgard/x/issue/types"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetCmdQueryIssue implements the query issue command.
func GetCmdQueryIssue(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "query [issue-id]",
		Args:    cobra.ExactArgs(1),
		Short:   "Query details of a single issue",
		Long:    "Query details for a issue. You can find the issue-id by running hashgardcli query issue coins",
		Example: "$ hashgardcli issue query gardh1c7d59vebq",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			// Query the issue
			res, err := issuequeriers.QueryIssueByID(issueID, cliCtx)
			if err != nil {
				return err
			}
			var issueInfo types.Issue
			cdc.MustUnmarshalJSON(res, &issueInfo)
			issueInfo.SetTotalSupply(issueutils.QuoDecimals(issueInfo.GetTotalSupply(), issueInfo.GetDecimals()))
			return cliCtx.PrintOutput(issueInfo)
		},
	}
}

// GetCmdQueryIssues implements the query issue command.
func GetCmdQueryIssues(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Query details of a address issues",
		Long:    "Query details for a address issues. You can find the address by running hashgardcli query issue coins",
		Example: "$ hashgardcli issue list gard10cm9l6ly924d37qksn2x93xt3ezhduc2ntdj04",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			address, err := sdk.AccAddressFromBech32(viper.GetString(flagAddress))
			if err != nil {
				return err
			}
			issueQueryParams := params.IssueQueryParams{
				StartIssueId: viper.GetString(flagStartIssueId),
				Owner:        address,
				Limit:        viper.GetInt(flagLimit),
			}
			// Query the issue
			res, err := issuequeriers.QueryIssuesList(issueQueryParams, cdc, cliCtx)
			if err != nil {
				return err
			}
			var coinIssues types.CoinIssues
			cdc.MustUnmarshalJSON(res, &coinIssues)
			for i, coin := range coinIssues {
				coinIssues[i].TotalSupply = issueutils.QuoDecimals(coin.TotalSupply, coin.Decimals)
			}
			return cliCtx.PrintOutput(coinIssues)
		},
	}

	cmd.Flags().String(flagAddress, "", "A account address of issue coin")
	cmd.Flags().String(flagSymbol, "", "Symbol of issue coin")
	cmd.Flags().String(flagStartIssueId, "", "Start issueId of issue coins")
	cmd.Flags().Int32(flagLimit, 30, "Query number of issue results per page returned")

	return cmd
}

// GetCmdQueryIssues implements the query issue command.
func GetCmdSearchIssues(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "search [symbol]",
		Args:    cobra.ExactArgs(1),
		Short:   "Query details of a address issues",
		Long:    "Query details for a address issues. You can find the address by running hashgardcli query issue coins",
		Example: "$ hashgardcli issue list gard10cm9l6ly924d37qksn2x93xt3ezhduc2ntdj04",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query the issue
			res, err := issuequeriers.QueryIssueBySymbol(strings.ToUpper(args[0]), cliCtx)
			if err != nil {
				return err
			}
			var coinIssues types.CoinIssues
			cdc.MustUnmarshalJSON(res, &coinIssues)
			for i, coin := range coinIssues {
				coinIssues[i].TotalSupply = issueutils.QuoDecimals(coin.TotalSupply, coin.Decimals)
			}
			return cliCtx.PrintOutput(coinIssues)
		},
	}
	return cmd
}
