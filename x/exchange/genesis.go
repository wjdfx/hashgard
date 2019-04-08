package exchange

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/exchange/keeper"
)

type GenesisState struct {
	StartingOrderId		uint64	`json:"starting_order_id"`
}

func NewGenesisState(startingOrderId uint64) GenesisState {
	return GenesisState{
		StartingOrderId: startingOrderId,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		StartingOrderId: 1,
	}
}

// Checks whether 2 GenesisState structs are equivalent.
func (data GenesisState) Equal(data2 GenesisState) bool {
	b1 := MsgCdc.MustMarshalBinaryBare(data)
	b2 := MsgCdc.MustMarshalBinaryBare(data2)
	return bytes.Equal(b1, b2)
}

// Returns if a GenesisState is empty or has data in it
func (data GenesisState) IsEmpty() bool {
	emptyGenState := GenesisState{}
	return data.Equal(emptyGenState)
}

func ValidateGenesis(data GenesisState) error {
	return nil
}

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data GenesisState) {
	err := keeper.SetInitialOrderId(ctx, data.StartingOrderId)
	if err != nil {
		// TODO: Handle this with #870
		panic(err)
	}
}

func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) GenesisState {
	startingOrderId, _ := keeper.PeekCurrentOrderId(ctx)
	return GenesisState{
		StartingOrderId: startingOrderId,
	}
}
