package utils

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
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
