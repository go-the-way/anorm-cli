package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "anorm-cli",
	Short: "An anorm Code generator written in Go",
	Long:  "An anorm Code generator written in Go",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
