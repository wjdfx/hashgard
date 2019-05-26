package gov

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var cdc = MakeCodec()

type ProposalParam struct {
	Key		 string	`json:"key"`
	Value	 string	`json:"value"`
}

type ProposalParams []ProposalParam

func MakeCodec() *codec.Codec {
	cdc := codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}

func ValidateProposalParam(proposalParam ProposalParam) sdk.Error {
	// check key
	switch proposalParam.Key {
	case "auth/max_memo_characters", "auth/tx_sig_limit", "auth/tx_size_cost_per_byte",
		"mint/blocks_per_year":
		var val uint64
		err := cdc.UnmarshalJSON([]byte(proposalParam.Value), &val)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case "bank/send_enabled", "distribution/withdraw_addr_enabled":
		var val bool
		err := cdc.UnmarshalJSON([]byte(proposalParam.Value), &val)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case "distribution/community_tax", "distribution/base_proposer_reward", "distribution/bonus_proposer_reward",
		"gov/quorum", "gov/threshold", "gov/veto",
		"mint/inflation_rate_change", "mint/inflation_max", "mint/inflation_min", "mint/goal_bonded",
		"slashing/min_signed_per_window", "slashing/slash_fraction_double_sign", "slashing/slash_fraction_downtime":
		var val sdk.Dec
		err := cdc.UnmarshalJSON([]byte(proposalParam.Value), &val)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case "gov/min_deposit":
		var val sdk.Coins
		err := cdc.UnmarshalJSON([]byte(proposalParam.Value), &val)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case "gov/max_deposit_period", "gov/voting_period", "slashing/downtime_jail_duration",
		"slashing/max_evidence_age", "staking/unbonding_time":
		var val time.Duration
		err := cdc.UnmarshalJSON([]byte(proposalParam.Value), &val)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case "mint/mint_denom":
		var val string
		err := cdc.UnmarshalJSON([]byte(proposalParam.Value), &val)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case "slashing/signed_blocks_window":
		var val int64
		err := cdc.UnmarshalJSON([]byte(proposalParam.Value), &val)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case "staking/max_validators", "staking/max_entries":
		var val uint16
		err := cdc.UnmarshalJSON([]byte(proposalParam.Value), &val)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	default:
		return ErrInvalidParamKey(DefaultCodespace, proposalParam.Key)
	}

	return nil
}

