package clitest

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/hashgard/hashgard/x/box/utils"

	"github.com/hashgard/hashgard/x/box/params"

	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	boxtests "github.com/hashgard/hashgard/x/box/tests"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
	"github.com/stretchr/testify/require"
)

var (
	decimals        = uint(18)
	IssueCoinAmount = sdk.NewInt(10000)
	BoxAmount       = sdk.NewInt(1000)
	Interest        = sdk.NewInt(200)
	DefaultFlag     = "--gas 200000 -y"

	futureBoxReceivers = []string{keyBar, keyBaz, keyVesting}
)

func AddIssue(t *testing.T, f *Fixtures, sender sdk.AccAddress) string {
	// create issue
	f.TxIssueCreate(keyFoo, "foocoin", "FOO", uint64(IssueCoinAmount.Int64()), fmt.Sprintf("--decimals %d --gas 200000 -y", decimals))
	tests.WaitForNextNBlocksTM(1, f.Port)
	// Ensure transaction tags can be queried
	txs := f.QueryTxs(1, 50, "action:issue", fmt.Sprintf("sender:%s", sender))
	require.NotNil(t, txs)
	bytes, _ := hex.DecodeString(txs[len(txs)-1].Data)
	issueId := string(bytes[2:])
	return issueId
}
func CreateLockBox(t *testing.T, f *Fixtures, issueID string, sender sdk.AccAddress) (string, *params.BoxLockParams) {
	params := boxtests.GetLockBoxInfo()
	params.Sender = sender
	params.TotalAmount.Token.Denom = issueID
	params.TotalAmount.Token.Amount = BoxAmount
	f.TxLockBoxCreate(params, DefaultFlag)

	tests.WaitForNextNBlocksTM(1, f.Port)
	// Ensure transaction tags can be queried
	txs1 := f.QueryTxs(1, 50, "action:box_create_lock", fmt.Sprintf("sender:%s", sender))
	require.Len(t, txs1, 1)
	bytes, _ := hex.DecodeString(txs1[0].Data)
	boxId := string(bytes[2:])
	return boxId, params
}
func CreateDepositBox(t *testing.T, f *Fixtures, issueAID string, issueBID string, sender sdk.AccAddress) (string, *params.BoxDepositParams) {
	params := boxtests.GetDepositBoxInfo()
	params.Sender = sender
	params.TotalAmount.Token.Amount = BoxAmount
	params.TotalAmount.Token.Denom = issueAID
	params.Deposit.Interest.Token.Denom = issueBID
	params.Deposit.Interest.Token.Amount = Interest

	params.Deposit.BottomLine = BoxAmount.QuoRaw(4)
	params.Deposit.Price = BoxAmount.QuoRaw(10)

	f.TxDepositBoxCreate(params, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	// Ensure transaction tags can be queried
	txs1 := f.QueryTxs(1, 50, "action:box_create_deposit", fmt.Sprintf("sender:%s", sender))
	require.Len(t, txs1, 1)
	bytes, _ := hex.DecodeString(txs1[0].Data)
	boxId := string(bytes[2:])
	return boxId, params
}
func CreateFutureBox(t *testing.T, f *Fixtures, issueID string, sender sdk.AccAddress) (string, *params.BoxFutureParams) {
	params := boxtests.GetFutureBoxInfo()
	params.Sender = sender
	params.TotalAmount.Token.Amount = issueutils.QuoDecimals(params.TotalAmount.Token.Amount, decimals)
	params.TotalAmount.Token.Denom = issueID

	for i, items := range params.Future.Receivers {
		for j, v := range items {
			if j == 0 {
				params.Future.Receivers[i][j] = f.KeyAddress(futureBoxReceivers[i]).String()
				continue
			}
			amount, _ := sdk.NewIntFromString(v)
			params.Future.Receivers[i][j] = issueutils.QuoDecimals(amount, decimals).String()
		}
	}

	f.TxFutureBoxCreate(params, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	// Ensure transaction tags can be queried
	txs1 := f.QueryTxs(1, 50, "action:box_create_future", fmt.Sprintf("sender:%s", sender))
	require.Len(t, txs1, 1)
	bytes, _ := hex.DecodeString(txs1[0].Data)
	boxId := string(bytes[2:])
	return boxId, params
}
func getWaitBlocks(endTime int64) int64 {
	duration := endTime - time.Now().Unix()
	if duration < 0 {
		duration = 1
	}
	return duration + 2
}
func TestHashgardCLILockBox(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)
	// start hashgard server
	proc := f.HGStart()
	defer proc.Stop(false)
	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)
	//barAddr := f.KeyAddress(keyBar)
	issueID := AddIssue(t, f, fooAddr)
	boxID, params := CreateLockBox(t, f, issueID, fooAddr)

	fooAcc := f.QueryAccount(fooAddr)
	//require.Equal(t, fooAcc.GetCoins().AmountOf(issueID), IssueCoinAmount.Sub(BoxAmount))
	require.Equal(t, fooAcc.GetCoins().AmountOf(boxID), issueutils.MulDecimals(params.TotalAmount.Token.Amount, decimals))

	tests.WaitForNextNBlocksTM(getWaitBlocks(params.Lock.EndTime), f.Port)

	fooAcc = f.QueryAccount(fooAddr)
	//require.Equal(t, fooAcc.GetCoins().AmountOf(issueID), IssueCoinAmount)
	require.Equal(t, fooAcc.GetCoins().AmountOf(boxID), sdk.ZeroInt())
}

