package config

import (
	"bytes"
	"fmt"

	"github.com/hashgard/hashgard/x/issue/msgs"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

var (
	// key for constant fee parameter
	ParamStoreKeyIssueFee               = []byte("IssueFee")
	ParamStoreKeyMintFee                = []byte("IssueMintFee")
	ParamStoreKeyFreezeFee              = []byte("IssueFreezeFee")
	ParamStoreKeyUnFreezeFee            = []byte("IssueUnFreezeFee")
	ParamStoreKeyBurnFee                = []byte("IssueBurnFee")
	ParamStoreKeyBurnFromFee            = []byte("IssueBurnFromFee")
	ParamStoreKeyTransferOwnerFee       = []byte("IssueTransferOwnerFee")
	ParamStoreKeyIssueDisableFeatureFee = []byte("IssueDisableFeatureFee")
	ParamStoreKeyDescribeFee            = []byte("IssueDescribeFee")
)

var _ params.ParamSet = &Params{}

// Param Config issue for issue
type Params struct {
	IssueFee          sdk.Coin `json:"issue_fee"`
	MintFee           sdk.Coin `json:"mint_fee"`
	FreezeFee         sdk.Coin `json:"freeze_fee"`
	UnFreezeFee       sdk.Coin `json:"unfreeze_fee"`
	BurnFee           sdk.Coin `json:"burn_fee"`
	BurnFromFee       sdk.Coin `json:"burn_from_fee"`
	TransferOwnerFee  sdk.Coin `json:"transfer_owner_fee"`
	DisableFeatureFee sdk.Coin `json:"disable_feature_fee"`
	DescribeFee       sdk.Coin `json:"describe_fee"`
}

// ParamKeyTable for auth module
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of auth module's parameters.
// nolint
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{ParamStoreKeyIssueFee, &p.IssueFee},
		{ParamStoreKeyMintFee, &p.MintFee},
		{ParamStoreKeyFreezeFee, &p.FreezeFee},
		{ParamStoreKeyUnFreezeFee, &p.UnFreezeFee},
		{ParamStoreKeyBurnFee, &p.BurnFee},
		{ParamStoreKeyBurnFromFee, &p.BurnFromFee},
		{ParamStoreKeyTransferOwnerFee, &p.TransferOwnerFee},
		{ParamStoreKeyIssueDisableFeatureFee, &p.DisableFeatureFee},
		{ParamStoreKeyDescribeFee, &p.DescribeFee},
	}
}

// Checks equality of Params
func (dp Params) Equal(dp2 Params) bool {
	b1 := msgs.MsgCdc.MustMarshalBinaryBare(dp)
	b2 := msgs.MsgCdc.MustMarshalBinaryBare(dp2)
	return bytes.Equal(b1, b2)
}

// DefaultParams returns a default set of parameters.
func DefaultParams(denom string) Params {
	return Params{
		IssueFee:          sdk.NewCoin(denom, sdk.NewIntWithDecimal(100, 18)),
		MintFee:           sdk.NewCoin(denom, sdk.NewIntWithDecimal(1000, 18)),
		FreezeFee:         sdk.NewCoin(denom, sdk.NewIntWithDecimal(1000, 18)),
		UnFreezeFee:       sdk.NewCoin(denom, sdk.NewIntWithDecimal(1000, 18)),
		BurnFee:           sdk.NewCoin(denom, sdk.NewIntWithDecimal(100, 18)),
		BurnFromFee:       sdk.NewCoin(denom, sdk.NewIntWithDecimal(100, 18)),
		TransferOwnerFee:  sdk.NewCoin(denom, sdk.NewIntWithDecimal(100, 18)),
		DisableFeatureFee: sdk.NewCoin(denom, sdk.NewIntWithDecimal(100, 18)),
		DescribeFee:       sdk.NewCoin(denom, sdk.NewIntWithDecimal(100, 18)),
	}
}

func (dp Params) String() string {
	return fmt.Sprintf(`Params:
  IssueFee:			%s
  MintFee:			%s
  FreezeFee:			%s
  UnFreezeFee:			%s
  BurnFee:			%s
  BurnFromFee:			%s
  TransferOwnerFee:			%s
  DisableFeatureFee:			%s
  DescribeFee:			%s`,
		dp.IssueFee.String(),
		dp.MintFee.String(),
		dp.FreezeFee.String(),
		dp.UnFreezeFee.String(),
		dp.BurnFee.String(),
		dp.BurnFromFee.String(),
		dp.TransferOwnerFee.String(),
		dp.DisableFeatureFee.String(),
		dp.DescribeFee.String(),
	)
}
