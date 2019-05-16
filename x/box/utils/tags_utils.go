package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/tags"
)

func GetBoxTags(boxID string, boxType string, sender sdk.AccAddress) sdk.Tags {
	return sdk.NewTags(
		tags.Category, tags.TxCategory,
		tags.BoxID, boxID,
		tags.BoxType, boxType,
		tags.Sender, sender.String(),
	)
}
