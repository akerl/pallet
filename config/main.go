package config

import (
	"github.com/spf13/viper"
)

const (
	configName = "config"
	configDir  = "$HOME/.pallet"
)

var defaults = map[string]interface{}{
	"root": configDir,
}

// LoadConfig loads the Pallet configuration
func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.AutomaticEnv()
	v.SetEnvPrefix("pallet")

	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.AddConfigPath(configDir)
	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return v, nil
		}
		return nil, err
	}
	return v, nil
}
