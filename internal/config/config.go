package config

import (
	"os"
)

// Config holds application configuration.
type Config struct {
	Port    string
	AppName string
	BaseURL string
	DataDir string
}

// Load returns configuration from environment variables with defaults.
func Load() *Config {
	return &Config{
		Port:    getEnv("PORT", "8080"),
		AppName: getEnv("APP_NAME", "RHYRED"),
		BaseURL: getEnv("BASE_URL", "https://rhyred.net"),
		DataDir: getEnv("DATA_DIR", "web/data"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
