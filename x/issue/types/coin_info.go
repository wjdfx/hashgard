package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//Issue interface
type Issue interface {
	GetIssueId() string
	SetIssueId(string)

	GetIssuer() sdk.AccAddress
	SetIssuer(sdk.AccAddress)

	GetOwner() sdk.AccAddress
	SetOwner(sdk.AccAddress)

	GetIssueTime() time.Time
	SetIssueTime(time.Time)

	GetName() string
	SetName(string)

	GetTotalSupply() sdk.Int
	SetTotalSupply(sdk.Int)

	GetDecimals() uint
	SetDecimals(uint)

	GetDescription() string
	SetDescription(string)

	GetBurningFinished() bool
	SetBurningFinished(bool)

	GetMintingFinished() bool
	SetMintingFinished(bool)

	GetSymbol() string
	SetSymbol(string)

	String() string
}

// CoinIssues is an array of Issue
type CoinIssues []CoinIssueInfo

//Coin Issue Info
type CoinIssueInfo struct {
	IssueId         string         `json:"issue_id"`
	Issuer          sdk.AccAddress `json:"issuer"`
	Owner           sdk.AccAddress `json:"owner"`
	IssueTime       time.Time      `json:"issue_time"`
	Name            string         `json:"name"`
	Symbol          string         `json:"symbol"`
	TotalSupply     sdk.Int        `json:"total_supply"`
	Decimals        uint           `json:"decimals"`
	Description     string         `json:"description"`
	BurningFinished bool           `json:"burning_finished"`
	MintingFinished bool           `json:"minting_finished"`
}

// Implements Issue Interface
var _ Issue = (*CoinIssueInfo)(nil)

//nolint
func (ci CoinIssueInfo) GetIssueId() string {
	return ci.IssueId
}
func (ci *CoinIssueInfo) SetIssueId(issueId string) {
	ci.IssueId = issueId
}
func (ci CoinIssueInfo) GetIssuer() sdk.AccAddress {
	return ci.Issuer
}
func (ci *CoinIssueInfo) SetIssuer(issuer sdk.AccAddress) {
	ci.Issuer = issuer
}
func (ci CoinIssueInfo) GetOwner() sdk.AccAddress {
	return ci.Owner
}
func (ci *CoinIssueInfo) SetOwner(owner sdk.AccAddress) {
	ci.Owner = owner
}
func (ci CoinIssueInfo) GetIssueTime() time.Time {
	return ci.IssueTime
}
func (ci *CoinIssueInfo) SetIssueTime(issueTime time.Time) {
	ci.IssueTime = issueTime
}
func (ci CoinIssueInfo) GetName() string {
	return ci.Name
}
func (ci *CoinIssueInfo) SetName(name string) {
	ci.Name = name
}
func (ci CoinIssueInfo) GetTotalSupply() sdk.Int {
	return ci.TotalSupply
}
func (ci *CoinIssueInfo) SetTotalSupply(totalSupply sdk.Int) {
	ci.TotalSupply = totalSupply
}
func (ci CoinIssueInfo) GetDecimals() uint {
	return ci.Decimals
}
func (ci *CoinIssueInfo) SetDecimals(decimals uint) {
	ci.Decimals = decimals
}
func (ci CoinIssueInfo) GetDescription() string {
	return ci.Description
}
func (ci *CoinIssueInfo) SetDescription(description string) {
	ci.Description = description
}
func (ci CoinIssueInfo) GetMintingFinished() bool {
	return ci.MintingFinished
}
func (ci *CoinIssueInfo) SetMintingFinished(mintingFinished bool) {
	ci.MintingFinished = mintingFinished
}
func (ci CoinIssueInfo) GetBurningFinished() bool {
	return ci.BurningFinished
}
func (ci *CoinIssueInfo) SetBurningFinished(burningFinished bool) {
	ci.BurningFinished = burningFinished
}
func (ci CoinIssueInfo) GetSymbol() string {
	return ci.Symbol
}
func (ci *CoinIssueInfo) SetSymbol(symbol string) {
	ci.Symbol = symbol
}

//nolint
func (ci CoinIssueInfo) String() string {
	return fmt.Sprintf(`Issue:
  IssueId:          %s
  Issuer:           %s
  Owner:            %s
  Name:             %s
  Symbol:    	    %s
  TotalSupply:      %s
  Decimals:         %d
  Description	    %s
  BurningFinished   %t 
  MintingFinished:  %t `,
		ci.IssueId, ci.Issuer.String(), ci.Owner.String(), ci.Name, ci.Symbol, ci.TotalSupply.String(),
		ci.Decimals, ci.Description, ci.BurningFinished, ci.MintingFinished)
}

//nolint
func (coinIssues CoinIssues) String() string {
	out := fmt.Sprintf("%-15s|%-10s|%-6s|%-18s|%-8s|%-15s|%s\n",
		"IssueID", "Name", "Symbol", "TotalSupply", "Decimals", "MintingFinished", "IssueTime")
	for _, issue := range coinIssues {
		out += fmt.Sprintf("%-15s|%-10s|%-6s|%-18s|%-8d|%-15t|%s\n",
			issue.IssueId, issue.Name, issue.Symbol, issue.TotalSupply.String(), issue.Decimals, issue.MintingFinished, issue.IssueTime.String())
	}
	return strings.TrimSpace(out)
}
