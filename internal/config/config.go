package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DatabaseURI string `mapstructure:"database_uri"`

    Server struct {
        Port           int    `mapstructure:"port"`
        LogLevel       string `mapstructure:"log_level"`
    } `mapstructure:"server"`
}

var App AppConfig

func LoadConfig(path string) {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	if err := viper.Unmarshal(&App); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

}
