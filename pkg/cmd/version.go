package cmd

// 展示版本信息

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information.",
	Long:  `Show version information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current version: v0.1")
	},
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
