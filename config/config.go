package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
}

func LoadConfig() (config Config, err error) {
    viper.AddConfigPath("$CONFIG_PATH")
    viper.SetConfigName("config")
    viper.SetConfigType("env")
    viper.AutomaticEnv()

    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    err = viper.Unmarshal(&config)
    return
}
