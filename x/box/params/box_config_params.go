package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

var (
	// key for constant fee parameter
	ParamStoreKeyLockBoxCreateFee     = []byte("LockBoxCreateFee")
	ParamStoreKeyDepositBoxCreateFee  = []byte("DepositBoxCreateFee")
	ParamStoreKeyFutureBoxCreateFee   = []byte("FutureBoxCreateFee")
	ParamStoreKeyBoxEnableTransferFee = []byte("BoxEnableTransferFee")
)

// type declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamStoreKeyLockBoxCreateFee, sdk.Coin{},
		ParamStoreKeyDepositBoxCreateFee, sdk.Coin{},
		ParamStoreKeyFutureBoxCreateFee, sdk.Coin{},
		ParamStoreKeyBoxEnableTransferFee, sdk.Coin{},
	)
}
