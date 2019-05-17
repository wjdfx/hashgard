package gov

import(
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ProposalParam struct {
	Subspace string `json:"subspace"`
	Key		 string	`json:"key"`
	Value	 string	`json:"value"`
}

type ProposalParams []ProposalParam

var (
	ChangableParams map[string]map[string]interface{}
	cdc
)

func init() {

}

func MakeCodec() *codec.Codec {
	cdc := codec.New()
	auth.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	staking.RegisterCodec(cdc)
	distribution.RegisterCodec(cdc)
	slashing.RegisterCodec(cdc)
	RegisterCodec(cdc)
	crisis.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	return cdc
}

func RegisterChangableParam(subspace string, key string, val interface{}) {
	ChangableParams[subspace][key] = val
}

