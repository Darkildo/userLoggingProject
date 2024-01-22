package config

import "github.com/spf13/viper"

type ConfigSystem struct {
	Port    int
	Address string
}

func Read(launchType string) (ConfigSystem, error) {
	viper.AddConfigPath("/config")
	viper.AddConfigPath("/")

	if err := viper.ReadInConfig(); err != nil {
		return ConfigSystem{}, err
	}

	return ConfigSystem{Port: viper.GetInt("Port"), Address: viper.GetString("Address")}, nil
}
