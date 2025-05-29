# code2txt - AI Ready Code Converter

A fast Golang CLI tool that converts code repositories to text files for AI analysis.

## Features

- **Fast scanning**: Process 1000+ files quickly
- **Token counting**: Estimate tokens per file
- **Tree visualization**: Display folder structure like `tree` command
- **Smart filtering**: Skip binary files, respect .gitignore
- **Multiple output formats**: Stdout or file output
- **Cross-platform**: Works on Windows, Mac, and Linux

## Installation

### Option 1: Download Pre-built Binary (Recommended)
Download the latest release for your platform:
- [Windows 64-bit](https://github.com/nav9v/code2txt/releases/latest/download/code2txt-windows-amd64.exe)
- [macOS Intel](https://github.com/nav9v/code2txt/releases/latest/download/code2txt-macos-intel)
- [macOS Apple Silicon](https://github.com/nav9v/code2txt/releases/latest/download/code2txt-macos-arm64)
- [Linux 64-bit](https://github.com/nav9v/code2txt/releases/latest/download/code2txt-linux-amd64)

### Option 2: Install with Go (requires Go 1.21+)
```bash
go install github.com/nav9v/code2txt@latest
```

### Option 3: Build from Source
```bash
git clone https://github.com/nav9v/code2txt.git
cd code2txt
go build -o code2txt
```

### Option 4: Local Build (if you have the source)
```bash
# In the project directory
go build -o code2txt
# On Windows: go build -o code2txt.exe
```

## Usage

### Basic Commands

```bash
# Scan folder, output to stdout
code2txt <folder>

# Scan folder, save to file
code2txt <folder> -o output.txt

# Show token count for each file
code2txt <folder> --tokens

# Include only specific file types
code2txt <folder> -i "*.go,*.md"

# Exclude specific patterns
code2txt <folder> -e "*.log,node_modules"
```

### Advanced Options

```bash
-o, --output <file>        # Output file (default: stdout)
-i, --include <patterns>   # Include patterns (*.go,*.md)
-e, --exclude <patterns>   # Exclude patterns (*.log,node_modules)
--tokens                   # Show token counts
--no-tree                  # Skip tree structure
--max-tokens <n>           # Skip files over N tokens
```

## Output Format

```
Directory Structure:
├── main.go (245 tokens)
├── config/
│   ├── config.go (456 tokens)
│   └── config_test.go (123 tokens)
└── README.md (234 tokens)

Total: 1,058 tokens

File Contents:
=============

File: main.go
-------------
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}

File: config/config.go
---------------------
[file contents here]
```

## Examples

### Scan a Go project with token counts
```bash
code2txt ./my-go-project --tokens
```

### Export only Go files to a text file
```bash
code2txt ./project -i "*.go" -o project-code.txt
```

### Scan excluding common build artifacts
```bash
code2txt ./project -e "*.exe,*.dll,node_modules,target"
```

## Performance

- Scans typical Go projects in under 2 seconds
- Accurate token counting (within 5% of actual GPT-4 tokens)
- Memory efficient for large codebases
- Respects .gitignore files automatically

## License

MIT License