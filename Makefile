.PHONY: build install test clean

build:
	go build -o bin/git-commit-ai ./cmd/git-commit-ai

install: build
	cp bin/git-commit-ai /usr/local/bin/git-commit-ai
	@echo "Installed git-commit-ai to /usr/local/bin/"

test:
	go test ./...

clean:
	rm -rf bin/
	go clean