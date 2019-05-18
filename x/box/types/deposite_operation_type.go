package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var DepositOperation = map[string]uint{DepositTo: 1, Fetch: 2}

func CheckDepositOperation(depositOperation string) sdk.Error {
	_, ok := DepositOperation[depositOperation]
	if !ok {
		return sdk.ErrInternal("unknown deposit operation:" + depositOperation)
	}
	return nil
}
