package cli

import (
	"fmt"

	"github.com/hashgard/hashgard/x/issue/errors"

	"github.com/hashgard/hashgard/x/issue/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/cosmos/cosmos-sdk/x/bank"
	boxutils "github.com/hashgard/hashgard/x/box/utils"
	issuequeriers "github.com/hashgard/hashgard/x/issue/client/queriers"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
	"github.com/spf13/cobra"
)

// SendTxCmd will create a send tx and sign it with the given key.
func SendTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [to_address] [amount]",
		Short: "Create and sign a send tx",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			to, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			// parse coins trying to be sent
			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}
			from := cliCtx.GetFromAddress()

			for i, coin := range coins {
				if boxutils.IsBoxId(coin.Denom) {
					return errors.Errorf(sdk.ErrInternal("box not support yet"))
				}
				if issueutils.IsIssueId(coin.Denom) {
					res, err := issuequeriers.QueryIssueByID(coin.Denom, cliCtx)
					if err == nil {
						var issueInfo types.Issue
						cdc.MustUnmarshalJSON(res, &issueInfo)
						coins[i].Amount = issueutils.MulDecimals(coin.Amount, issueInfo.GetDecimals())
						if err = issueutils.CheckFreeze(cdc, cliCtx, issueInfo.GetIssueId(), from, to); err != nil {
							return err
						}
					}
				}
			}

			account, err := cliCtx.GetAccount(from)
			if err != nil {
				return err
			}

			// ensure account has enough coins
			if !account.GetCoins().IsAllGTE(coins) {
				return fmt.Errorf("address %s doesn't have enough coins to pay for this transaction", from)
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := bank.NewMsgSend(from, to, coins)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	return client.PostCommands(cmd)[0]
}
