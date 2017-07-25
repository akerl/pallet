package cmd

import (
	"fmt"

	"github.com/akerl/pallet/dispatch"

	"github.com/spf13/cobra"
)

func pluginUninstallRunner(cmd *cobra.Command, args []string) error {
	pluginName := args[0]
	found, err := dispatch.IsPluginInstalled(pluginName)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("Plugin not installed")
	}

	dispatch.InstallPlugin(pluginName, pluginURL)
	return nil
}

var pluginUninstallCmd = &cobra.Command{
	Use:   "uninstall PLUGIN",
	Short: "Uninstall a new plugin",
	RunE:  pluginUninstallRunner,
	Args: cobra.ExactArgs(1),
}

func init() {
	pluginCmd.AddCommand(pluginUninstallCmd)
}
