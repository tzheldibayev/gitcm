# gitcm - AI-Powered Git Commit Messages

`gitcm` is a command-line tool that uses AI to generate meaningful commit messages based on your git diff. It analyzes the changes you've staged and suggests appropriate commit messages following conventional commit format.

## Features

- Generate contextual commit messages based on staged changes
- Support for multiple AI providers (OpenAI and Anthropic's Claude)
- Easy configuration and usage
- Seamless integration with your git workflow
- Conventional commit format for consistent messaging

## Installation

### Prerequisites

- Go 1.18 or later
- Git
- API key for at least one supported AI provider

### Install from source

# Clone the repository
git clone https://github.com/tzheldibayev/gitcm.git
cd gitcm

# Build the application
go build -o gitcm

# Move to a location in your PATH
sudo mv gitcm /usr/local/bin/

## Configuration

Before using `gitcm`, you need to set up your API keys:

### OpenAI configuration

```bash
gitcm config set-openai-key YOUR_OPENAI_API_KEY
```

### Claude configuration

```bash
gitcm config set-claude-key YOUR_CLAUDE_API_KEY
```
### Select which provider to use

```bash
gitcm config use-provider openai  # Use OpenAI (default)
# OR
gitcm config use-provider claude  # Use Claude
```
## Usage

### Get a commit suggestion

1. Stage your changes with `git add`
2. Run the suggestion command:

```bash
gitcm suggest
```
This will:
- Analyze your staged changes
- Generate a commit message suggestion
- Ask if you want to use it for your commit

### Example workflow

```bash
# Make changes to your code
# Stage the changes
git add .

# Get a commit suggestion
gitcm suggest

# Review the suggested message and confirm
# Your commit is created!
```

## Command Reference

- `gitcm config set-openai-key KEY` - Set your OpenAI API key
- `gitcm config set-claude-key KEY` - Set your Claude API key  
- `gitcm config use-provider PROVIDER` - Set active provider (openai or claude)
- `gitcm suggest` - Get a commit suggestion for staged changes
- `gitcm suggest-plain` - Get just the suggestion text (for scripting)

## Example Output
```bash
$ gitcm suggest

Getting commit suggestion using openai...

Suggested commit message:
"feat(auth): implement OAuth2 authentication flow with Google provider"

Use this message? (y/n): y
Commit completed!
```
## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

---

*Note: This project was created with assistance from Claude 3.7 Sonnet AI*