func TestHashgardCLIDepositBox(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)
	// start hashgard server
	proc := f.HGStart()
	defer proc.Stop(false)
	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)
	barAddr := f.KeyAddress(keyBar)
	//barAddr := f.KeyAddress(keyBar)
	issueAID := AddIssue(t, f, fooAddr)
	issueBID := AddIssue(t, f, fooAddr)

	boxID, params := CreateDepositBox(t, f, issueAID, issueBID, fooAddr)

	f.TxDepositBoxInterestInjection(keyFoo, boxID, params.Deposit.Interest.Token.Amount, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	// Ensure transaction tags can be queried
	txsInjection := f.QueryTxs(1, 50, "action:box_interest", "operation:injection", fmt.Sprintf("sender:%s", fooAddr))
	require.Len(t, txsInjection, 1)

	depositTo := BoxAmount.QuoRaw(int64(2))
	f.TxSend(keyFoo, barAddr, sdk.NewCoin(issueAID, depositTo), DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)

	tests.WaitForNextNBlocksTM(getWaitBlocks(params.Deposit.StartTime), f.Port)

	f.TxDepositTo(keyBar, boxID, depositTo, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	txsDeposit := f.QueryTxs(1, 50, "action:box_deposit", "operation:deposit-to", fmt.Sprintf("sender:%s", barAddr.String()))
	require.Len(t, txsDeposit, 1)

	fooAcc := f.QueryAccount(barAddr)
	require.Equal(t, fooAcc.GetCoins().AmountOf(boxID), depositTo.Quo(params.Deposit.Price))

	tests.WaitForNextNBlocksTM(getWaitBlocks(params.Deposit.MaturityTime), f.Port)

	f.TxWithdraw(keyBar, boxID, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	txsWithdraw := f.QueryTxs(1, 50, "action:box_withdraw", fmt.Sprintf("sender:%s", barAddr))
	require.Len(t, txsWithdraw, 1)

	fooAcc = f.QueryAccount(barAddr)
	require.Equal(t, fooAcc.GetCoins().AmountOf(boxID), sdk.ZeroInt())
}

func TestHashgardCLIFutureBox(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)
	// start hashgard server
	proc := f.HGStart()
	defer proc.Stop(false)
	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)
	//barAddr := f.KeyAddress(keyBar)
	issueID := AddIssue(t, f, fooAddr)
	boxID, params := CreateFutureBox(t, f, issueID, fooAddr)

	depositTo := params.TotalAmount.Token.Amount
	f.TxDepositTo(keyFoo, boxID, depositTo, DefaultFlag)

	tests.WaitForNextNBlocksTM(1, f.Port)
	txsDeposit := f.QueryTxs(1, 50, "action:box_deposit", "operation:deposit-to", fmt.Sprintf("sender:%s", fooAddr.String()))
	require.Len(t, txsDeposit, 1)

	tests.WaitForNextNBlocksTM(getWaitBlocks(params.Future.TimeLine[len(params.Future.TimeLine)-1]), f.Port)

	var address sdk.AccAddress
	for i, v := range params.Future.Receivers {
		address, _ = sdk.AccAddressFromBech32(v[0])
		account := f.QueryAccount(address)
		totalAmount := sdk.ZeroInt()
		for _, coin := range account.GetCoins() {
			if utils.IsBoxId(coin.Denom) {
				f.TxWithdraw(futureBoxReceivers[i], coin.Denom, DefaultFlag)
				tests.WaitForNextNBlocksTM(1, f.Port)
				totalAmount = totalAmount.Add(coin.Amount)
			}
		}
		account = f.QueryAccount(address)
		require.Equal(t, account.GetCoins().AmountOf(params.TotalAmount.Token.Denom), totalAmount)
	}
}
