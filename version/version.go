package version

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

// Version - Hashgard Version
var Version = ""

// return version of CLI/node and commit hash
func GetVersion() string {
	return Version
}

// ServeVersionCommand
func ServeVersionCommand(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show executable binary version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(GetVersion())
			return nil
		},
	}
	return cmd
}
