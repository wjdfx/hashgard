package cli

import (
	"fmt"
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
		Use:     "query-issue [issue-id]",
		Args:    cobra.ExactArgs(1),
		Short:   "Query a single issue",
		Long:    "Query details for a issue. You can find the issue-id by running hashgardcli issue list-issues",
		Example: "$ hashgardcli issue query-issue gardh1c7d59vebq",
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

// GetCmdQueryAllowance implements the query allowance command.
func GetCmdQueryAllowance(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "query-allowance [issue-id] [owner-address] [spender-address]",
		Args:    cobra.ExactArgs(3),
		Short:   "Query allowance",
		Long:    "Query the amount of tokens that an owner allowed to a spender",
		Example: "$ hashgardcli issue query-allowance coin174876e800 gard1zu85q8a7wev675k527y7keyrea7wu7crr9vdrs gard1vud9ptwagudgq7yht53cwuf8qfmgkd0qcej0ah",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			ownerAddress, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}
			spenderAddress, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			res, err := issuequeriers.QueryIssueAllowance(issueID, ownerAddress, spenderAddress, cliCtx)
			if err != nil {
				return err
			}
			var approval types.Approval
			cdc.MustUnmarshalJSON(res, &approval)

			return cliCtx.PrintOutput(approval)
		},
	}
}

// GetCmdQueryIssues implements the query issue command.
func GetCmdQueryIssues(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list-issues",
		Short:   "Query issue list",
		Long:    "Query all or one of the account issue list, the limit default is 30",
		Example: "$ hashgardcli issue list-issues gard10cm9l6ly924d37qksn2x93xt3ezhduc2ntdj04",
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

			var tokenIssues types.CoinIssues
			cdc.MustUnmarshalJSON(res, &tokenIssues)
			if len(tokenIssues) == 0 {
				fmt.Println("No records")
				return nil
			}
			for i, token := range tokenIssues {
				tokenIssues[i].TotalSupply = issueutils.QuoDecimals(token.TotalSupply, token.Decimals)
			}
			return cliCtx.PrintOutput(tokenIssues)
		},
	}

	cmd.Flags().String(flagAddress, "", "Token owner address")
	cmd.Flags().String(flagSymbol, "", "Symbol of issue token")
	cmd.Flags().String(flagStartIssueId, "", "Start issueId of issues")
	cmd.Flags().Int32(flagLimit, 30, "Query number of issue results per page returned")

	return cmd
}

// GetCmdQueryIssues implements the query issue command.
func GetCmdSearchIssues(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "search [symbol]",
		Args:    cobra.ExactArgs(1),
		Short:   "Search issues",
		Long:    "Search issues based on symbol",
		Example: "$ hashgardcli issue search fo",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query the issue
			res, err := issuequeriers.QueryIssueBySymbol(strings.ToUpper(args[0]), cliCtx)
			if err != nil {
				return err
			}
			var tokenIssues types.CoinIssues
			cdc.MustUnmarshalJSON(res, &tokenIssues)
			for i, token := range tokenIssues {
				tokenIssues[i].TotalSupply = issueutils.QuoDecimals(token.TotalSupply, token.Decimals)
			}
			return cliCtx.PrintOutput(tokenIssues)
		},
	}
	return cmd
}
