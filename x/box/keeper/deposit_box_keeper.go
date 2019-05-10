package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/types"
)

//Process deposit box

func (keeper Keeper) ProcessDepositBoxCreate(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	box.Deposit.Status = types.BoxCreated
	box.Deposit.TotalDeposit = sdk.ZeroInt()
	box.Deposit.Share = sdk.ZeroInt()
	box.Deposit.Coupon = box.Deposit.Interest.Amount.Quo(box.TotalAmount.Amount.Quo(box.Deposit.Price))
	keeper.InsertActiveBoxQueue(ctx, box.Deposit.StartTime, box.BoxId)
	return nil
}
func (keeper Keeper) ProcessDepositBoxInterest(ctx sdk.Context, boxID string, sender sdk.AccAddress, interest sdk.Coin, operation string) (*types.BoxInfo, sdk.Error) {

	switch operation {
	case types.Injection:
		return keeper.injectionDepositBoxInterest(ctx, boxID, sender, interest)
	case types.Fetch:
		return keeper.fetchInterestFromDepositBox(ctx, boxID, sender, interest)

	}
	return nil, errors.ErrUnknownOperation()
}
func (keeper Keeper) injectionDepositBoxInterest(ctx sdk.Context, boxID string, sender sdk.AccAddress, interest sdk.Coin) (*types.BoxInfo, sdk.Error) {
	box := keeper.GetBox(ctx, boxID)
	if box == nil {
		return nil, errors.ErrUnknownBox(boxID)
	}
	if types.BoxCreated != box.Deposit.Status {
		return nil, errors.ErrNotAllowedOperation(box.Deposit.Status)
	}
	if box.Deposit.Interest.Denom != interest.Denom {
		return nil, errors.ErrInterestInjectionNotValid(interest)
	}
	if len(box.Deposit.InterestInjection) >= types.BoxMaxInjectionInterest {
		return nil, errors.ErrInterestInjectionNotValid(interest)
	}

	totalInterest := sdk.ZeroInt()
	if box.Deposit.InterestInjection == nil {
		box.Deposit.InterestInjection = []types.AddressDeposit{{Address: sender, Amount: interest.Amount}}
	} else {
		exist := false
		for i, v := range box.Deposit.InterestInjection {
			totalInterest = totalInterest.Add(v.Amount)
			if v.Address.Equals(sender) {
				box.Deposit.InterestInjection[i].Amount = box.Deposit.InterestInjection[i].Amount.Add(interest.Amount)
				exist = true
			}
		}
		if !exist {
			box.Deposit.InterestInjection = append(box.Deposit.InterestInjection, types.NewAddressDeposit(sender, interest.Amount))
		}
	}
	totalInterest = totalInterest.Add(interest.Amount)
	if totalInterest.GT(box.Deposit.Interest.Amount) {
		return nil, errors.ErrInterestInjectionNotValid(interest)
	}

	_, err := keeper.ck.SubtractCoins(ctx, sender, sdk.Coins{interest})
	if err != nil {
		return nil, err
	}
	keeper.setBox(ctx, box)

	return box, nil
}
func (keeper Keeper) fetchInterestFromDepositBox(ctx sdk.Context, boxID string, sender sdk.AccAddress, interest sdk.Coin) (*types.BoxInfo, sdk.Error) {
	box := keeper.GetBox(ctx, boxID)
	if box == nil {
		return nil, errors.ErrUnknownBox(boxID)
	}
	if types.BoxCreated != box.Deposit.Status {
		return nil, errors.ErrNotAllowedOperation(box.Deposit.Status)
	}
	if box.Deposit.Interest.Denom != interest.Denom {
		return nil, errors.ErrInterestInjectionNotValid(interest)
	}
	if box.Deposit.InterestInjection == nil {
		return nil, errors.ErrInterestFetchNotValid(interest)
	} else {
		remove := -1
		for i, v := range box.Deposit.InterestInjection {
			if v.Address.Equals(sender) {
				if box.Deposit.InterestInjection[i].Amount.LT(interest.Amount) {
					return nil, errors.ErrNotEnoughAmount()
				}
				box.Deposit.InterestInjection[i].Amount = box.Deposit.InterestInjection[i].Amount.Sub(interest.Amount)
				if box.Deposit.InterestInjection[i].Amount.IsZero() {
					remove = i
				}
			}
		}
		if remove != -1 {
			box.Deposit.InterestInjection = append(box.Deposit.InterestInjection[:remove], box.Deposit.InterestInjection[remove+1:]...)
		}
	}
	_, err := keeper.ck.AddCoins(ctx, sender, sdk.Coins{interest})
	if err != nil {
		return nil, err
	}

	keeper.setBox(ctx, box)
	return box, nil
}
func (keeper Keeper) ProcessDepositBoxDeposit(ctx sdk.Context, boxID string, sender sdk.AccAddress, deposit sdk.Coin, operation string) (*types.BoxInfo, sdk.Error) {

	switch operation {
	case types.DepositTo:
		return keeper.depositToDepositBox(ctx, boxID, sender, deposit)
	case types.Fetch:
		return keeper.fetchDepositFromDepositBox(ctx, boxID, sender, deposit)

	}
	return nil, errors.ErrUnknownOperation()
}
func (keeper Keeper) depositToDepositBox(ctx sdk.Context, boxID string, sender sdk.AccAddress, deposit sdk.Coin) (*types.BoxInfo, sdk.Error) {
	boxInfo := keeper.GetBox(ctx, boxID)
	if boxInfo == nil {
		return nil, errors.ErrUnknownBox(boxID)
	}
	if types.DepositBoxDeposit != boxInfo.Deposit.Status {
		return nil, errors.ErrNotAllowedOperation(boxInfo.Deposit.Status)
	}
	if !deposit.Amount.Mod(boxInfo.Deposit.Price).IsZero() {
		return nil, errors.ErrAmountNotValid(deposit.Denom)
	}
	if boxInfo.TotalAmount.Denom != deposit.Denom {
		return nil, errors.ErrAmountNotValid(deposit.Denom)
	}
	boxInfo.Deposit.TotalDeposit = boxInfo.Deposit.TotalDeposit.Add(deposit.Amount)
	if boxInfo.Deposit.TotalDeposit.GT(boxInfo.TotalAmount.Amount) {
		return nil, errors.ErrAmountNotValid(deposit.Denom)
	}
	_, err := keeper.ck.SubtractCoins(ctx, sender, sdk.Coins{deposit})
	if err != nil {
		return nil, err
	}
	keeper.addAddressDeposit(ctx, boxID, sender, deposit.Amount)
	boxInfo.Deposit.Share = boxInfo.Deposit.Share.Add(deposit.Amount.Quo(boxInfo.Deposit.Price))
	keeper.setBox(ctx, boxInfo)
	return boxInfo, nil
}
func (keeper Keeper) fetchDepositFromDepositBox(ctx sdk.Context, boxID string, sender sdk.AccAddress, deposit sdk.Coin) (*types.BoxInfo, sdk.Error) {
	boxInfo := keeper.GetBox(ctx, boxID)
	if boxInfo == nil {
		return nil, errors.ErrUnknownBox(boxID)
	}
	if types.DepositBoxDeposit != boxInfo.Deposit.Status {
		return nil, errors.ErrNotAllowedOperation(boxInfo.Deposit.Status)
	}
	if !deposit.Amount.Mod(boxInfo.Deposit.Price).IsZero() {
		return nil, errors.ErrAmountNotValid(deposit.Amount.String())
	}
	amount := keeper.GetDepositByAddress(ctx, boxID, sender)
	if amount.LT(deposit.Amount) {
		return nil, errors.ErrAmountNotValid(deposit.Denom)
	}
	_, err := keeper.ck.AddCoins(ctx, sender, sdk.Coins{deposit})
	if err != nil {
		return nil, err
	}
	amount = amount.Sub(deposit.Amount)
	if amount.IsZero() {
		keeper.removeAddressDeposit(ctx, boxID, sender)
	} else {
		keeper.setAddressDeposit(ctx, boxID, sender, amount)
	}

	boxInfo.Deposit.Share = boxInfo.Deposit.Share.Sub(deposit.Amount.Quo(boxInfo.Deposit.Price))
	keeper.setBox(ctx, boxInfo)
	return boxInfo, nil
}
func (keeper Keeper) ProcessDepositBoxByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {

	var err sdk.Error
	switch box.Deposit.Status {
	case types.BoxCreated:
		err = keeper.processBoxCreatedByEndBlocker(ctx, box)
	case types.DepositBoxDeposit:
		err = keeper.processDepositBoxDepositByEndBlocker(ctx, box)
	case types.DepositBoxInterest:
		err = keeper.processDepositBoxInterestByEndBlocker(ctx, box)
	}
	return err
}
func (keeper Keeper) processBoxCreatedByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if box.Deposit.Status != types.BoxCreated {
		return nil
	}
	totalInterest := sdk.ZeroInt()
	if box.Deposit.InterestInjection != nil {
		for _, v := range box.Deposit.InterestInjection {
			totalInterest = totalInterest.Add(v.Amount)
		}
	}
	if box.Deposit.Interest.Amount.Equal(totalInterest) {
		box.Deposit.Status = types.DepositBoxDeposit
		keeper.InsertActiveBoxQueue(ctx, box.Deposit.EstablishTime, box.BoxId)
	} else {
		box.Deposit.Status = types.BoxClosed
		for _, v := range box.Deposit.InterestInjection {
			_, err := keeper.ck.AddCoins(ctx, v.Address, sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Denom, v.Amount)))
			if err != nil {
				return err
			}
		}
		//keeper.RemoveBox(ctx, box)
	}

	keeper.setBox(ctx, box)
	keeper.RemoveFromActiveBoxQueue(ctx, box.Deposit.StartTime, box.BoxId)

	return nil
}
func (keeper Keeper) processDepositBoxDepositByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if box.Deposit.Status != types.DepositBoxDeposit {
		return nil
	}
	store := ctx.KVStore(keeper.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, PrefixKeyDeposit(box.BoxId))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		if len(bz) == 0 {
			continue
		}
		var amount sdk.Int
		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &amount)
		if box.Deposit.TotalDeposit.GTE(box.Deposit.BottomLine) {
			box.Deposit.Status = types.DepositBoxInterest
			_, err := keeper.ck.AddCoins(ctx, GetAddressFromKeyAddressDeposit(iterator.Key()), sdk.NewCoins(sdk.NewCoin(box.BoxId, amount)))
			if err != nil {
				return err
			}
		} else {
			box.Deposit.Status = types.BoxClosed
			_, err := keeper.ck.AddCoins(ctx, GetAddressFromKeyAddressDeposit(iterator.Key()), sdk.NewCoins(sdk.NewCoin(box.TotalAmount.Denom, amount)))
			if err != nil {
				return err
			}
			//keeper.RemoveBox(ctx, box)
		}
	}

	keeper.InsertActiveBoxQueue(ctx, box.Deposit.MaturityTime, box.BoxId)
	keeper.setBox(ctx, box)

	keeper.RemoveFromActiveBoxQueue(ctx, box.Deposit.EstablishTime, box.BoxId)

	return nil
}
func (keeper Keeper) processDepositBoxInterestByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if box.Deposit.Status != types.DepositBoxInterest {
		return nil
	}
	store := ctx.KVStore(keeper.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, PrefixKeyDeposit(box.BoxId))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		if len(bz) == 0 {
			continue
		}
		var amount sdk.Int
		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &amount)
		address := GetAddressFromKeyAddressDeposit(iterator.Key())

		_, err := keeper.ck.SubtractCoins(ctx, address, sdk.NewCoins(sdk.NewCoin(box.BoxId, amount)))
		if err != nil {
			return err
		}
		_, err = keeper.ck.AddCoins(ctx, address, sdk.NewCoins(sdk.NewCoin(box.TotalAmount.Denom, amount)))
		if err != nil {
			return err
		}
		_, err = keeper.ck.AddCoins(ctx, address, sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Denom, amount.Quo(box.Deposit.Price).Mul(box.Deposit.Coupon))))
		if err != nil {
			return err
		}
	}
	keeper.RemoveFromActiveBoxQueue(ctx, box.Deposit.MaturityTime, box.BoxId)
	box.Deposit.Status = types.BoxFinished
	keeper.setBox(ctx, box)
	return nil
}

//Queries
//Query deposit list
func (keeper Keeper) QueryDepositListFromDepositBox(ctx sdk.Context, boxID string, accAddress sdk.AccAddress) types.DepositBoxDepositToList {
	var list = make(types.DepositBoxDepositToList, 0)
	if accAddress != nil && !accAddress.Empty() {
		amount := keeper.GetDepositByAddress(ctx, boxID, accAddress)
		list = append(list, types.NewAddressDeposit(accAddress, amount))
		return list
	}
	store := ctx.KVStore(keeper.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, PrefixKeyDeposit(boxID))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		if len(bz) == 0 {
			continue
		}
		var amount sdk.Int
		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &amount)
		list = append(list, types.NewAddressDeposit(GetAddressFromKeyAddressDeposit(iterator.Key()), amount))
	}
	return list
}
