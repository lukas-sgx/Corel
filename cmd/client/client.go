package client

import (
	"github.com/spf13/cobra"
)

var ClientCmd = &cobra.Command{
	Use:     "client",
	Short:   "Start the Corel client",
	Long:    `Start the Corel client, it will connect to the specified server and handle incoming requests.`,
	GroupID: "modules",
	Run: func(cmd *cobra.Command, args []string) {
		println("Client started")
	},
}

func Client() error {
	return ClientCmd.Execute()
}
