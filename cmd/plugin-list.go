package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func pluginListRunner(cmd *cobra.Command, args []string) error {
	fmt.Println("Hi")
	return nil
}

var pluginListCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed language plugins",
	RunE:  pluginListRunner,
}

func init() {
	pluginCmd.AddCommand(pluginListCmd)
}
