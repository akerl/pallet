package config

import (
	"github.com/spf13/viper"
)

const (
	configName = "config"
	configDir  = "$HOME/.pallet"
)

var defaults = map[string]interface{}{
	"root": "$HOME/.pallet",
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.AutomaticEnv()
	v.SetEnvPrefix("pallet")

	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.AddConfigPath(configDir)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
}
