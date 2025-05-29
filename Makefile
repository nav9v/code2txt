BINARY_NAME=code2txt
VERSION?=$(shell git describe --tags --always --dirty)
LDFLAGS=-ldflags="-s -w -X main.version=$(VERSION)"

.PHONY: build test clean install lint fmt vet deps

# Default target
all: test build

# Run tests
test:
	go test -v ./...

# Build for current platform
build:
	go build $(LDFLAGS) -o $(BINARY_NAME)

# Build for all platforms
build-all:
	mkdir -p dist
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64

# Install dependencies
deps:
	go mod download
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)
	rm -rf dist/
	rm -f output.txt

# Install to GOPATH/bin
install:
	go install $(LDFLAGS)

# Run with sample data
demo:
	./$(BINARY_NAME) . --tokens

# Development workflow
dev: deps fmt vet test build

# Release preparation
release-prep: clean deps fmt vet test build-all