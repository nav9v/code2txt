# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Comprehensive test suite for all packages
- Modern GitHub Actions workflows for CI/CD
- Cross-platform build support (Windows, Linux, macOS Intel/ARM)
- Automated release process with binary uploads
- Version information in build artifacts
- Comprehensive documentation and README
- Makefile for development workflow
- Code coverage reporting
- Multi-platform testing in CI

### Changed
- Updated GitHub Actions to use latest versions (v4, v5)
- Improved error handling and test coverage
- Enhanced documentation with usage examples
- Modernized development workflow

### Removed
- Deprecated GitHub Actions (actions/setup-go@v2, actions/checkout@v2)
- Unnecessary build scripts (build.bat, build.sh)
- Output files and temporary artifacts

### Fixed
- Token counting test expectations
- Cross-platform compatibility issues
- Build process reliability

## [Previous Versions]

### Added
- Initial release with basic functionality
- CLI interface with Cobra framework
- Directory scanning with gitignore support
- Token estimation for AI models
- Tree visualization output
- File filtering and pattern matching