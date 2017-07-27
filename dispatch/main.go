package dispatch

type PluginSet map[string]Plugin

type Plugin struct {
	Name, URL string
}

func LoadPluginSet() (PluginSet, error) {
	ps := PluginSet{}
	err := ps.Load()
	return ps, err
}

func (ps *PluginSet) load() error {
	return nil
}

func (p Plugin) Install() error {
}

func (p Plugin) Uninstall() error {
}

func (p Plugin) Upgrade() error {
}
