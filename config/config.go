package config

import (
	"os"
	
	"github.com/kelseyhightower/envconfig"
	"github.com/joho/godotenv"
)

type Config struct {
	Database *DatabaseConfig
}

func LoadConfig() (config Config, err error) {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "dev" {
		err = godotenv.Load(".env")
		if err != nil {
			return
		}
	}
	
	err = envconfig.Process("app", &config)
	if err != nil {
		return
	}

	return
}
