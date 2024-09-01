package cmd

// 展示版本信息

import (
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/web"
	"github.com/spf13/cobra"
)

var port uint = 10086

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the Web server in the test range.",
	Long:  `Start the Web server in the test range.`,
	Run: func(cmd *cobra.Command, args []string) {
		web.Run(port)
	},
	TraverseChildren: true,
}

func init() {
	serverCmd.Flags().UintVarP(&port, "port", "p", 10086, "Specify the port on which to start the Web service, with the default port being 10086.\n")
	rootCmd.AddCommand(serverCmd)
}
