package cmd

import (
	"fmt"
	"os"

	"github.com/nav9v/code2txt/internal"
	"github.com/spf13/cobra"
)

var (
	outputFile      string
	includePatterns []string
	excludePatterns []string
	showTokens      bool
	noTree          bool
	maxTokens       int
)

var rootCmd = &cobra.Command{
	Use:   "code2txt <folder>",
	Short: "Convert code repositories to text files for AI analysis",
	Long: `code2txt - AI Ready Code Converter

A fast CLI tool that scans code repositories and converts them into AI-friendly
text format. Perfect for feeding codebases to ChatGPT, Claude, or other AI models.

Features:
  • Fast scanning with smart filtering
  • Accurate token counting for AI models
  • Beautiful tree structure visualization
  • Respects .gitignore files automatically
  • Cross-platform support (Windows, Mac, Linux)

Examples:
  code2txt ./my-project                    # Scan project, output to console
  code2txt ./src --tokens                  # Show token counts for each file
  code2txt ./app -o analysis.txt           # Save output to file
  code2txt ./code -i "*.go,*.js"           # Only include Go and JS files
  code2txt ./proj -e "*.log,node_modules"  # Exclude logs and dependencies`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		folderPath := args[0]

		// Validate folder exists
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			return fmt.Errorf("folder does not exist: %s", folderPath)
		}

		// Create scanner with options
		scanner := internal.NewScanner(&internal.ScanOptions{
			IncludePatterns: includePatterns,
			ExcludePatterns: excludePatterns,
			MaxTokens:       maxTokens,
		})

		// Scan the directory
		result, err := scanner.ScanDirectory(folderPath)
		if err != nil {
			return fmt.Errorf("failed to scan directory: %w", err)
		}

		// Create output formatter
		formatter := internal.NewOutputFormatter(&internal.OutputOptions{
			ShowTokens: showTokens,
			ShowTree:   !noTree,
		})

		// Generate output
		output := formatter.FormatOutput(result)

		// Write to file or stdout
		if outputFile != "" {
			if err := os.WriteFile(outputFile, []byte(output), 0644); err != nil {
				return fmt.Errorf("failed to write output file: %w", err)
			}
			fmt.Printf("Output written to: %s\n", outputFile)
		} else {
			fmt.Print(output)
		}

		return nil
	},
}

func init() {
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "",
		"Save output to file instead of printing to console\n"+
			"Example: -o report.txt")

	rootCmd.Flags().StringSliceVarP(&includePatterns, "include", "i", []string{},
		"Only include files matching these patterns (comma-separated)\n"+
			"Example: -i \"*.go,*.js,*.py\" (only Go, JavaScript, Python files)")

	rootCmd.Flags().StringSliceVarP(&excludePatterns, "exclude", "e", []string{},
		"Exclude files/folders matching these patterns (comma-separated)\n"+
			"Example: -e \"*.log,node_modules,target,dist\" (skip logs & build dirs)")

	rootCmd.Flags().BoolVar(&showTokens, "tokens", false,
		"Display estimated token count for each file and total\n"+
			"Useful for estimating AI model costs (GPT-4, Claude, etc.)")

	rootCmd.Flags().BoolVar(&noTree, "no-tree", false,
		"Skip the directory tree visualization in output\n"+
			"Only show file contents (faster for large projects)")

	rootCmd.Flags().IntVar(&maxTokens, "max-tokens", 0,
		"Skip files larger than N tokens (0 = no limit)\n"+
			"Example: --max-tokens 5000 (skip files over 5k tokens)")
}

func Execute() error {
	return rootCmd.Execute()
}
