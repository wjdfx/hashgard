package clitest

import (
	"encoding/hex"
	"fmt"
	"testing"

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
	txs1 := f.QueryTxs(1, 50, "action:box_create", "box-type:lock", fmt.Sprintf("sender:%s", sender))
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
	txs1 := f.QueryTxs(1, 50, "action:box_create", "box-type:deposit", fmt.Sprintf("sender:%s", sender))
	require.Len(t, txs1, 1)
	bytes, _ := hex.DecodeString(txs1[0].Data)
	boxId := string(bytes[2:])
	return boxId, params
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
	boxID, _ := CreateLockBox(t, f, issueID, fooAddr)

	fooAcc := f.QueryAccount(fooAddr)
	//require.Equal(t, fooAcc.GetCoins().AmountOf(issueID), IssueCoinAmount.Sub(BoxAmount))
	require.Equal(t, fooAcc.GetCoins().AmountOf(boxID), issueutils.MulDecimals(BoxAmount, decimals))

	tests.WaitForNextNBlocksTM(8, f.Port)

	fooAcc = f.QueryAccount(fooAddr)
	//require.Equal(t, fooAcc.GetCoins().AmountOf(issueID), IssueCoinAmount)
	require.Equal(t, fooAcc.GetCoins().AmountOf(boxID), sdk.ZeroInt())
}

func TestHashgardCLILockDepositBox(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)
	// start hashgard server
	proc := f.HGStart()
	defer proc.Stop(false)
	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)
	//barAddr := f.KeyAddress(keyBar)
	issueAID := AddIssue(t, f, fooAddr)
	issueBID := AddIssue(t, f, fooAddr)
	fmt.Println(issueAID, issueBID)
	boxID, params := CreateDepositBox(t, f, issueAID, issueBID, fooAddr)

	f.TxDepositBoxInterestInjection(keyFoo, boxID, params.Deposit.Interest.Token.Amount, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	// Ensure transaction tags can be queried
	txs1 := f.QueryTxs(1, 50, "action:box_interest", "operation:injection", fmt.Sprintf("sender:%s", fooAddr))
	require.Len(t, txs1, 1)
}
