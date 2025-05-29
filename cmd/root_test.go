package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRootCommand(t *testing.T) {
	// Test that the command can be executed without errors
	cmd := rootCmd

	// Test command configuration
	if cmd.Use != "code2txt <folder>" {
		t.Errorf("Expected Use to be 'code2txt <folder>', got %s", cmd.Use)
	}

	if cmd.Short == "" {
		t.Error("Expected Short description to be set")
	}

	if cmd.Long == "" {
		t.Error("Expected Long description to be set")
	}
}

func TestExecuteWithNonExistentFolder(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()
	nonExistentPath := filepath.Join(tempDir, "nonexistent")

	// Set up command args
	rootCmd.SetArgs([]string{nonExistentPath})

	// Execute should return an error for non-existent folder
	err := rootCmd.Execute()
	if err == nil {
		t.Error("Expected error for non-existent folder, got nil")
	}
}

func TestExecuteWithValidFolder(t *testing.T) {
	// Create a temporary directory with a test file
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.go")

	content := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`

	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Set up command args
	rootCmd.SetArgs([]string{tempDir})

	// Execute should not return an error for valid folder
	err = rootCmd.Execute()
	if err != nil {
		t.Errorf("Expected no error for valid folder, got: %v", err)
	}
}
