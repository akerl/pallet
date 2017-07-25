package cmd

import (
	"fmt"

	"github.com/akerl/pallet/dispatch"

	"github.com/spf13/cobra"
)

func pluginInstallRunner(cmd *cobra.Command, args []string) error {
	pluginName := args[0]
	pluginURL := args[1]
	plugins, err := dispatch.Plugins()
	if err != nil {
		return err
	}
	for _, plugin := range plugins {
		if pluginName == plugin.Name {
			return fmt.Errorf("Plugin already installed")
		}
	}
	return dispatch.InstallPlugin(pluginName, pluginURL)
}

var pluginInstallCmd = &cobra.Command{
	Use:   "install PLUGIN URL",
	Short: "Install a new plugin",
	RunE:  pluginInstallRunner,
	Args: cobra.ExactArgs(2),
}

func init() {
	pluginCmd.AddCommand(pluginInstallCmd)
}
