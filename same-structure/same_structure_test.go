package samestructure

import "testing"

func TestSameStructure(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "valid anagram",
			s1:       "kakap",
			s2:       "kapak",
			expected: true,
		},
		{
			name:     "different characters",
			s1:       "kocak",
			s2:       "kacik",
			expected: false,
		},
		{
			name:     "case sensitive mismatch",
			s1:       "valas",
			s2:       "vaSal",
			expected: false,
		},
		{
			name:     "different length",
			s1:       "abc",
			s2:       "ab",
			expected: false,
		},
		{
			name:     "empty strings",
			s1:       "",
			s2:       "",
			expected: true,
		},
		{
			name:     "single character same",
			s1:       "a",
			s2:       "a",
			expected: true,
		},
		{
			name:     "single character different",
			s1:       "a",
			s2:       "A",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SameStructure(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("SameStructure(%q, %q) = %v; want %v",
					tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}