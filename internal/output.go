package internal

import (
	"fmt"
	"strings"
)

type OutputOptions struct {
	ShowTokens bool
	ShowTree   bool
}

type OutputFormatter struct {
	options *OutputOptions
}

func NewOutputFormatter(options *OutputOptions) *OutputFormatter {
	if options == nil {
		options = &OutputOptions{
			ShowTokens: false,
			ShowTree:   true,
		}
	}

	return &OutputFormatter{
		options: options,
	}
}

func (f *OutputFormatter) FormatOutput(result *ScanResult) string {
	var output strings.Builder

	// Generate tree structure if enabled
	if f.options.ShowTree {
		output.WriteString("Directory Structure:\n")
		tree := BuildTree(result)
		treeOutput := RenderTree(tree, f.options.ShowTokens)
		output.WriteString(treeOutput)
		output.WriteString("\n")

		// Add summary statistics
		if f.options.ShowTokens {
			output.WriteString(fmt.Sprintf("Total: %s tokens (%s)\n\n",
				formatNumber(result.TotalTokens),
				GetTokenCountSummary(result.TotalTokens)))
		} else {
			output.WriteString(fmt.Sprintf("Total files: %d\n\n", result.TotalFiles))
		}
	}

	// Generate file contents section
	output.WriteString("File Contents:\n")
	output.WriteString(strings.Repeat("=", 50) + "\n\n")

	// Get only files (not directories) and sort them
	files := make([]*FileInfo, 0)
	for _, file := range result.Files {
		if !file.IsDirectory {
			files = append(files, file)
		}
	}

	// Sort files by relative path
	sortFiles(files)

	for i, file := range files {
		if i > 0 {
			output.WriteString("\n")
		}

		// File header
		header := fmt.Sprintf("File: %s", file.RelativePath)
		if f.options.ShowTokens {
			header += fmt.Sprintf(" (%d tokens)", file.TokenCount)
		}
		output.WriteString(header + "\n")
		output.WriteString(strings.Repeat("-", len(header)) + "\n")

		// File content
		if file.Content != "" {
			output.WriteString(file.Content)
			// Ensure file ends with newline
			if !strings.HasSuffix(file.Content, "\n") {
				output.WriteString("\n")
			}
		} else {
			output.WriteString("(empty file)\n")
		}
	}

	return output.String()
}

func sortFiles(files []*FileInfo) {
	// Simple bubble sort for file paths
	n := len(files)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if files[j].RelativePath > files[j+1].RelativePath {
				files[j], files[j+1] = files[j+1], files[j]
			}
		}
	}
}

func formatNumber(num int) string {
	if num < 1000 {
		return fmt.Sprintf("%d", num)
	} else if num < 1000000 {
		return fmt.Sprintf("%.1fk", float64(num)/1000)
	} else {
		return fmt.Sprintf("%.1fM", float64(num)/1000000)
	}
}

// FormatFileList creates a simple list of files with token counts
func (f *OutputFormatter) FormatFileList(result *ScanResult) string {
	var output strings.Builder

	output.WriteString("Files and Token Counts:\n")
	output.WriteString(strings.Repeat("=", 30) + "\n")

	for _, file := range result.Files {
		if !file.IsDirectory {
			line := fmt.Sprintf("%-40s %6d tokens\n", file.RelativePath, file.TokenCount)
			output.WriteString(line)
		}
	}

	output.WriteString(strings.Repeat("-", 50) + "\n")
	output.WriteString(fmt.Sprintf("Total: %d files, %s tokens\n",
		result.TotalFiles, formatNumber(result.TotalTokens)))

	return output.String()
}

// FormatSummary creates a brief summary of the scan results
func (f *OutputFormatter) FormatSummary(result *ScanResult) string {
	var output strings.Builder

	output.WriteString("Scan Summary:\n")
	output.WriteString(fmt.Sprintf("Root Path: %s\n", result.RootPath))
	output.WriteString(fmt.Sprintf("Files: %d\n", result.TotalFiles))
	output.WriteString(fmt.Sprintf("Total Tokens: %s (%s)\n",
		formatNumber(result.TotalTokens),
		GetTokenCountSummary(result.TotalTokens)))

	return output.String()
}
