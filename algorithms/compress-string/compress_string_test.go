package compressstring

import "testing"

func TestCompressString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "no compression benefit",
			input:    "aabbcc",
			expected: "aabbcc",
		},
		{
			name:     "compressible string",
			input:    "aaabcccccaaa",
			expected: "a3b1c5a3",
		},
		{
			name:     "all unique characters",
			input:    "abcd",
			expected: "abcd",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "long repeating character",
			input:    "aaaaaaaaaa",
			expected: "a10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompressString(tt.input)
			if result != tt.expected {
				t.Errorf("CompressString(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}