package types

import (
	"fmt"
)

type ExchangeParams struct {
	MaxOrdersPerAddr	uint64		`json:"max_orders_per_addr"`
}

func (ep ExchangeParams) String() string {
	return fmt.Sprintf(`Exchange Params:
  Max Orders Per Address:		%s`, ep.MaxOrdersPerAddr)
}