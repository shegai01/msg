package config

import (
	"log/slog"
	"os"
)

type Config struct {
	HTTPPort string
}

func Load() *Config {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		slog.Info("http port is required")
	}
	return &Config{
		HTTPPort: port,
	}
}
