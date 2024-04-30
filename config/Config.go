package config

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	SvcDBGo      string `env:"SERVICE_DB_GO"`
	SvcDBSekolah string `env:"SERVICE_DB_SEKOLAH"`
}

var Cfg *Config

func Init() {
	var err error
	Cfg, err = LoadConfig()
	if err != nil {
		// Handle error, potentially exit application
	}
}

func LoadConfig() (*Config, error) {
	// Load environment variables from .env file (optional)
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	// Create a new Config instance
	cfg := &Config{}

	cfg.SvcDBGo = os.Getenv("SERVICE_DB_GO")
	cfg.SvcDBSekolah = os.Getenv("SERVICE_DB_SEKOLAH")

	// Validate configuration values (optional)
	// ...

	return cfg, nil
}
