package config

import (
	"github.com/spf13/viper"
	"userLoggingProject/pkg/logger"
)

type LaunchConfig struct {
	Port       int
	Address    string
	LaunchType string
}

func Read() (LaunchConfig, error) {
	viper.AddConfigPath("/config")
	viper.AddConfigPath("/")
	viper.SetDefault("LaunchType", "rest")
	viper.SetDefault("Port", 8080)
	viper.SetDefault("Address", "localhost")
	if err := viper.ReadInConfig(); err != nil {
		logger.Info("Config file not found")
		logger.Info(err)
	}

	return LaunchConfig{Port: viper.GetInt("Port"), Address: viper.GetString("Address"), LaunchType: viper.GetString("LaunchType")}, nil
}
