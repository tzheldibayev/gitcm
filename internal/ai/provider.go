// internal/ai/provider.go
package ai

import (
	"fmt"

	"github.com/tzheldibayev/gitcm/internal/config"
)

// Provider defines the interface for AI providers
type Provider interface {
	SuggestCommitMessage(diff string) (string, error)
}

// NewProvider creates a provider based on the configuration
func NewProvider(cfg *config.Config) (Provider, error) {
	provider := cfg.ActiveProvider
	apiKey, ok := cfg.APIKeys[string(provider)]
	if !ok || apiKey == "" {
		return nil, fmt.Errorf("API key not set for provider %s", provider)
	}

	model, ok := cfg.Models[string(provider)]
	if !ok {
		return nil, fmt.Errorf("model not configured for provider %s", provider)
	}

	switch provider {
	case config.ProviderOpenAI:
		return NewOpenAIProvider(apiKey, model), nil
	case config.ProviderClaude:
		return NewClaudeProvider(apiKey, model), nil
	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}
}
