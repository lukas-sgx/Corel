package cmd

import (
	"github.com/lukas-sgx/corel/cmd/agent"
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
	agent.InitFlags(rootCmd.Version)
	server.InitFlags()
	rootCmd.AddCommand(agent.ClientCmd)
	rootCmd.AddCommand(server.ServerCmd)

	return rootCmd.Execute()
}