package domain

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Issue interface {
	GetIssueId() string
	SetIssueId(string)

	GetIssuer() sdk.AccAddress
	SetIssuer(sdk.AccAddress)

	GetName() string
	SetName(string)

	GetTotalSupply() sdk.Int
	SetTotalSupply(sdk.Int)

	GetDecimals() uint
	SetDecimals(uint)

	GetMintingFinished() bool
	SetMintingFinished(bool)

	GetSymbol() string
	SetSymbol(string)

	String() string
}
type CoinIssueInfo struct {
	IssueId         string         `json:"issue_id"`
	Issuer          sdk.AccAddress `json:"issuer"`
	Name            string         `json:"name"`
	Symbol          string         `json:"symbol"`
	TotalSupply     sdk.Int        `json:"total_supply"`
	Decimals        uint           `json:"decimals"`
	MintingFinished bool           `json:"minting_finished"`
}

// Implements Issue Interface
var _ Issue = (*CoinIssueInfo)(nil)

func (ci CoinIssueInfo) GetIssueId() string {
	return ci.IssueId
}
func (ci CoinIssueInfo) SetIssueId(issueId string) {
	ci.IssueId = issueId
}
func (ci CoinIssueInfo) GetIssuer() sdk.AccAddress {
	return ci.Issuer
}
func (ci CoinIssueInfo) SetIssuer(issuer sdk.AccAddress) {
	ci.Issuer = issuer
}

func (ci CoinIssueInfo) GetName() string {
	return ci.Name
}
func (ci CoinIssueInfo) SetName(name string) {
	ci.Name = name
}

func (ci CoinIssueInfo) GetTotalSupply() sdk.Int {
	return ci.TotalSupply
}
func (ci CoinIssueInfo) SetTotalSupply(totalSupply sdk.Int) {
	ci.TotalSupply = totalSupply
}

func (ci CoinIssueInfo) GetDecimals() uint {
	return ci.Decimals
}
func (ci CoinIssueInfo) SetDecimals(decimals uint) {
	ci.Decimals = decimals
}

func (ci CoinIssueInfo) GetMintingFinished() bool {
	return ci.MintingFinished
}
func (ci CoinIssueInfo) SetMintingFinished(mintingFinished bool) {
	ci.MintingFinished = mintingFinished
}

func (ci CoinIssueInfo) GetSymbol() string {
	return ci.Symbol
}
func (ci CoinIssueInfo) SetSymbol(symbol string) {
	ci.Symbol = symbol
}

//TODO
func (ci CoinIssueInfo) String() string {
	return fmt.Sprintf(`Issue:
  IssueId:          %s
  Issuer:           %s
  Name:             %s
  Symbol:    	    %s
  TotalSupply:      %s
  Decimals:         %d
  MintingFinished:  %t `,
		ci.IssueId, ci.Issuer.String(), ci.Name, ci.Symbol, ci.TotalSupply.String(),
		ci.Decimals, ci.MintingFinished)
}

type CoinIssue struct {
	*sdk.Coin
	IssueId string         `json:"issue_id"`
	Issuer  sdk.AccAddress `json:"issuer"`
}
