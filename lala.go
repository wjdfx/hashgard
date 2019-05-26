package main

import (
	"fmt"
	"encoding/json"
)

type UsageType byte

const (
	UsageTypeBurn       UsageType = 0x01
	UsageTypeDistribute UsageType = 0x02
	UsageTypeGrant      UsageType = 0x03
)

// Turns VoteOption byte to String
func (ut UsageType) String() string {
	switch ut {
	case UsageTypeBurn:
		return "Burn"
	case UsageTypeDistribute:
		return "Distribute"
	case UsageTypeGrant:
		return "Grant"
	default:
		return ""
	}
}

// Marshal needed for protobuf compatibility
func (ut UsageType) Marshal() ([]byte, error) {
	return []byte{byte(ut)}, nil
}

// Marshals to JSON using string
func (ut UsageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ut.String())
}

func main() {
	xixi := UsageTypeBurn
	fmt.Println(xixi)

	fmt.Printf("%T\n", xixi)

	fmt.Println("----------")

	guagua, _ := json.Marshal(xixi)
	fmt.Println(guagua)
	fmt.Printf("%T", guagua)

	fmt.Println("----------")
	dada, _ := json.Marshal(xixi.String())
	fmt.Println(dada)
	fmt.Printf("%T", dada)
}
