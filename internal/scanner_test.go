package internal

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewScanner(t *testing.T) {
	// Test with nil options
	scanner := NewScanner(nil)
	if scanner == nil {
		t.Error("Expected scanner to be created, got nil")
	}

	if len(scanner.options.ExcludePatterns) == 0 {
		t.Error("Expected default exclude patterns to be set")
	}

	// Test with custom options
	options := &ScanOptions{
		IncludePatterns: []string{"*.go"},
		ExcludePatterns: []string{"*.log"},
		MaxTokens:       1000,
	}

	scanner = NewScanner(options)
	if len(scanner.options.IncludePatterns) != 1 {
		t.Error("Expected include patterns to be preserved")
	}
}

func TestScanDirectory(t *testing.T) {
	// Create a temporary directory with test files
	tempDir := t.TempDir()

	// Create test files
	goFile := filepath.Join(tempDir, "main.go")
	goContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`

	err := os.WriteFile(goFile, []byte(goContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test Go file: %v", err)
	}

	txtFile := filepath.Join(tempDir, "readme.txt")
	txtContent := "This is a test file."

	err = os.WriteFile(txtFile, []byte(txtContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test txt file: %v", err)
	}

	// Create scanner and scan directory
	scanner := NewScanner(&ScanOptions{})
	result, err := scanner.ScanDirectory(tempDir)

	if err != nil {
		t.Fatalf("Failed to scan directory: %v", err)
	}

	if result == nil {
		t.Fatal("Expected scan result, got nil")
	}

	if result.TotalFiles == 0 {
		t.Error("Expected files to be found")
	}

	if result.RootPath != tempDir {
		t.Errorf("Expected root path %s, got %s", tempDir, result.RootPath)
	}
}

func TestShouldExclude(t *testing.T) {
	scanner := NewScanner(&ScanOptions{
		ExcludePatterns: []string{"*.log", "node_modules"},
	})

	tests := []struct {
		path     string
		isDir    bool
		expected bool
	}{
		{"test.log", false, true},
		{"test.go", false, false},
		{"node_modules", true, true},
		{"src", true, false},
	}

	for _, test := range tests {
		result := scanner.shouldExclude(test.path, test.isDir)
		if result != test.expected {
			t.Errorf("shouldExclude(%s, %t) = %t, expected %t",
				test.path, test.isDir, result, test.expected)
		}
	}
}

func TestShouldInclude(t *testing.T) {
	scanner := NewScanner(&ScanOptions{
		IncludePatterns: []string{"*.go", "*.js"},
	})

	tests := []struct {
		path     string
		expected bool
	}{
		{"main.go", true},
		{"app.js", true},
		{"style.css", false},
		{"readme.txt", false},
	}

	for _, test := range tests {
		result := scanner.shouldInclude(test.path)
		if result != test.expected {
			t.Errorf("shouldInclude(%s) = %t, expected %t",
				test.path, result, test.expected)
		}
	}
}
