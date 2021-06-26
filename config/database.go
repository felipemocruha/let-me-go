package config

type DatabaseConfig struct {
	Host string `envconfig:"DB_HOST"`
	Port string `envconfig:"DB_PORT"`
	Name string `envconfig:"DB_NAME"`
	User string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
}
