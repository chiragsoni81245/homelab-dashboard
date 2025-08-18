package config

import (
	"homelab-dashboard/internal/logger"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DatabaseURI string `mapstructure:"database_uri"`

    Server struct {
        Port           int    `mapstructure:"port"`
        LogLevel       string `mapstructure:"log_level"`
		SecretKey      string `mapstructure:"secret_key"`
		UpdateFrequency int   `mapstructure:"update_frequency"` // This will be in milliseconds
		AdminAuth      struct {
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
		} `mapstructure:"admin_auth"`
    } `mapstructure:"server"`
}

var App AppConfig

func validateConfig() {
    missing := []string{}

    if App.DatabaseURI == "" {
        missing = append(missing, "APP_DATABASE_URI / database_uri")
    }
    if App.Server.Port == 0 {
        missing = append(missing, "APP_SERVER_PORT / server.port")
    }
    if App.Server.SecretKey == "" {
        missing = append(missing, "APP_SERVER_SECRET_KEY / server.secret_key")
    }
    if App.Server.AdminAuth.Username == "" {
        missing = append(missing, "APP_SERVER_ADMIN_AUTH_USERNAME / server.admin_auth.username")
    }
    if App.Server.AdminAuth.Password == "" {
        missing = append(missing, "APP_SERVER_ADMIN_AUTH_PASSWORD / server.admin_auth.password")
    }
	if App.Server.UpdateFrequency == 0 {
		App.Server.UpdateFrequency = 2000
	}

    if len(missing) > 0 {
        log.Fatalf("Missing required configuration: %v", missing)
    }
}

func LoadConfig(path string) {
	viper.SetConfigFile(path)

	// Read from ENV
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP") // optional prefix, e.g. APP_DATABASE_URI
	viper.BindEnv("database_uri") 
	viper.BindEnv("server.port") 
    viper.BindEnv("server.log_level") 
    viper.BindEnv("server.secret_key") 
    viper.BindEnv("server.admin_auth.username") 
    viper.BindEnv("server.admin_auth.password")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Infof("Config file not found: %v (falling back to ENV only)", err)
	} else {
		logger.Log.Info("Config file loaded successfully")
	}

	if err := viper.Unmarshal(&App); err != nil {
		logger.Log.Fatalf("Failed to unmarshal config: %v", err)
	}

	validateConfig()
}
