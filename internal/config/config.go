package config

import (
	"encoding/json"
	"os"
)

// BisnodeConfig holds configuration for the Bisnode API
type BisnodeConfig struct {
	BaseURL      string `json:"base_url"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type Config struct {
	Bisnode BisnodeConfig `json:"bisnode"`
}

// Load loads configuration from config.json
func Load() (*Config, error) {
	// In production, you might want to use environment variables or a config management system
	// For simplicity, we'll use a config file
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
