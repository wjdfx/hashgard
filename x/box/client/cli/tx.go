package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/hashgard/hashgard/x/box/types"

	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientutils "github.com/hashgard/hashgard/x/box/client/utils"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/msgs"
	boxutils "github.com/hashgard/hashgard/x/box/utils"
	"github.com/spf13/cobra"
)

// GetCmdBoxDescription implements box a coin transaction command.
func GetCmdBoxDescription(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "describe [box-id] [description-file]",
		Args:    cobra.ExactArgs(2),
		Short:   "Describe a box",
		Long:    "Owner can add description of the box by owner, and the description need to be in json format. You can customize preferences or use recommended templates.",
		Example: "$ hashgardcli box describe boxab3jlxpt2ps path/description.json --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			boxID := args[0]
			if err := boxutils.CheckBoxId(boxID); err != nil {
				return errors.Errorf(err)
			}
			txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
			if err != nil {
				return err
			}
			contents, err := ioutil.ReadFile(args[1])
			if err != nil {
				return err
			}
			buffer := bytes.Buffer{}
			err = json.Compact(&buffer, contents)
			if err != nil {
				return errors.ErrBoxDescriptionNotValid()
			}
			contents = buffer.Bytes()

			_, err = boxutils.BoxOwnerCheck(cdc, cliCtx, account, boxID)
			if err != nil {
				return err
			}
			msg := msgs.NewMsgBoxDescription(boxID, account.GetAddress(), contents)

			validateErr := msg.ValidateBasic()

			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	return cmd
}

// GetCmdBoxDisableFeature implements disable feature a box transaction command.
func GetCmdBoxDisableFeature(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disable [box-id] [feature]",
		Args:  cobra.ExactArgs(2),
		Short: "Disable feature from a box",
		Long: fmt.Sprintf("Box Owner disabled the features:\n"+
			"%s:Box holder can trade the box", types.Trade),
		Example: fmt.Sprintf("$ hashgardcli box disable boxab3jlxpt2ps %s --from foo", types.Trade),
		RunE: func(cmd *cobra.Command, args []string) error {
			feature := args[1]

			_, ok := types.Features[feature]
			if !ok {
				return errors.Errorf(errors.ErrUnknownFeatures())
			}

			boxID := args[0]
			if err := boxutils.CheckBoxId(boxID); err != nil {
				return errors.Errorf(err)
			}
			txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
			if err != nil {
				return err
			}
			boxInfo, err := boxutils.BoxOwnerCheck(cdc, cliCtx, account, boxID)
			if err != nil {
				return err
			}
			if feature == types.Trade && boxInfo.GetBoxType() == types.Lock {
				return errors.Errorf(errors.ErrNotSupportOperation())
			}

			msg := msgs.NewMsgBoxDisableFeature(boxID, account.GetAddress(), feature)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	return cmd
}
