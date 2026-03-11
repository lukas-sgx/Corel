package cmd

import (
	"github.com/lukas-sgx/corel/cmd/client"
	"github.com/lukas-sgx/corel/cmd/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "corel",
	Short: "Corel CLI tool",
	Long:  `Lightweight, cross-platform network relay and lateral movement agent for containerized environments.`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Cli() error {
	Init_templates(rootCmd)

	rootCmd.AddGroup(&cobra.Group{
		ID:    "modules",
		Title: "Available Modules:",
	})
	rootCmd.AddCommand(client.ClientCmd)
	rootCmd.AddCommand(server.ServerCmd)

	return rootCmd.Execute()
}