package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"strings"

	issuequeriers "github.com/hashgard/hashgard/x/issue/client/queriers"
	"github.com/hashgard/hashgard/x/issue/domain"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
)

// GetAccountCmd returns a query account that will display the state of the
// account at a given address.
// nolint: unparam
func GetAccountCmd(storeName string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account [address]",
		Short: "Query account balance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).WithAccountDecoder(cdc)
			key, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			if err = cliCtx.EnsureAccountExistsFromAddr(key); err != nil {
				return err
			}
			acc, err := cliCtx.GetAccount(key)
			if err != nil {
				return err
			}
			if acc.GetCoins().Empty() {
				return cliCtx.PrintOutput(acc)
			}
			coins := make(sdk.Coins, acc.GetCoins().Len())
			i := 0
			for _, coin := range acc.GetCoins() {
				denom := coin.Denom
				if issueutils.IsIssueId(coin.Denom) {
					res, err := issuequeriers.QueryIssueByID(coin.Denom, cliCtx, cdc, domain.QuerierRoute)
					if err == nil {
						var issueInfo domain.Issue
						cdc.MustUnmarshalJSON(res, &issueInfo)
						denom = fmt.Sprintf("%s(%s)", issueInfo.GetName(), coin.Denom)
					}
				}
				newCoin := sdk.Coin{Denom: denom, Amount: coin.Amount}
				coins[i] = newCoin
				i += 1
			}
			acc.SetCoins(coins)
			return cliCtx.PrintOutput(acc)
		},
	}
	return client.GetCommands(cmd)[0]
}

// GetCmdQueryIssue implements the query issue command.
func GetCmdQueryIssue(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "query-issue [issue-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query details of a single issue",
		Long: strings.TrimSpace(`
Query details for a issue. You can find the issue-id by running hashgardcli query issue coins:

$ hashgardcli issue query-issue gardh1c7d59vebq
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			issueID := args[0]
			if error := issueutils.CheckIssueId(issueID); error != nil {
				return error
			}
			// Query the issue
			res, err := issuequeriers.QueryIssueByID(issueID, cliCtx, cdc, queryRoute)
			if err != nil {
				return err
			}
			var issueInfo domain.Issue
			cdc.MustUnmarshalJSON(res, &issueInfo)
			return cliCtx.PrintOutput(issueInfo)
		},
	}
}
