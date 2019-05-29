package cli

import (
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientutils "github.com/hashgard/hashgard/x/box/client/utils"
	"github.com/hashgard/hashgard/x/box/errors"
	boxutils "github.com/hashgard/hashgard/x/box/utils"
)

// ProcessBoxWithdrawCmd implements withdraw a box transaction command.
func ProcessBoxWithdrawCmd(cdc *codec.Codec, boxType string, id string) error {
	if boxutils.GetBoxTypeByValue(id) != boxType {
		return errors.Errorf(errors.ErrUnknownBox(id))
	}
	txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
	if err != nil {
		return err
	}
	msg, err := clientutils.GetWithdrawMsg(cdc, cliCtx, account, id)
	if err != nil {
		return err
	}
	return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)

}
