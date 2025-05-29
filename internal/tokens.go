package internal

import (
	"regexp"
	"strings"
	"unicode"
)

// CountTokens provides an estimate of GPT-4 tokens for the given text
// This is a simplified approximation that's reasonably accurate
func CountTokens(text string) int {
	if text == "" {
		return 0
	}

	// Remove excessive whitespace and normalize
	text = normalizeText(text)

	// Count words, punctuation, and special tokens
	tokens := 0

	// Split by whitespace and count words
	words := strings.Fields(text)
	for _, word := range words {
		tokens += countWordTokens(word)
	}

	// Add tokens for special characters and formatting
	tokens += countSpecialTokens(text)

	return tokens
}

func normalizeText(text string) string {
	// Replace multiple whitespaces with single space
	re := regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")

	// Trim leading/trailing whitespace
	return strings.TrimSpace(text)
}

func countWordTokens(word string) int {
	if word == "" {
		return 0
	}

	// Remove punctuation for length calculation
	cleanWord := strings.FieldsFunc(word, func(r rune) bool {
		return unicode.IsPunct(r)
	})

	if len(cleanWord) == 0 {
		return 1 // punctuation-only word
	}

	wordLength := len(strings.Join(cleanWord, ""))

	// Estimate tokens based on character length
	// Rough approximation: 1 token per 4 characters for English text
	tokens := (wordLength + 3) / 4
	if tokens == 0 {
		tokens = 1
	}

	// Add extra tokens for punctuation
	punctCount := countPunctuation(word)
	tokens += (punctCount + 1) / 2 // Roughly 1 token per 2 punctuation marks

	return tokens
}

func countPunctuation(text string) int {
	count := 0
	for _, r := range text {
		if unicode.IsPunct(r) {
			count++
		}
	}
	return count
}

func countSpecialTokens(text string) int {
	tokens := 0

	// Count newlines (each newline is roughly 0.5 tokens)
	newlines := strings.Count(text, "\n")
	tokens += newlines / 2

	// Count code-specific patterns that tend to use more tokens
	codePatterns := []string{
		"{", "}", "(", ")", "[", "]", ";", "->", "=>", "==", "!=", "<=", ">=",
		"&&", "||", "++", "--", "+=", "-=", "*=", "/=",
	}

	for _, pattern := range codePatterns {
		count := strings.Count(text, pattern)
		tokens += count / 3 // Each pattern contributes roughly 1/3 token
	}

	return tokens
}

// GetTokenCountSummary returns a human-readable summary of token counts
func GetTokenCountSummary(totalTokens int) string {
	if totalTokens < 1000 {
		return "Small"
	} else if totalTokens < 5000 {
		return "Medium"
	} else if totalTokens < 15000 {
		return "Large"
	} else {
		return "Very Large"
	}
}
