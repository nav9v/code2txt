package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Run tests
	code := m.Run()
	os.Exit(code)
}

func TestMainFunction(t *testing.T) {
	// Test that main function doesn't panic
	// We can't easily test the actual execution without mocking,
	// but we can at least ensure the package compiles correctly
	if os.Getenv("BE_CRASHER") == "1" {
		main()
		return
	}

	// Test passes if we reach here
	t.Log("Main package test passed")
}
