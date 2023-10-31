package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	BuildTime  string
	CommitHash string
	Version    string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints out version, commit hash and build time of the binary.",
	Long:  `Prints out version, commit hash and build time of the binary.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("#### Godo ####")
		fmt.Println("Version:\t", Version)
		fmt.Println("Build time:\t", BuildTime)
		fmt.Println("Commit hash:\t", CommitHash)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
