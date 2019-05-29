package mint

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// mint parameters
type Params struct {
	Inflation		sdk.Dec `json:"inflation"` // inflation rate
	InflationBase	sdk.Int `json:"inflation_base"`
}

func NewParams(inflation sdk.Dec, inflationBase sdk.Int) Params {
	return Params{
		Inflation:		inflation,
		InflationBase:	inflationBase,
	}
}

// default minting module parameters
func DefaultParams() Params {
	return Params{
		Inflation: 		sdk.NewDecWithPrec(8, 2),
		InflationBase:  sdk.NewIntWithDecimal(1, 11).Mul(sdk.NewIntWithDecimal(1, 18)), // 1*(10^11)gard, 1*(10^11)*(10^18)agard
	}
}

func validateParams(params Params) error {
	if params.Inflation.LT(sdk.ZeroDec()) || params.Inflation.GT(sdk.OneDec()) {
		return fmt.Errorf("minter inflation (%s) should between 0 and 1", params.Inflation.String())
	}

	if !params.InflationBase.GT(sdk.ZeroInt()) {
		return fmt.Errorf("minter inflation basement (%s) should be positive", params.InflationBase.String())
	}
	return nil
}

func (p Params) String() string {
	return fmt.Sprintf(`Minting Params:
  Inflation :			%s
  Inflation Base :      %s
`,
		p.Inflation, p.InflationBase,
	)
}
