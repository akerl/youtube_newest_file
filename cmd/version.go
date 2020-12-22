package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/akerl/slack_newest_file/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of slack_newest_file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
