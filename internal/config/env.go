package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PAT string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
		PAT: os.Getenv("PAT"),
	}, nil
}
