package cmd

import (
	"github.com/spf13/cobra"
)

var Template = `Usage:
  {{.UseLine}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [module]{{end}}

Available Modules:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}

Use "{{.CommandPath}} [command] --help" for more information about a command.
`

var VersionTemplate = "Corel version {{.Version}}\n"

func Init_templates(rootCmd *cobra.Command) {
	rootCmd.SetHelpCommand(&cobra.Command{Use: "no-help", Hidden: true})
	rootCmd.SetUsageTemplate(Template)

	rootCmd.Version = "0.1.0"
	rootCmd.SetVersionTemplate(VersionTemplate)
}