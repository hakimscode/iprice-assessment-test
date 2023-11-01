package string_manipulation

import "testing"

type Test struct {
	name     string
	input    string
	expected string
}

func TestGetUppercase(t *testing.T) {
	tests := []Test{
		{
			name:     "Test input normal string",
			input:    "uppercase",
			expected: "UPPERCASE",
		},
		{
			name:     "Test input string with space",
			input:    "hello world",
			expected: "HELLO WORLD",
		},
		{
			name:     "Test input string containing uppercase",
			input:    "Hello World",
			expected: "HELLO WORLD",
		},
	}

	for _, test := range tests {
		result := GetUppercase(test.input)

		if result != test.expected {
			t.Errorf("Failed, expected output is %s but the result is %s", test.expected, result)
		}
	}
}

func TestAlternateCase(t *testing.T) {
	tests := []Test{
		{
			name:     "Test input normal string",
			input:    "alternatecase",
			expected: "AlTeRnAtEcAsE",
		},
		{
			name:     "Test input string with space",
			input:    "hello world",
			expected: "HeLlO WoRlD",
		},
		{
			name:     "Test input string containing uppercase",
			input:    "Hello World",
			expected: "HeLlO WoRlD",
		},
		{
			name:     "Test input string nearly alternate case",
			input:    "AltErnAtecASe",
			expected: "AlTeRnAtEcAsE",
		},
	}

	for _, test := range tests {
		result := GetAlternateCase(test.input)

		if result != test.expected {
			t.Errorf("Failed, expected output is %s but the result is %s", test.expected, result)
		}
	}
}
