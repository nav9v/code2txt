# code2txt 🚀

[![CI](https://github.com/nav9v/code2txt/actions/workflows/ci.yml/badge.svg)](https://github.com/nav9v/code2txt/actions/workflows/ci.yml)
[![Release](https://github.com/nav9v/code2txt/actions/workflows/release.yml/badge.svg)](https://github.com/nav9v/code2txt/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/nav9v/code2txt)](https://goreportcard.com/report/github.com/nav9v/code2txt)

A fast CLI tool that converts code repositories into AI-friendly text format. Perfect for feeding codebases to ChatGPT, Claude, or other AI models.

## ✨ Features

- **🚀 Fast scanning**: Process 1000+ files quickly
- **🧮 Token counting**: Accurate GPT-4/Claude token estimation  
- **🌳 Tree visualization**: Beautiful directory structure display
- **🎯 Smart filtering**: Respects .gitignore automatically
- **🔧 Cross-platform**: Works on Windows, Mac, and Linux
- **📝 Flexible patterns**: Include/exclude file patterns
- **💻 Easy integration**: Works great in CI/CD pipelines

## 📥 Installation

### Download Binary (Recommended)

1. Go to the [Releases page](https://github.com/nav9v/code2txt/releases)
2. Download the binary for your platform:
   - **Windows**: `code2txt-windows-amd64.exe`
   - **Linux**: `code2txt-linux-amd64`
   - **macOS Intel**: `code2txt-darwin-amd64`
   - **macOS Apple Silicon**: `code2txt-darwin-arm64`
3. Make executable (Linux/macOS): `chmod +x code2txt-*`
4. Move to PATH or run directly

### Build from Source

```bash
# Clone the repository
git clone https://github.com/nav9v/code2txt.git
cd code2txt

# Build
make build
# or
go build -o code2txt

# Install to GOPATH/bin
make install
# or
go install
```

## 🔧 Usage

### Basic Usage

```bash
# Scan current directory
code2txt .

# Scan specific folder
code2txt ./my-project

# Show token counts
code2txt ./src --tokens

# Save to file
code2txt ./app -o analysis.txt
```

### Advanced Usage

```bash
# Only include specific file types
code2txt ./src -i "*.go,*.js,*.py"

# Exclude certain patterns
code2txt ./project -e "*.log,node_modules,target"

# Skip large files (over 5000 tokens)
code2txt ./code --max-tokens 5000

# Skip tree visualization for faster processing
code2txt ./large-project --no-tree
```

### Useful Examples

```bash
# Analyze a Go project for AI review
code2txt ./my-go-app -i "*.go" --tokens -o go-analysis.txt

# Prepare frontend code for AI assistance
code2txt ./frontend -i "*.js,*.ts,*.jsx,*.tsx,*.css" --tokens

# Quick scan without binary/media files
code2txt ./project -e "*.jpg,*.png,*.mp4,node_modules,dist"

# Large project analysis with limits
code2txt ./enterprise-app --max-tokens 10000 --no-tree
```

## 🛠️ Development

### Prerequisites

- Go 1.21 or later
- Make (optional, for convenience)

### Development Workflow

```bash
# Install dependencies
make deps

# Run tests
make test

# Format code
make fmt

# Lint code (requires golangci-lint)
make lint

# Build for current platform
make build

# Build for all platforms
make build-all

# Full development cycle
make dev
```

### Testing

```bash
# Run all tests
go test -v ./...

# Run with coverage
go test -race -coverprofile=coverage.out ./...

# View coverage report
go tool cover -html=coverage.out
```

## 📦 Release Process

### Automated Releases (Recommended)

1. **Create a new tag**:
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```

2. **Create GitHub Release**:
   - Go to GitHub → Releases → "Create a new release"
   - Select your tag
   - GitHub Actions will automatically build and upload binaries

### Manual Release

```bash
# Prepare release
make release-prep

# Upload dist/ files to GitHub release manually
```

## 🔄 GitHub Workflows

### Continuous Integration (`ci.yml`)

- **Triggers**: Push to main/develop, Pull requests
- **Tests**: Cross-platform testing (Ubuntu, Windows, macOS)
- **Go versions**: 1.21, 1.22
- **Checks**: Tests, formatting, vetting, building

### Release (`release.yml`)

- **Triggers**: GitHub release creation
- **Builds**: Multi-platform binaries
- **Assets**: Automated upload with checksums
- **Description**: Auto-generated release notes

## 📋 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Make changes and add tests
4. Run the development workflow: `make dev`
5. Commit changes: `git commit -m 'Add amazing feature'`
6. Push to branch: `git push origin feature/amazing-feature`
7. Open a Pull Request

### Code Style

- Follow standard Go formatting (`gofmt`)
- Add tests for new functionality
- Update documentation as needed
- Ensure all CI checks pass

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Support

- 🐛 **Issues**: [GitHub Issues](https://github.com/nav9v/code2txt/issues)
- 💡 **Feature Requests**: [GitHub Discussions](https://github.com/nav9v/code2txt/discussions)
- 📧 **Contact**: Create an issue for questions

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI functionality
- Inspired by the need for better AI-code interaction tools