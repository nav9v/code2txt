package internal

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

type ScanOptions struct {
	IncludePatterns []string
	ExcludePatterns []string
	MaxTokens       int
}

type FileInfo struct {
	Path         string
	RelativePath string
	Size         int64
	TokenCount   int
	Content      string
	IsDirectory  bool
}

type ScanResult struct {
	RootPath    string
	Files       []*FileInfo
	TotalTokens int
	TotalFiles  int
}

type Scanner struct {
	options           *ScanOptions
	gitignorePatterns []string
}

func NewScanner(options *ScanOptions) *Scanner {
	if options == nil {
		options = &ScanOptions{}
	}

	// Default exclude patterns
	if len(options.ExcludePatterns) == 0 {
		options.ExcludePatterns = []string{
			"*.exe", "*.dll", "*.so", "*.dylib",
			"*.jpg", "*.jpeg", "*.png", "*.gif", "*.bmp",
			"*.mp3", "*.mp4", "*.avi", "*.mov",
			"*.zip", "*.tar", "*.gz", "*.rar",
			"node_modules", ".git", ".svn", ".hg",
			"*.log", "*.tmp", "*.cache",
			".DS_Store", "Thumbs.db",
		}
	}

	return &Scanner{
		options: options,
	}
}

func (s *Scanner) ScanDirectory(rootPath string) (*ScanResult, error) {
	result := &ScanResult{
		RootPath: rootPath,
		Files:    make([]*FileInfo, 0),
	}

	// Load .gitignore if it exists
	s.loadGitignore(rootPath)

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(rootPath, path)

		// Skip if matches exclude patterns
		if s.shouldExclude(relPath, d.IsDir()) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip if doesn't match include patterns (when specified)
		if len(s.options.IncludePatterns) > 0 && !d.IsDir() {
			if !s.shouldInclude(relPath) {
				return nil
			}
		}

		fileInfo := &FileInfo{
			Path:         path,
			RelativePath: relPath,
			IsDirectory:  d.IsDir(),
		}

		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}

			fileInfo.Size = info.Size()

			// Skip large files (over 10MB)
			if fileInfo.Size > 10*1024*1024 {
				return nil
			}

			// Read and process file content
			if err := s.processFile(fileInfo); err != nil {
				// Skip files that can't be read or processed
				return nil
			}

			// Skip if over max tokens limit
			if s.options.MaxTokens > 0 && fileInfo.TokenCount > s.options.MaxTokens {
				return nil
			}

			result.TotalTokens += fileInfo.TokenCount
			result.TotalFiles++
		}

		result.Files = append(result.Files, fileInfo)
		return nil
	})

	return result, err
}

func (s *Scanner) processFile(fileInfo *FileInfo) error {
	content, err := os.ReadFile(fileInfo.Path)
	if err != nil {
		return err
	}

	// Check if file is binary
	if !utf8.Valid(content) && !s.isTextFile(fileInfo.Path) {
		return fmt.Errorf("binary file")
	}

	fileInfo.Content = string(content)
	fileInfo.TokenCount = CountTokens(fileInfo.Content)

	return nil
}

func (s *Scanner) isTextFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	textExtensions := []string{
		".txt", ".md", ".rst", ".json", ".xml", ".yaml", ".yml",
		".go", ".py", ".js", ".ts", ".java", ".c", ".cpp", ".h", ".hpp",
		".cs", ".php", ".rb", ".swift", ".kt", ".rs", ".sh", ".bat",
		".html", ".css", ".scss", ".less", ".sql", ".r", ".m",
	}

	for _, textExt := range textExtensions {
		if ext == textExt {
			return true
		}
	}

	return false
}

func (s *Scanner) shouldExclude(path string, isDir bool) bool {
	// Check gitignore patterns
	for _, pattern := range s.gitignorePatterns {
		if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
			return true
		}
	}

	// Check exclude patterns
	for _, pattern := range s.options.ExcludePatterns {
		if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
			return true
		}
		// Also check full path for directory patterns
		if strings.Contains(path, pattern) {
			return true
		}
	}

	return false
}

func (s *Scanner) shouldInclude(path string) bool {
	for _, pattern := range s.options.IncludePatterns {
		if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
			return true
		}
	}
	return false
}

func (s *Scanner) loadGitignore(rootPath string) {
	gitignorePath := filepath.Join(rootPath, ".gitignore")
	file, err := os.Open(gitignorePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			s.gitignorePatterns = append(s.gitignorePatterns, line)
		}
	}
}
