.PHONY: build install test clean

build:
	go build -o bin/git-commit-ai ./cmd/git-commit-ai/main.go

install: build
	mkdir -p $(HOME)/bin
	cp bin/git-commit-ai $(HOME)/bin/
	@echo "Installed git-commit-ai to $(HOME)/bin/"
	@echo "Make sure $(HOME)/bin is in your PATH"

test:
	go test ./...

clean:
	rm -rf bin/
	go clean