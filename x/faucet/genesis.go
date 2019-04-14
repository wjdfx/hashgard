package faucet

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

var (
	DefaultFaucetOrigin = sdk.AccAddress(crypto.AddressHash([]byte("defaultFaucetOriginAddress")))
	DefaultSendCoins	= []sdk.Coin{
		sdk.NewCoin("gard", sdk.NewInt(100)),
		sdk.NewCoin("apple", sdk.NewInt(100)),
	}
	DefaultLimitCoins	= []sdk.Coin{
		sdk.NewCoin("gard", sdk.NewInt(300)),
		sdk.NewCoin("apple", sdk.NewInt(300)),
	}
)

type GenesisState struct {
	FaucetOrigin	sdk.AccAddress	`json:"faucet_origin"`
	SendCoins		sdk.Coins		`json:"send_coins"`
	LimitCoins		sdk.Coins		`json:"limit_coins"`
}

func NewGenesisState(origin sdk.AccAddress, send sdk.Coins, limit sdk.Coins) GenesisState {
	return GenesisState{
		FaucetOrigin:	origin,
		SendCoins:		send,
		LimitCoins:		limit,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		FaucetOrigin:	DefaultFaucetOrigin,
		SendCoins:		DefaultSendCoins,
		LimitCoins:		DefaultLimitCoins,
	}
}

func (data GenesisState) Equal(data2 GenesisState) bool {
	b1 := msgCdc.MustMarshalBinaryBare(data)
	b2 := msgCdc.MustMarshalBinaryBare(data2)
	return bytes.Equal(b1, b2)
}

func (data GenesisState) IsEmpty() bool {
	emptyGenState := GenesisState{}
	return data.Equal(emptyGenState)
}

func ValidateGenesis(data GenesisState) error {
	if !data.FaucetOrigin.Empty() && (data.SendCoins.Empty() || data.SendCoins.IsZero()) {
		return sdk.NewError(DefaultCodespace, CodeInvalidGenesis, "if set faucet_origin, the send_coins must be valid")
	}
	return nil
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	keeper.setFaucetOrigin(ctx, data.FaucetOrigin)
	keeper.setSendCoins(ctx, data.SendCoins)
	keeper.setLimitCoins(ctx, data.LimitCoins)
}

func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	faucetOrigin := keeper.GetFaucetOrigin(ctx)
	sendCoins := keeper.GetSendCoins(ctx)
	limitCoins := keeper.GetLimitCoins(ctx)

	return GenesisState {
		FaucetOrigin:	faucetOrigin,
		SendCoins:		sendCoins,
		LimitCoins:		limitCoins,
	}
}