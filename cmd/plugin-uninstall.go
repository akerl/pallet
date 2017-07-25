package cmd

import (
	"fmt"

	"github.com/akerl/pallet/dispatch"

	"github.com/spf13/cobra"
)

func pluginUninstallRunner(cmd *cobra.Command, args []string) error {
	pluginName := args[0]
	installed, err := dispatch.IsPluginInstalled(pluginName)
	if err != nil {
		return err
	}
	if !installed {
		return fmt.Errorf("Plugin not installed")
	}
	plugin := dispatch.GetPlugin(pluginName)
	return plugin.Uninstall()
}

var pluginUninstallCmd = &cobra.Command{
	Use:   "uninstall PLUGIN",
	Short: "Uninstall a new plugin",
	RunE:  pluginUninstallRunner,
	Args:  cobra.ExactArgs(1),
}

func init() {
	pluginCmd.AddCommand(pluginUninstallCmd)
}
