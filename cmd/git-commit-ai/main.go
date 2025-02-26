// cmd/git-commit-ai/main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tzheldibayev/git-commit-ai/internal/ai"
	"github.com/tzheldibayev/git-commit-ai/internal/config"
	"github.com/tzheldibayev/git-commit-ai/internal/git"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "config":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing subcommand for config")
			printConfigUsage()
			return
		}

		switch os.Args[2] {
		case "set-openai-key":
			if len(os.Args) < 4 {
				fmt.Println("Error: API key required")
				printConfigUsage()
				return
			}

			if err := config.SaveAPIKey(config.ProviderOpenAI, os.Args[3]); err != nil {
				fmt.Printf("Error saving OpenAI API key: %v\n", err)
				return
			}

			fmt.Println("OpenAI API key saved successfully!")

		case "set-claude-key":
			if len(os.Args) < 4 {
				fmt.Println("Error: API key required")
				printConfigUsage()
				return
			}

			if err := config.SaveAPIKey(config.ProviderClaude, os.Args[3]); err != nil {
				fmt.Printf("Error saving Claude API key: %v\n", err)
				return
			}

			fmt.Println("Claude API key saved successfully!")

		case "use-provider":
			if len(os.Args) < 4 {
				fmt.Println("Error: Provider name required")
				printConfigUsage()
				return
			}

			providerName := os.Args[3]
			var provider config.AIProvider

			switch providerName {
			case "openai":
				provider = config.ProviderOpenAI
			case "claude":
				provider = config.ProviderClaude
			default:
				fmt.Printf("Error: Unknown provider '%s'. Use 'openai' or 'claude'.\n", providerName)
				return
			}

			if err := config.SetActiveProvider(provider); err != nil {
				fmt.Printf("Error setting active provider: %v\n", err)
				return
			}

			fmt.Printf("Now using %s as the active provider.\n", providerName)

		default:
			fmt.Printf("Unknown config command: %s\n", os.Args[2])
			printConfigUsage()
		}

	case "suggest":
		if err := suggestCommit(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}

	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Git Commit AI Assistant")
	fmt.Println("Usage:")
	fmt.Println("  git-commit-ai config [subcommand]            - Configure the assistant")
	fmt.Println("  git-commit-ai suggest                        - Get a commit suggestion for staged changes")
	fmt.Println("\nRun 'git-commit-ai config' to see configuration options.")
}

func printConfigUsage() {
	fmt.Println("Configuration Commands:")
	fmt.Println("  git-commit-ai config set-openai-key KEY      - Set your OpenAI API key")
	fmt.Println("  git-commit-ai config set-claude-key KEY      - Set your Claude API key")
	fmt.Println("  git-commit-ai config use-provider PROVIDER   - Set active provider (openai or claude)")
}

func suggestCommit() error {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	// Get diff
	diff, err := git.GetStagedDiff()
	if err != nil {
		return fmt.Errorf("error getting git diff: %w", err)
	}

	if diff == "" {
		return fmt.Errorf("no staged changes found. Add files with 'git add' first")
	}

	// Create AI provider
	provider, err := ai.NewProvider(cfg)
	if err != nil {
		if strings.Contains(err.Error(), "API key not set") {
			return fmt.Errorf("%v. Please run appropriate config command to set API key", err)
		}
		return fmt.Errorf("error initializing AI provider: %w", err)
	}

	fmt.Printf("Getting commit suggestion using %s...\n", cfg.ActiveProvider)

	// Get suggestion
	suggestion, err := provider.SuggestCommitMessage(diff)
	if err != nil {
		return fmt.Errorf("error getting suggestion: %w", err)
	}

	// Show the suggestion
	fmt.Println("\nSuggested commit message:")
	fmt.Printf("\"%s\"\n\n", suggestion)

	// Confirm with user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Use this message? (y/n): ")
	response, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	response = strings.ToLower(strings.TrimSpace(response))
	if response == "y" || response == "yes" {
		if err := git.Commit(suggestion); err != nil {
			return fmt.Errorf("error committing: %w", err)
		}
		fmt.Println("Commit completed!")
	} else {
		fmt.Println("Commit cancelled.")
	}

	return nil
}
