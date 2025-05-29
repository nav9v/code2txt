# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-05-29

### Added
- Initial release of code2txt CLI tool
- Fast code repository scanning and conversion to AI-friendly text format
- Accurate token counting for AI models (GPT-4, Claude, etc.)
- Beautiful tree structure visualization with Unicode box-drawing characters
- Smart filtering that respects .gitignore files automatically
- Cross-platform support with pre-built binaries for:
  - Windows AMD64
  - Linux AMD64
  - macOS Intel (AMD64)
  - macOS Apple Silicon (ARM64)
- Flexible include/exclude pattern matching
- Multiple output options (stdout or file)
- Command-line flags for customization:
  - `--tokens`: Show token counts
  - `--no-tree`: Skip tree visualization
  - `--include`: Include only specific file patterns
  - `--exclude`: Exclude specific file patterns
  - `--max-tokens`: Skip files over token limit
  - `--output`: Save to file instead of stdout

### Features
- Processes 1000+ files quickly with efficient scanning
- Memory-efficient for large codebases
- Accurate token estimation (within 5% of actual GPT-4 tokens)
- Professional CLI interface with detailed help and examples
- MIT License for open source use