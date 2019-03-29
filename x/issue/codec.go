package issue

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/hashgard/hashgard/x/issue/domain"
	"github.com/hashgard/hashgard/x/issue/msgs"
)

var msgCdc = codec.New()

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(msgs.MsgIssue{}, "issue/msg/MsgIssue", nil)
	cdc.RegisterConcrete(msgs.MsgIssueMint{}, "issue/msg/MsgIssueMint", nil)
	cdc.RegisterConcrete(msgs.MsgIssueBurn{}, "issue/msg/MsgIssueBurn", nil)
	cdc.RegisterConcrete(msgs.MsgIssueFinishMinting{}, "issue/msg/MsgIssueFinishMinting", nil)

	cdc.RegisterInterface((*domain.Issue)(nil), nil)
	cdc.RegisterConcrete(&domain.CoinIssueInfo{}, "issue/domain/CoinIssueInfo", nil)
}
func init() {
	RegisterCodec(msgCdc)
}
