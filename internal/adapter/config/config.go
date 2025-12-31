package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server Server
	}

	Server struct {
		Port int
	}
)

var (
	configInstance Config
	once           sync.Once
)

func NewConfig() *Config {
	once.Do(func() {
		viper.SetConfigFile("config.yml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/srv/EchoNotes/")
		viper.AddConfigPath("/app/")

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Sprintf("failed to read config: %v", err))
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(fmt.Sprintf("failed to unmarshal config: %v", err))
		}
	})

	return &configInstance
}
