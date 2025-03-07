// internal/config/config.go
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type AIProvider string

const (
	ProviderOpenAI  AIProvider = "openai"
	ProviderClaude  AIProvider = "claude"
	DefaultProvider AIProvider = ProviderOpenAI
)

// Config holds application configuration
type Config struct {
	ActiveProvider AIProvider        `json:"active_provider"`
	APIKeys        map[string]string `json:"api_keys"`
	Models         map[string]string `json:"models"`
}

// GetDefaultConfig returns a config with default values
func GetDefaultConfig() *Config {
	return &Config{
		ActiveProvider: DefaultProvider,
		APIKeys: map[string]string{
			string(ProviderOpenAI): "",
			string(ProviderClaude): "",
		},
		Models: map[string]string{
			string(ProviderOpenAI): "gpt-4o",
			string(ProviderClaude): "claude-3-sonnet-20240229",
		},
	}
}

// LoadConfig reads configuration from disk
func LoadConfig() (*Config, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return nil, err
	}

	configFile := filepath.Join(configDir, "config.json")

	// If config file doesn't exist, return default config
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return GetDefaultConfig(), nil
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return &config, nil
}

// SaveConfig saves the configuration to disk
func SaveConfig(config *Config) error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	configFile := filepath.Join(configDir, "config.json")

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling config: %w", err)
	}

	return os.WriteFile(configFile, data, 0600)
}

// SaveAPIKey saves an API key for a specific provider
func SaveAPIKey(provider AIProvider, apiKey string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	if config.APIKeys == nil {
		config.APIKeys = make(map[string]string)
	}

	config.APIKeys[string(provider)] = apiKey
	return SaveConfig(config)
}

// SetActiveProvider sets the active AI provider
func SetActiveProvider(provider AIProvider) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	config.ActiveProvider = provider
	return SaveConfig(config)
}

// getConfigDir returns the path to the config directory
func getConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not find home directory: %w", err)
	}

	configDir := filepath.Join(home, ".config", "gitcm")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", fmt.Errorf("could not create config directory: %w", err)
	}

	return configDir, nil
}
