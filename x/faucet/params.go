package faucet

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

const (
	DefaultParamspace = ModuleName
)

var (
	ParamStoreKeyFaucetOrigin = []byte("faucetorigin")
	ParamStoreKeySendCoins = []byte("sendcoins")
	ParamStoreKeyLimitCoins = []byte("limitcoins")
)

// ParamKeyTable type declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamStoreKeyFaucetOrigin, sdk.AccAddress{},
		ParamStoreKeySendCoins, sdk.Coins{},
		ParamStoreKeyLimitCoins, sdk.Coins{},
	)
}
