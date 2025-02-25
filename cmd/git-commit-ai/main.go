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
			printUsage()
			return
		}

		if os.Args[2] == "set-api-key" {
			if len(os.Args) < 4 {
				fmt.Println("Error: API key required")
				printUsage()
				return
			}

			if err := config.SaveAPIKey(os.Args[3]); err != nil {
				fmt.Printf("Error saving API key: %v\n", err)
				return
			}

			fmt.Println("API key saved successfully!")
		} else {
			fmt.Printf("Unknown config command: %s\n", os.Args[2])
			printUsage()
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
	fmt.Println("  git-commit-ai config set-api-key YOUR_API_KEY   - Set your OpenAI API key")
	fmt.Println("  git-commit-ai suggest                           - Get a commit suggestion for staged changes")
}

func suggestCommit() error {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	if cfg.OpenAIAPIKey == "" {
		return fmt.Errorf("API key not set. Please run: git-commit-ai config set-api-key YOUR_API_KEY")
	}

	// Get diff
	diff, err := git.GetStagedDiff()
	if err != nil {
		return fmt.Errorf("error getting git diff: %w", err)
	}

	if diff == "" {
		return fmt.Errorf("no staged changes found. Add files with 'git add' first")
	}

	fmt.Println("Getting commit suggestion based on diff...")

	// Get suggestion
	suggestion, err := ai.SuggestCommitMessage(cfg.OpenAIAPIKey, diff, cfg.Model)
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
