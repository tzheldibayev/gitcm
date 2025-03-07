.PHONY: build install test clean

build:
	go build -o bin/gitcm ./cmd/gitcm/main.go

install: build
	mkdir -p $(HOME)/bin
	cp bin/gitcm $(HOME)/bin/
	@echo "Installed gitcm to $(HOME)/bin/"
	@echo "Make sure $(HOME)/bin is in your PATH"

test:
	go test ./...

clean:
	rm -rf bin/
	go clean