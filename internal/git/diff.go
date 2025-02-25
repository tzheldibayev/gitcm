// internal/git/diff.go
package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

// GetStagedDiff returns the git diff for staged changes
func GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--staged")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error running git diff: %w", err)
	}

	return out.String(), nil
}

// Commit performs a git commit with the provided message
func Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error committing: %w", err)
	}

	return nil
}
