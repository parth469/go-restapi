package utils

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("failed to load .env file")

	}

	config := &Config{
		App:      os.Getenv("APP_NAME"),
		Port:     os.Getenv("PORT"),
		Database: os.Getenv("DATABASE"),
	}

	if config.App == "" || config.Port == "" || config.Database == "" {
		return nil, errors.New("required environment variables missing")
	}
	return config, nil
}

type Config struct {
	Port     string
	App      string
	Database string
}
