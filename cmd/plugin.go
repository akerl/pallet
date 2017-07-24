package cmd

import (
	"github.com/spf13/cobra"
)

func pluginRunner(cmd *cobra.Command, args []string) error {
	return nil
}

var pluginCmd = &cobra.Command{
	Use:   "plugin SUBCOMMAND",
	Short: "Manage installed language plugins",
}

func init() {
	rootCmd.AddCommand(pluginCmd)
}
