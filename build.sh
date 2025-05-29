#!/bin/bash

# Build script for code2txt
# Creates binaries for all supported platforms

set -e

VERSION=${1:-"dev"}
LDFLAGS="-s -w -X main.version=$VERSION"

echo "Building code2txt $VERSION..."

# Create dist directory
mkdir -p dist

# Windows AMD64
echo "Building Windows AMD64..."
GOOS=windows GOARCH=amd64 go build -ldflags="$LDFLAGS" -o dist/code2txt-windows-amd64.exe

# Linux AMD64
echo "Building Linux AMD64..."
GOOS=linux GOARCH=amd64 go build -ldflags="$LDFLAGS" -o dist/code2txt-linux-amd64

# macOS Intel
echo "Building macOS Intel..."
GOOS=darwin GOARCH=amd64 go build -ldflags="$LDFLAGS" -o dist/code2txt-macos-intel

# macOS Apple Silicon
echo "Building macOS Apple Silicon..."
GOOS=darwin GOARCH=arm64 go build -ldflags="$LDFLAGS" -o dist/code2txt-macos-arm64

echo "Build complete! Binaries available in dist/"
echo "File sizes:"
ls -lh dist/