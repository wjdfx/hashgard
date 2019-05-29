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
	"github.com/spf13/viper"
)

// ProcessQueryBoxCmd implements the query box command.
func ProcessQueryBoxCmd(cdc *codec.Codec, boxType string, id string) error {
	if boxutils.GetBoxTypeByValue(id) != boxType {
		return errors.Errorf(errors.ErrUnknownBox(id))
	}
	cliCtx := context.NewCLIContext().WithCodec(cdc)
	if err := boxutils.CheckId(id); err != nil {
		return errors.Errorf(err)
	}
	// Query the box
	res, err := boxqueriers.QueryBoxByID(id, cliCtx)
	if err != nil {
		return err
	}
	var box types.BoxInfo
	cdc.MustUnmarshalJSON(res, &box)
	return cliCtx.PrintOutput(utils.GetBoxInfo(box))
}

// ProcessListBoxCmd implements the query box command.
func ProcessListBoxCmd(cdc *codec.Codec, boxType string) error {
	cliCtx := context.NewCLIContext().WithCodec(cdc)
	_, ok := types.BoxType[boxType]
	if !ok {
		return errors.Errorf(errors.ErrUnknownBoxType())
	}
	address, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
	if err != nil {
		return err
	}
	boxQueryParams := params.BoxQueryParams{
		StartId: viper.GetString(FlagStartId),
		BoxType: boxType,
		Owner:   address,
		Limit:   viper.GetInt(FlagLimit),
	}
	// Query the box
	res, err := boxqueriers.QueryBoxsList(boxQueryParams, cdc, cliCtx)
	if err != nil {
		return err
	}
	var boxs types.BoxInfos
	cdc.MustUnmarshalJSON(res, &boxs)
	return cliCtx.PrintOutput(utils.GetBoxList(boxs, boxQueryParams.BoxType))
}

// ProcessSearchBoxsCmd implements the query box command.
func ProcessSearchBoxsCmd(cdc *codec.Codec, boxType string, name string) error {
	_, ok := types.BoxType[boxType]

	if !ok {
		return errors.Errorf(errors.ErrUnknownBoxType())
	}
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	// Query the box
	res, err := boxqueriers.QueryBoxByName(boxType, strings.ToLower(name), cliCtx)
	if err != nil {
		return err
	}
	var boxs types.BoxInfos
	cdc.MustUnmarshalJSON(res, &boxs)
	return cliCtx.PrintOutput(utils.GetBoxList(boxs, boxType))
}
