package main

import "testing"

func TestWordCount(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input:    "hello world",
			expected: map[string]int{"hello": 1, "world": 1},
		},
		{
			input:    "go go go",
			expected: map[string]int{"go": 3},
		},
		{
			input:    "this is a test this is only a test",
			expected: map[string]int{"this": 2, "is": 2, "a": 2, "test": 2, "only": 1},
		},
		{
			input:    "",
			expected: map[string]int{},
		},
		{
			input:    "one-word",
			expected: map[string]int{"one-word": 1},
		},
	}

	for _, test := range tests {
		result := wordCount(test.input)
		if !equalMaps(result, test.expected) {
			t.Errorf("For input '%s', expected %v but got %v", test.input, test.expected, result)
		}
	}
}

func equalMaps(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, value := range a {
		if b[key] != value {
			return false
		}
	}
	return true
}
