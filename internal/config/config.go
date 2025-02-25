// internal/config/config.go
package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// Config holds application configuration
type Config struct {
	OpenAIAPIKey string
	Model        string
}

// LoadConfig reads configuration from disk
func LoadConfig() (*Config, error) {
	config := &Config{
		Model: "gpt-4o",
	}

	configDir, err := getConfigDir()
	if err != nil {
		return nil, err
	}

	configFile := filepath.Join(configDir, "config.txt")
	if _, err := os.Stat(configFile); err == nil {
		data, err := os.ReadFile(configFile)
		if err != nil {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
		config.OpenAIAPIKey = string(data)
	}

	return config, nil
}

// SaveAPIKey saves the OpenAI API key to disk
func SaveAPIKey(apiKey string) error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	configFile := filepath.Join(configDir, "config.txt")
	return os.WriteFile(configFile, []byte(apiKey), 0600)
}

// getConfigDir returns the path to the config directory
func getConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not find home directory: %w", err)
	}

	configDir := filepath.Join(home, ".config", "git-commit-ai")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", fmt.Errorf("could not create config directory: %w", err)
	}

	return configDir, nil
}
