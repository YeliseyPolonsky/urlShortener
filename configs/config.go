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
		log.Println("Warning: error loading .env file, using system enviroment variables!")
	}
	config := &Config{
		DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		AuthConfig{
			Secret: os.Getenv("SECRET"),
		},
	}

	if config.Db.Dsn == "" || config.Auth.Secret == "" {
		log.Fatal("Some variablies of Config is empty!")
	}

	return config
}
