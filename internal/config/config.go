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
		SecretKey      string `mapstructure:"secret_key"`
		AdminAuth      struct {
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
		} `mapstructure:"admin_auth"`
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
