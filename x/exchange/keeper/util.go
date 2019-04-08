package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetGratestDivisor(a sdk.Int, b sdk.Int) sdk.Int {
	for c := sdk.NewInt(0); !b.IsZero(); {
		c = a.Mod(b)
		a = b
		b = c
	}
	return a
}
