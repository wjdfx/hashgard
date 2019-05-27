package box

import (
	"github.com/hashgard/hashgard/x/box/client"
	"github.com/hashgard/hashgard/x/box/keeper"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/types"
)

type (
	Keeper  = keeper.Keeper
	BoxInfo = types.BoxInfo
	Hooks	= keeper.Hooks
)

var (
	MsgCdc          = msgs.MsgCdc
	NewKeeper       = keeper.NewKeeper
	NewModuleClient = client.NewModuleClient
	RegisterCodec   = msgs.RegisterCodec
)

const (
	StoreKey          = types.StoreKey
	RouterKey         = types.RouterKey
	QuerierRoute      = types.QuerierRoute
	DefaultParamspace = types.DefaultParamspace
	DefaultCodespace  = types.DefaultCodespace
)
