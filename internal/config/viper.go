package config

import (
	"fmt"
	"myapp/internal/model"

	"github.com/spf13/viper"
)

func NewViper() *model.Config {
	var config *model.Config
	viper := viper.New()

	// config.SetConfigName("env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./../")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}

	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshal config file: %w", err))
	}

	return config
}
