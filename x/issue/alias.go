package issue

import (
	"github.com/hashgard/hashgard/x/issue/client"
	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/types"
	"github.com/hashgard/hashgard/x/lock/client/cli"
)

type (
	Keeper        = keeper.Keeper
	CoinIssueInfo = types.CoinIssueInfo
)

var (
	MsgCdc          = msgs.MsgCdc
	NewKeeper       = keeper.NewKeeper
	NewModuleClient = client.NewModuleClient
	GetAccountCmd   = cli.GetAccountCmd
	RegisterCodec   = msgs.RegisterCodec
)

const (
	StoreKey          = types.StoreKey
	RouterKey         = types.RouterKey
	QuerierRoute      = types.QuerierRoute
	DefaultParamspace = types.DefaultParamspace
	DefaultCodespace  = types.DefaultCodespace
)
