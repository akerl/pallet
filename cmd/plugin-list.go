package cmd

import (
	"fmt"

	"github.com/akerl/pallet/dispatch"

	"github.com/spf13/cobra"
)

func pluginListRunner(cmd *cobra.Command, args []string) error {
	plugins, err := dispatch.Plugins()
	if err != nil {
		return err
	}
	for _, plugin := range plugins {
		fmt.Printf("%s\n", plugin.Name)
	}
	return nil
}

var pluginListCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed language plugins",
	RunE:  pluginListRunner,
	Args: cobra.NoArgs(),
}

func init() {
	pluginCmd.AddCommand(pluginListCmd)
}
