package utils

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/hashgard/hashgard/x/issue/msgs"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/types"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
)

func GetCliContext(cdc *codec.Codec) (authtxb.TxBuilder, context.CLIContext, auth.Account, error) {
	txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
	cliCtx := context.NewCLIContext().
		WithCodec(cdc).
		WithAccountDecoder(cdc)
	from := cliCtx.GetFromAddress()
	account, err := cliCtx.GetAccount(from)

	return txBldr, cliCtx, account, err
}
func burnCheck(sender auth.Account, burnFrom sdk.AccAddress, issueInfo types.Issue, amount sdk.Int, burnType string, cli bool) error {

	coins := sender.GetCoins()

	switch burnType {
	case types.BurnOwner:
		{
			if !sender.GetAddress().Equals(issueInfo.GetOwner()) {
				return errors.Errorf(errors.ErrOwnerMismatch(issueInfo.GetIssueId()))
			}
			if !sender.GetAddress().Equals(burnFrom) {
				return errors.Errorf(errors.ErrOwnerMismatch(issueInfo.GetIssueId()))
			}
			if issueInfo.IsBurnOwnerDisabled() {
				return errors.Errorf(errors.ErrCanNotBurn(issueInfo.GetIssueId(), burnType))
			}
			break
		}
	case types.BurnHolder:
		{
			if issueInfo.IsBurnHolderDisabled() {
				return errors.Errorf(errors.ErrCanNotBurn(issueInfo.GetIssueId(), burnType))
			}
			if !sender.GetAddress().Equals(burnFrom) {
				return errors.Errorf(errors.ErrOwnerMismatch(issueInfo.GetIssueId()))
			}
			break
		}
	case types.BurnFrom:
		{
			if !sender.GetAddress().Equals(issueInfo.GetOwner()) {
				return errors.Errorf(errors.ErrOwnerMismatch(issueInfo.GetIssueId()))
			}
			if issueInfo.IsBurnFromDisabled() {
				return errors.Errorf(errors.ErrCanNotBurn(issueInfo.GetIssueId(), burnType))
			}
			if issueInfo.GetOwner().Equals(burnFrom) {
				//burnFrom
				if issueInfo.IsBurnOwnerDisabled() {
					return errors.Errorf(errors.ErrCanNotBurn(issueInfo.GetIssueId(), types.BurnOwner))
				}
			}
			break
		}
	default:
		{
			panic("not support")
		}

	}
	if cli {
		amount = issueutils.MulDecimals(amount, issueInfo.GetDecimals())
	}
	// ensure account has enough coins
	if !coins.IsAllGTE(sdk.NewCoins(sdk.NewCoin(issueInfo.GetIssueId(), amount))) {
		return fmt.Errorf("address %s doesn't have enough coins to pay for this transaction", sender.GetAddress())
	}
	return nil
}
func GetBurnMsg(cdc *codec.Codec, cliCtx context.CLIContext, sender auth.Account,
	burnFrom sdk.AccAddress, issueID string, amount sdk.Int, burnFromType string, cli bool) (sdk.Msg, error) {

	issueInfo, err := issueutils.GetIssueByID(cdc, cliCtx, issueID)
	if err != nil {
		return nil, err
	}

	if types.BurnHolder == burnFromType {
		if issueInfo.GetOwner().Equals(sender.GetAddress()) {
			burnFromType = types.BurnOwner
		}
	}
	err = burnCheck(sender, burnFrom, issueInfo, amount, burnFromType, cli)
	if err != nil {
		return nil, err
	}
	if cli {
		amount = issueutils.MulDecimals(amount, issueInfo.GetDecimals())
	}

	var msg sdk.Msg

	switch burnFromType {

	case types.BurnOwner:
		msg = msgs.NewMsgIssueBurnOwner(issueID, sender.GetAddress(), amount)
		break
	case types.BurnHolder:
		msg = msgs.NewMsgIssueBurnHolder(issueID, sender.GetAddress(), amount)
		break
	case types.BurnFrom:
		msg = msgs.NewMsgIssueBurnFrom(issueID, sender.GetAddress(), burnFrom, amount)
		break
	default:
		return nil, errors.ErrCanNotBurn(issueID, burnFromType)
	}

	validateErr := msg.ValidateBasic()
	if validateErr != nil {
		return nil, errors.Errorf(validateErr)
	}

	return msg, nil
}
