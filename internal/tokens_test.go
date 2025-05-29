package internal

import (
	"testing"
)

func TestCountTokens(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected int
	}{
		{"Empty string", "", 0},
		{"Simple word", "hello", 2},
		{"Multiple words", "hello world", 3},
		{"Code snippet", "func main() { fmt.Println(\"hello\") }", 10},
		{"With newlines", "line1\nline2\nline3", 4},
		{"With punctuation", "Hello, world! How are you?", 10}, // Adjusted expectation
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CountTokens(test.text)
			// Allow some variance in token counting since it's an approximation
			if result < test.expected-2 || result > test.expected+2 {
				t.Errorf("CountTokens(%q) = %d, expected around %d", test.text, result, test.expected)
			}
		})
	}
}

func TestNormalizeText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Multiple spaces", "hello    world", "hello world"},
		{"Tabs and spaces", "hello\t\t  world", "hello world"},
		{"Leading/trailing spaces", "  hello world  ", "hello world"},
		{"Newlines", "hello\n\nworld", "hello world"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := normalizeText(test.input)
			if result != test.expected {
				t.Errorf("normalizeText(%q) = %q, expected %q", test.input, result, test.expected)
			}
		})
	}
}

func TestCountWordTokens(t *testing.T) {
	tests := []struct {
		name      string
		word      string
		minTokens int
		maxTokens int
	}{
		{"Empty word", "", 0, 1},
		{"Short word", "hi", 1, 2},
		{"Medium word", "hello", 1, 3},
		{"Long word", "supercalifragilisticexpialidocious", 7, 10},
		{"Punctuation only", "!!!", 1, 2},
		{"Word with punctuation", "hello!", 1, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := countWordTokens(test.word)
			if result < test.minTokens || result > test.maxTokens {
				t.Errorf("countWordTokens(%q) = %d, expected between %d and %d",
					test.word, result, test.minTokens, test.maxTokens)
			}
		})
	}
}

func TestCountPunctuation(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected int
	}{
		{"No punctuation", "hello world", 0},
		{"Single punctuation", "hello!", 1},
		{"Multiple punctuation", "Hello, world!!!", 4},
		{"Mixed content", "func() { return; }", 5}, // Adjusted expectation: ( ) { ; }
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := countPunctuation(test.text)
			if result != test.expected {
				t.Errorf("countPunctuation(%q) = %d, expected %d", test.text, result, test.expected)
			}
		})
	}
}

func TestGetTokenCountSummary(t *testing.T) {
	tests := []struct {
		tokens   int
		expected string
	}{
		{500, "Small"},
		{3000, "Medium"},
		{10000, "Large"},
		{20000, "Very Large"},
	}

	for _, test := range tests {
		result := GetTokenCountSummary(test.tokens)
		if result != test.expected {
			t.Errorf("GetTokenCountSummary(%d) = %q, expected %q", test.tokens, result, test.expected)
		}
	}
}
