package keeper

import (
	"github.com/hashgard/hashgard/x/box/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/types"
	issueerr "github.com/hashgard/hashgard/x/issue/errors"
)

//Process deposit box

func (keeper Keeper) ProcessDepositBoxCreate(ctx sdk.Context, box *types.BoxInfo) sdk.Error {

	coinIssueInfo := keeper.ik.GetIssue(ctx, box.Deposit.Interest.Token.Denom)
	if coinIssueInfo == nil {
		return issueerr.ErrUnknownIssue(box.Deposit.Interest.Token.Denom)
	}
	if box.Deposit.Interest.Decimals != coinIssueInfo.GetDecimals() {
		return errors.ErrDecimalsNotValid(box.Deposit.Interest.Decimals)
	}

	box.BoxStatus = types.BoxCreated
	box.Deposit.TotalDeposit = sdk.ZeroInt()
	//box.Deposit.TotalInterestInjection = sdk.ZeroInt()
	box.Deposit.Share = sdk.ZeroInt()
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
	if types.BoxCreated != box.BoxStatus {
		return nil, errors.ErrNotAllowedOperation(box.BoxStatus)
	}
	if box.Deposit.Interest.Token.Denom != interest.Denom {
		return nil, errors.ErrInterestInjectionNotValid(interest)
	}
	if len(box.Deposit.InterestInjections) >= types.BoxMaxInjectionInterest {
		return nil, errors.ErrInterestInjectionNotValid(interest)
	}

	totalInterest := sdk.ZeroInt()
	if box.Deposit.InterestInjections == nil {
		box.Deposit.InterestInjections = []types.AddressDeposit{{Address: sender, Amount: interest.Amount}}
	} else {
		exist := false
		for i, v := range box.Deposit.InterestInjections {
			totalInterest = totalInterest.Add(v.Amount)
			if v.Address.Equals(sender) {
				box.Deposit.InterestInjections[i].Amount = box.Deposit.InterestInjections[i].Amount.Add(interest.Amount)
				exist = true
			}
		}
		if !exist {
			box.Deposit.InterestInjections = append(box.Deposit.InterestInjections, types.NewAddressDeposit(sender, interest.Amount))
		}
	}
	totalInterest = totalInterest.Add(interest.Amount)
	if totalInterest.GT(box.Deposit.Interest.Token.Amount) {
		return nil, errors.ErrInterestInjectionNotValid(interest)
	}

	if err := keeper.SendDepositedCoin(ctx, sender, sdk.Coins{interest}, boxID); err != nil {
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
	if types.BoxCreated != box.BoxStatus {
		return nil, errors.ErrNotAllowedOperation(box.BoxStatus)
	}
	if box.Deposit.Interest.Token.Denom != interest.Denom {
		return nil, errors.ErrInterestInjectionNotValid(interest)
	}
	if box.Deposit.InterestInjections == nil {
		return nil, errors.ErrInterestFetchNotValid(interest)
	} else {
		remove := -1
		for i, v := range box.Deposit.InterestInjections {
			if v.Address.Equals(sender) {
				if box.Deposit.InterestInjections[i].Amount.LT(interest.Amount) {
					return nil, errors.ErrNotEnoughAmount()
				}
				box.Deposit.InterestInjections[i].Amount = box.Deposit.InterestInjections[i].Amount.Sub(interest.Amount)
				if box.Deposit.InterestInjections[i].Amount.IsZero() {
					remove = i
				}
			}
		}
		if remove != -1 {
			box.Deposit.InterestInjections = append(box.Deposit.InterestInjections[:remove], box.Deposit.InterestInjections[remove+1:]...)
		}
	}
	if err := keeper.FetchDepositedCoin(ctx, sender, sdk.Coins{interest}, boxID); err != nil {
		return nil, err
	}
	keeper.setBox(ctx, box)
	return box, nil
}
func (keeper Keeper) processDepositBoxDeposit(ctx sdk.Context, box *types.BoxInfo, sender sdk.AccAddress, deposit sdk.Coin, operation string) sdk.Error {

	switch operation {
	case types.DepositTo:
		return keeper.depositToDepositBox(ctx, box, sender, deposit)
	case types.Fetch:
		return keeper.fetchDepositFromDepositBox(ctx, box, sender, deposit)

	}
	return errors.ErrUnknownOperation()
}
func (keeper Keeper) depositToDepositBox(ctx sdk.Context, box *types.BoxInfo, sender sdk.AccAddress, deposit sdk.Coin) sdk.Error {
	if !deposit.Amount.Mod(box.Deposit.Price).IsZero() {
		return errors.ErrAmountNotValid(deposit.Denom)
	}
	if box.TotalAmount.Token.Denom != deposit.Denom {
		return errors.ErrAmountNotValid(deposit.Denom)
	}
	box.Deposit.TotalDeposit = box.Deposit.TotalDeposit.Add(deposit.Amount)
	if box.Deposit.TotalDeposit.GT(box.TotalAmount.Token.Amount) {
		return errors.ErrAmountNotValid(deposit.Denom)
	}

	if err := keeper.SendDepositedCoin(ctx, sender, sdk.Coins{deposit}, box.BoxId); err != nil {
		return err
	}
	keeper.addAddressDeposit(ctx, box.BoxId, sender, deposit.Amount)
	box.Deposit.Share = box.Deposit.Share.Add(deposit.Amount.Quo(box.Deposit.Price))

	keeper.setBox(ctx, box)
	return nil
}
func (keeper Keeper) fetchDepositFromDepositBox(ctx sdk.Context, box *types.BoxInfo, sender sdk.AccAddress, deposit sdk.Coin) sdk.Error {

	if !deposit.Amount.Mod(box.Deposit.Price).IsZero() {
		return errors.ErrAmountNotValid(deposit.Amount.String())
	}
	amount := keeper.GetDepositByAddress(ctx, box.BoxId, sender)
	if amount.LT(deposit.Amount) {
		return errors.ErrAmountNotValid(deposit.Denom)
	}
	if err := keeper.FetchDepositedCoin(ctx, sender, sdk.Coins{deposit}, box.BoxId); err != nil {
		return err
	}
	amount = amount.Sub(deposit.Amount)
	if amount.IsZero() {
		keeper.removeAddressDeposit(ctx, box.BoxId, sender)
	} else {
		keeper.setAddressDeposit(ctx, box.BoxId, sender, amount)
	}

	box.Deposit.Share = box.Deposit.Share.Sub(deposit.Amount.Quo(box.Deposit.Price))
	keeper.setBox(ctx, box)
	return nil
}
func (keeper Keeper) ProcessDepositBoxByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {

	var err sdk.Error
	switch box.BoxStatus {
	case types.BoxCreated:
		err = keeper.processBoxCreatedByEndBlocker(ctx, box)
	case types.BoxDepositing:
		err = keeper.processDepositBoxDepositToByEndBlocker(ctx, box)
	case types.DepositBoxInterest:
		err = keeper.processDepositBoxInterestByEndBlocker(ctx, box)
	}
	return err
}
func (keeper Keeper) backBoxInterestInjections(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	for _, v := range box.Deposit.InterestInjections {
		if err := keeper.FetchDepositedCoin(ctx, v.Address, sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Token.Denom, v.Amount)), box.BoxId); err != nil {
			return err
		}
	}
	return nil
}
func (keeper Keeper) backBoxUnUsedInterestInjections(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	totalCoupon := box.TotalAmount.Token.Amount.Quo(box.Deposit.Price)
	if totalCoupon.Equal(box.Deposit.Share) {
		return nil
	}

	unused := utils.MulMaxPrecisionByDecimal(box.Deposit.PerCoupon.MulInt(totalCoupon.Sub(box.Deposit.Share)), box.Deposit.Interest.Decimals)
	interestInjectionsLen := len(box.Deposit.InterestInjections)

	if interestInjectionsLen == 0 {
		if err := keeper.FetchDepositedCoin(ctx, box.Deposit.InterestInjections[0].Address,
			sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Token.Denom, unused)), box.BoxId); err != nil {
			return err
		}
		box.Deposit.InterestInjections[0].Amount = box.Deposit.InterestInjections[0].Amount.Sub(unused)
	} else {
		total := sdk.ZeroInt()
		for i, v := range box.Deposit.InterestInjections {
			var amount sdk.Int
			if i == interestInjectionsLen-1 {
				amount = unused.Sub(total)
			} else {
				amount = sdk.NewDecFromInt(v.Amount).QuoInt(box.Deposit.Interest.Token.Amount).MulInt(unused).TruncateInt()
			}
			if err := keeper.FetchDepositedCoin(ctx, v.Address, sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Token.Denom, amount)), box.BoxId); err != nil {
				return err
			}
			box.Deposit.InterestInjections[i].Amount = box.Deposit.InterestInjections[i].Amount.Sub(amount)
			total = total.Add(amount)
		}
	}

	box.Deposit.Interest.Token.Amount = box.Deposit.Interest.Token.Amount.Sub(unused)
	keeper.setBox(ctx, box)
	return nil
}
func (keeper Keeper) backBoxAllDeposit(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	for _, v := range box.Deposit.InterestInjections {
		if err := keeper.FetchDepositedCoin(ctx, v.Address, sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Token.Denom, v.Amount)), box.BoxId); err != nil {
			return err
		}
	}
	return nil
}
func (keeper Keeper) processBoxCreatedByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if box.BoxStatus != types.BoxCreated {
		return nil
	}
	totalInterest := sdk.ZeroInt()
	if box.Deposit.InterestInjections != nil {
		for _, v := range box.Deposit.InterestInjections {
			totalInterest = totalInterest.Add(v.Amount)
		}
	}
	if box.Deposit.Interest.Token.Amount.Equal(totalInterest) {
		box.BoxStatus = types.BoxDepositing
		keeper.InsertActiveBoxQueue(ctx, box.Deposit.EstablishTime, box.BoxId)
		keeper.setBox(ctx, box)
	} else {
		box.BoxStatus = types.BoxClosed
		if err := keeper.backBoxInterestInjections(ctx, box); err != nil {
			return err
		}
		keeper.RemoveBox(ctx, box)
	}

	keeper.RemoveFromActiveBoxQueue(ctx, box.Deposit.StartTime, box.BoxId)

	return nil
}
func (keeper Keeper) processDepositBoxDepositToByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if box.BoxStatus != types.BoxDepositing {
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
			box.BoxStatus = types.DepositBoxInterest
			_, err := keeper.ck.SubtractCoins(ctx, keeper.getDepositedCoinsAddress(box.BoxId), sdk.NewCoins(sdk.NewCoin(box.TotalAmount.Token.Denom, amount)))
			if err != nil {
				return err
			}
			_, err = keeper.ck.AddCoins(ctx, GetAddressFromKeyAddressDeposit(iterator.Key()), sdk.NewCoins(sdk.NewCoin(box.BoxId, amount.Quo(box.Deposit.Price))))
			if err != nil {
				return err
			}
		} else {
			box.BoxStatus = types.BoxClosed
			if err := keeper.FetchDepositedCoin(ctx, GetAddressFromKeyAddressDeposit(iterator.Key()),
				sdk.NewCoins(sdk.NewCoin(box.TotalAmount.Token.Denom, amount)), box.BoxId); err != nil {
				return err
			}
		}
	}

	if box.BoxStatus == types.BoxClosed {
		if err := keeper.backBoxInterestInjections(ctx, box); err != nil {
			return err
		}
		keeper.RemoveBox(ctx, box)
	} else {
		if err := keeper.backBoxUnUsedInterestInjections(ctx, box); err != nil {
			return err
		}
		keeper.InsertActiveBoxQueue(ctx, box.Deposit.MaturityTime, box.BoxId)
		keeper.setBox(ctx, box)
	}

	keeper.RemoveFromActiveBoxQueue(ctx, box.Deposit.EstablishTime, box.BoxId)
	return nil
}
func (keeper Keeper) processDepositBoxInterestByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if box.BoxStatus != types.DepositBoxInterest {
		return nil
	}
	store := ctx.KVStore(keeper.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, PrefixKeyDeposit(box.BoxId))
	defer iterator.Close()
	totalInterest := sdk.ZeroInt()
	var address sdk.AccAddress
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		if len(bz) == 0 {
			continue
		}
		var amount sdk.Int
		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &amount)
		address = GetAddressFromKeyAddressDeposit(iterator.Key())

		share := amount.Quo(box.Deposit.Price)
		_, err := keeper.ck.SubtractCoins(ctx, address, sdk.NewCoins(sdk.NewCoin(box.BoxId, share)))
		if err != nil {
			return err
		}
		_, err = keeper.ck.AddCoins(ctx, address, sdk.NewCoins(sdk.NewCoin(box.TotalAmount.Token.Denom, amount)))
		if err != nil {
			return err
		}
		interest := utils.MulMaxPrecisionByDecimal(box.Deposit.PerCoupon.MulInt(share), box.Deposit.Interest.Decimals)
		_, err = keeper.ck.AddCoins(ctx, address, sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Token.Denom, interest)))
		if err != nil {
			return err
		}
		totalInterest = totalInterest.Add(interest)
	}
	if !totalInterest.Equal(box.Deposit.Interest.Token.Amount) {
		_, err := keeper.ck.AddCoins(ctx, address, sdk.NewCoins(sdk.NewCoin(box.TotalAmount.Token.Denom, box.Deposit.Interest.Token.Amount.Sub(totalInterest))))
		if err != nil {
			return err
		}
	}
	keeper.RemoveFromActiveBoxQueue(ctx, box.Deposit.MaturityTime, box.BoxId)
	box.BoxStatus = types.BoxFinished
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
