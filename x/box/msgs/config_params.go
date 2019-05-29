package msgs

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

var (
	// key for constant fee parameter
	ParamStoreKeyLockCreateFee        = []byte("LockCreateFee")
	ParamStoreKeyDepositBoxCreateFee  = []byte("DepositBoxCreateFee")
	ParamStoreKeyFutureBoxCreateFee   = []byte("FutureBoxCreateFee")
	ParamStoreKeyBoxDisableFeatureFee = []byte("BoxDisableFeatureFee")
	ParamStoreKeyBoxDescribeFee       = []byte("BoxDescribeFee")
)

var _ params.ParamSet = &BoxConfigParams{}

// Param Config issue for issue
type BoxConfigParams struct {
	LockCreateFee       sdk.Coin `json:"lock_create_fee"`
	DepositBoxCreateFee sdk.Coin `json:"deposit_box_create_fee"`
	FutureBoxCreateFee  sdk.Coin `json:"future_box_create_fee"`
	DisableFeatureFee   sdk.Coin `json:"disable_feature_fee"`
	DescribeFee         sdk.Coin `json:"describe_fee"`
}

// ParamKeyTable for auth module
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&BoxConfigParams{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of auth module's parameters.
// nolint
func (p *BoxConfigParams) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{ParamStoreKeyLockCreateFee, &p.LockCreateFee},
		{ParamStoreKeyDepositBoxCreateFee, &p.DepositBoxCreateFee},
		{ParamStoreKeyFutureBoxCreateFee, &p.FutureBoxCreateFee},
		{ParamStoreKeyBoxDisableFeatureFee, &p.DisableFeatureFee},
		{ParamStoreKeyBoxDescribeFee, &p.DescribeFee},
	}
}

// Checks equality of BoxConfigParams
func (dp BoxConfigParams) Equal(dp2 BoxConfigParams) bool {
	b1 := MsgCdc.MustMarshalBinaryBare(dp)
	b2 := MsgCdc.MustMarshalBinaryBare(dp2)
	return bytes.Equal(b1, b2)
}

// DefaultParams returns a default set of parameters.
func DefaultParams(denom string) BoxConfigParams {
	return BoxConfigParams{
		LockCreateFee:       sdk.NewCoin(denom, sdk.NewIntWithDecimal(100, 18)),
		DepositBoxCreateFee: sdk.NewCoin(denom, sdk.NewIntWithDecimal(1000, 18)),
		FutureBoxCreateFee:  sdk.NewCoin(denom, sdk.NewIntWithDecimal(1000, 18)),
		DisableFeatureFee:   sdk.NewCoin(denom, sdk.NewIntWithDecimal(1000, 18)),
		DescribeFee:         sdk.NewCoin(denom, sdk.NewIntWithDecimal(100, 18)),
	}
}

func (dp BoxConfigParams) String() string {
	return fmt.Sprintf(`BoxConfigParams:
  LockCreateFee:			%s
  DepositBoxCreateFee:			%s
  FutureBoxCreateFee:			%s
  DisableFeatureFee:			%s
  DescribeFee:			%s`,
		dp.LockCreateFee.String(),
		dp.DepositBoxCreateFee.String(),
		dp.FutureBoxCreateFee.String(),
		dp.DisableFeatureFee.String(),
		dp.DescribeFee.String())
}
