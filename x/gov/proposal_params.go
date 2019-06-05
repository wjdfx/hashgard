package gov

import (
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	communityTax          = "distribution/community_tax"
	foundationAddress     = "distribution/foundation_address"
	inflation             = "mint/inflation"
	inflationBase         = "mint/inflation_base"
	minDeposit            = "gov/min_deposit"
	signedBlocksWindow    = "slashing/signed_blocks_window"
	minSignedPerWindow    = "slashing/min_signed_per_window"
	slashFractionDowntime = "slashing/slash_fraction_downtime"
	downtimeJailDuration  = "slashing/downtime_jail_duration"
	maxValidators         = "staking/max_validators"
	unbondingTime         = "staking/unbonding_time"
)

//fee
const (
	Fee         = "Fee"
	BoxModule   = "box/"
	IssueModule = "issue/"
)

//fee/box/lock_create=10000,fee/box/deposit_box_create=10000,fee/box/future_box_create=10000,
//fee/box/disable_feature=10000,fee/box/describe_fee=10000,fee/issue/create=10000,fee/issue/mint=10000
var cdc = MakeCodec()

type ProposalParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ProposalParams []ProposalParam

func MakeCodec() *codec.Codec {
	cdc := codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}

func ValidateProposalParam(proposalParam ProposalParam) sdk.Error {
	if strings.HasPrefix(proposalParam.Key, BoxModule) || strings.HasPrefix(proposalParam.Key, IssueModule) {
		return nil
	}
	// check key
	switch proposalParam.Key {
	case communityTax, inflation, minSignedPerWindow, slashFractionDowntime:
		var val sdk.Dec
		err := cdc.UnmarshalJSON([]byte(proposalParam.Value), &val)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case minDeposit:
		var val sdk.Coins
		err := cdc.UnmarshalJSON([]byte(proposalParam.Value), &val)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case downtimeJailDuration, unbondingTime:
		var val time.Duration
		err := cdc.UnmarshalJSON([]byte(proposalParam.Value), &val)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case signedBlocksWindow:
		var val int64
		err := cdc.UnmarshalJSON([]byte(proposalParam.Value), &val)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case maxValidators:
		var val uint16
		err := cdc.UnmarshalJSON([]byte(proposalParam.Value), &val)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case inflationBase:
		var val sdk.Int
		err := cdc.UnmarshalJSON([]byte(proposalParam.Value), &val)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case foundationAddress:
		var val sdk.AccAddress
		err := cdc.UnmarshalJSON([]byte(proposalParam.Value), &val)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	default:
		return ErrInvalidParamKey(DefaultCodespace, proposalParam.Key)
	}

	return nil
}
