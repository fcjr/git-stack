package cmd

import (
	"fmt"

	"github.com/fcjr/git-stack/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Args:  cobra.NoArgs,
	Short: "Get the version of git-stack",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.String())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
