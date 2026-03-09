package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "corel",
	Short: "Corel is a tool to manage your projects",
	Long:  `Corel is a tool to manage your projects, it allows you to create, delete, list and update your projects.`,
}

func Cli() error {
	return rootCmd.Execute()
}