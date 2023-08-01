package cmd

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var Version string = ""

func getVersion(version string) string {
	return fmt.Sprintf(`tagli

Version: %s
OS: %s
Arch: %s`, version, runtime.GOOS, runtime.GOARCH)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			Version = "unknown"
		}
		Version = info.Main.Version
		fmt.Println(getVersion(Version))
	},
	Short: "Show cli version info",
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
