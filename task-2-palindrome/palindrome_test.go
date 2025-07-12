package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "hello",
			expected: "olleh",
		},
		{
			input:    "madam",
			expected: "madam",
		},
		{
			input:    "racecar",
			expected: "racecar",
		},
		{
			input:    "GoLang",
			expected: "gnaLoG",
		},
		{
			input:    "",
			expected: "",
		},
	}

	for _, test := range tests {
		result := reverse(test.input)
		if result != test.expected {
			t.Errorf("for input '%s', expected '%s' but got '%s'", test.input, test.expected, result)
		}
	}
}

func TestPalindrome(t *testing.T) {
	tests := []struct {
		original string
		reversed string
		expected bool
	}{
		{
			original: "madam",
			reversed: "madam",
			expected: true,
		},
		{
			original: "hello",
			reversed: "olleh",
			expected: false,
		},
		{
			original: "racecar",
			reversed: "racecar",
			expected: true,
		},
		{
			original: "GoLang",
			reversed: "gnaLoG",
			expected: false,
		},
		{
			original: "",
			reversed: "",
			expected: true,
		},
	}

	for _, test := range tests {
		result := palindromeCheck(test.original, test.reversed)
		if result != test.expected {
			t.Errorf("For original '%s' and reversed '%s', expected %v but got %v", test.original, test.reversed, test.expected, result)
		}
	}
}
