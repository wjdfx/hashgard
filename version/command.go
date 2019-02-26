package version

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagLong = "long"
)

var (
	// VersionCmd prints out the current sdk version
	VersionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the app version",
		Run:   printVersion,
	}
)

// return version of CLI/node and commit hash
func GetVersion() string {
	return Version
}

func GetCommit() string {
	return Commit
}

// CMD
func printVersion(cmd *cobra.Command, args []string) {
	fmt.Println("hashgard:", GetVersion())

	if viper.GetBool(flagLong) {
		fmt.Println("git commit:", GetCommit())
		fmt.Printf("go version %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	}
}

func init() {
	VersionCmd.Flags().Bool(flagLong, false, "Print long version information")
}