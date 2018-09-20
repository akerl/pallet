package dispatch

import (
	"github.com/akerl/pallet/config"

	"github.com/spf13/viper"
)

var c *viper.Viper

// LoadConfig loads the viper config
func LoadConfig() error {
	var err error
	c, err = config.LoadConfig()
	return err
}

// PluginSet is a group of plugins
type PluginSet map[string]Plugin

// Plugin defines a repo for configuring a language
type Plugin struct {
	Name, URL string
}

// LoadPluginSet generates a set of Plugins
func LoadPluginSet() (PluginSet, error) {
	ps := PluginSet{}
	err := ps.load()
	return ps, err
}

func (ps *PluginSet) load() error {
	return nil
}

// Install installs the plugin
func (p Plugin) Install() error {
	return nil
}

// Uninstall removes the plugin
func (p Plugin) Uninstall() error {
	return nil
}

// Upgrade upgrades the plugin
func (p Plugin) Upgrade() error {
	return nil
}
