package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"strings"

	issueutils "github.com/hashgard/hashgard/x/issue/client/utils"

	"github.com/hashgard/hashgard/x/issue"
)

const (
	flagName        = "name"
	flagTotalSupply = "totalSupply"
	flagIssue       = "issue"
)

var issueFlags = []string{
	flagName,
	flagTotalSupply,
}

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

			for _, coin := range acc.GetCoins() {
				if len(coin.Denom) == 15 && strings.HasPrefix(coin.Denom, issue.IDPreStr) {
					res, err := issueutils.QueryIssueByID(coin.Denom, cliCtx, cdc, issue.QuerierRoute)
					if err != nil {
						var coinIssueInfo issue.CoinIssueInfo
						if err := cdc.UnmarshalBinaryLengthPrefixed(res, &coinIssueInfo); err != nil {
							//TODO 需重建结构体
							coin.Denom = coinIssueInfo.Name
						}
					}
				}
			}
			return cliCtx.PrintOutput(acc)
		},
	}
	return client.GetCommands(cmd)[0]
}

//TODO
// GetCmdIssue implements issue a coin transaction command.
func GetCmdIssue(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "issue a coin",
		Long: strings.TrimSpace(`
issue a coin along . For example:

$ hashgardcli issue mycoin 200000
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			coinIssueInfo, err := parseIssueFlags()
			if err != nil {
				return err
			}

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)

			// Get from address
			from := cliCtx.GetFromAddress()

			//Pull associated account
			//account, err := cliCtx.GetAccount(from)
			//if err != nil {
			//	return err
			//}

			msg := issue.NewMsgIssue(from, coinIssueInfo)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd.Flags().String(flagName, "", "name of coin")
	cmd.Flags().String(flagTotalSupply, "", "total supply of coin")
	return cmd
}
