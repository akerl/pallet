package dispatch

import (
	"github.com/akerl/pallet/config"

	"github.com/spf13/viper"
)

var c *viper.Viper

func LoadConfig() error {
	var err error
	c, err = config.LoadConfig()
	return err
}

type PluginSet map[string]Plugin

type Plugin struct {
	Name, URL string
}

func LoadPluginSet() (PluginSet, error) {
	ps := PluginSet{}
	err := ps.load()
	return ps, err
}

func (ps *PluginSet) load() error {
	return nil
}

func (p Plugin) Install() error {
	return nil
}

func (p Plugin) Uninstall() error {
	return nil
}

func (p Plugin) Upgrade() error {
	return nil
}
