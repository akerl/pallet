package dispatch

type Plugin struct {
	Name string
}

func GetPlugin(name string) Plugin, error {
	plugins, err := GetPlugins()
	if err != nil {
		return Plugin{}, err
	}
	for _, plugin := range plugins {
		if plugin.Name == name {
			return plugin, nil
		}
	}
	return Plugin{}, nil
}

func GetPlugins() []Plugin, error {
}

func InstallPlugin(name, url string) error {
}

func (p Plugin) Uninstall() error {
}

func IsPluginInstalled(name string) bool, error {
	plugin, err := GetPlugin(name)
	if err != nil {
		return false, err
	}
	if plugin.Name != "" {
		return true, nil
	}
	return false, nil
}
