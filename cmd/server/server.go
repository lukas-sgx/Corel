package server

import (
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start the Corel server",
	Long:    `Start the Corel server, it will listen on the specified port and handle incoming requests.`,
	GroupID: "modules",
	Run: func(cmd *cobra.Command, args []string) {
		println("Server started")
	},
}

func Server() error {
	return ServerCmd.Execute()
}
