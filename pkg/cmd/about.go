package cmd

import (
	"fmt"
	"github.com/fatih/color"
	go_StringBuilder "github.com/golang-infrastructure/go-StringBuilder"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(aboutCmd)
}

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "About this tool github, about author, blabla.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		msg := go_StringBuilder.New().
			AppendString("\n\n\n").
			AppendString(fmt.Sprintf("%20s", "Repo")).AppendString(" : ").AppendString("https://github.com/cryptography-research-lab/jwt-secret-fake-random-goat\n\n").
			AppendString(fmt.Sprintf("%20s", "Problem feedback")).AppendString(" : ").AppendString("https://github.com/cryptography-research-lab/jwt-secret-fake-random-goat/issues\n\n").
			AppendString(fmt.Sprintf("%20s", "Author")).AppendString(" : ").AppendString("CC11001100\n").
			AppendString("\n\n\n").
			String()
		color.HiGreen(msg)
	},
}
