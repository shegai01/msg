package config

import (
	"log/slog"
	"os"
)

type Config struct {
	HTTPPort string
	DB_port  string
	DB_host  string
	DB_Port  string
	DB_user  string
	DB_pass  string
	DB_Name  string
}

func Load() *Config {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		slog.Info("http port is required")
	}
	return &Config{
		HTTPPort: port,
		DB_host:  os.Getenv("DBHost"),
		DB_port:  os.Getenv("DBPort"),
		DB_user:  os.Getenv("DBuser"),
		DB_pass:  os.Getenv("DBpass"),
		DB_Name:  os.Getenv("DBName"),
	}
}
