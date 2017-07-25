package dispatch

type Plugin struct {
	Name string
}

func GetPlugin(name string) Plugin, error {

}

func GetPlugins() []Plugin, error {
}

func InstallPlugin(name, url string) error {
}

func (p Plugin) Uninstall() error {
}

func IsPluginInstalled(name string) bool, error {
	plugins, err := GetPlugins()
	if err != nil {
		return false, err
	}
	for _, plugin := range plugins {
		if plugin.Name == name {
			return true, nil
		}
	}
	return false, nil
}
