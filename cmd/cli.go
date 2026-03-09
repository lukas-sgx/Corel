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

func init_templates() {
	rootCmd.SetHelpCommand(&cobra.Command{Use: "no-help", Hidden: true})
	rootCmd.SetUsageTemplate(`Usage:
  {{.UseLine}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [module]{{end}}

Available Modules:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}

Use "{{.CommandPath}} [command] --help" for more information about a command.
`)

	rootCmd.Version = "0.1.0"
	rootCmd.SetVersionTemplate("Corel version {{.Version}}\n")
}

func Cli() error {
	init_templates()
	rootCmd.AddGroup(&cobra.Group{
		ID:    "modules",
		Title: "Available Modules:",
	})
	rootCmd.AddCommand(client.ClientCmd)
	rootCmd.AddCommand(server.ServerCmd)
	return rootCmd.Execute()
}