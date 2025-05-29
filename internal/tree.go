package internal

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

// TreeNode represents a node in the directory tree
type TreeNode struct {
	Name        string
	Path        string
	IsDirectory bool
	TokenCount  int
	Children    []*TreeNode
	Parent      *TreeNode
}

// BuildTree creates a tree structure from the scan results
func BuildTree(result *ScanResult) *TreeNode {
	root := &TreeNode{
		Name:        filepath.Base(result.RootPath),
		Path:        result.RootPath,
		IsDirectory: true,
		Children:    make([]*TreeNode, 0),
	}

	// Create a map for quick lookup
	nodeMap := make(map[string]*TreeNode)
	nodeMap["."] = root

	// Sort files by path for consistent tree building
	sortedFiles := make([]*FileInfo, len(result.Files))
	copy(sortedFiles, result.Files)
	sort.Slice(sortedFiles, func(i, j int) bool {
		return sortedFiles[i].RelativePath < sortedFiles[j].RelativePath
	})

	// Build tree structure
	for _, file := range sortedFiles {
		if file.RelativePath == "." {
			continue
		}

		parts := strings.Split(filepath.ToSlash(file.RelativePath), "/")
		currentPath := ""

		for i, part := range parts {
			if currentPath == "" {
				currentPath = part
			} else {
				currentPath = currentPath + "/" + part
			}

			// Check if node already exists
			if _, exists := nodeMap[currentPath]; !exists {
				// Create new node
				node := &TreeNode{
					Name:        part,
					Path:        filepath.Join(result.RootPath, filepath.FromSlash(currentPath)),
					IsDirectory: i < len(parts)-1 || file.IsDirectory,
					TokenCount:  0,
					Children:    make([]*TreeNode, 0),
				}

				if !node.IsDirectory {
					node.TokenCount = file.TokenCount
				}

				// Find parent
				parentPath := "."
				if i > 0 {
					parentParts := parts[:i]
					parentPath = strings.Join(parentParts, "/")
				}

				parent := nodeMap[parentPath]
				node.Parent = parent
				parent.Children = append(parent.Children, node)
				nodeMap[currentPath] = node
			}
		}
	}

	// Sort children in each node
	sortTreeChildren(root)

	return root
}

func sortTreeChildren(node *TreeNode) {
	// Sort children: directories first, then files, both alphabetically
	sort.Slice(node.Children, func(i, j int) bool {
		if node.Children[i].IsDirectory && !node.Children[j].IsDirectory {
			return true
		}
		if !node.Children[i].IsDirectory && node.Children[j].IsDirectory {
			return false
		}
		return node.Children[i].Name < node.Children[j].Name
	})

	// Recursively sort children
	for _, child := range node.Children {
		sortTreeChildren(child)
	}
}

// RenderTree generates a tree-like string representation
func RenderTree(root *TreeNode, showTokens bool) string {
	var result strings.Builder
	renderNode(root, "", true, true, showTokens, &result)
	return result.String()
}

func renderNode(node *TreeNode, prefix string, isLast bool, isRoot bool, showTokens bool, result *strings.Builder) {
	if !isRoot {
		// Determine the tree characters
		var connector string
		if isLast {
			connector = "└── "
		} else {
			connector = "├── "
		}

		// Write the node line
		line := prefix + connector + node.Name
		if showTokens && !node.IsDirectory && node.TokenCount > 0 {
			line += fmt.Sprintf(" (%d tokens)", node.TokenCount)
		}
		result.WriteString(line + "\n")

		// Update prefix for children
		if isLast {
			prefix += "    "
		} else {
			prefix += "│   "
		}
	} else {
		// Root node
		line := node.Name
		if showTokens {
			totalTokens := calculateTotalTokens(node)
			if totalTokens > 0 {
				line += fmt.Sprintf(" (%d tokens)", totalTokens)
			}
		}
		result.WriteString(line + "\n")
	}

	// Render children
	for i, child := range node.Children {
		isLastChild := i == len(node.Children)-1
		renderNode(child, prefix, isLastChild, false, showTokens, result)
	}
}

func calculateTotalTokens(node *TreeNode) int {
	total := node.TokenCount
	for _, child := range node.Children {
		total += calculateTotalTokens(child)
	}
	return total
}

// GetTreeStats returns statistics about the tree
func GetTreeStats(root *TreeNode) (int, int) {
	files := 0
	directories := 0

	countNodes(root, &files, &directories)

	return files, directories
}

func countNodes(node *TreeNode, files *int, directories *int) {
	if node.IsDirectory {
		*directories++
	} else {
		*files++
	}

	for _, child := range node.Children {
		countNodes(child, files, directories)
	}
}
