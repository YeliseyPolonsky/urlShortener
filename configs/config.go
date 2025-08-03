package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	if godotenv.Load() != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &Config{
		DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		AuthConfig{
			Secret: os.Getenv("SECRET"),
		},
	}
}
