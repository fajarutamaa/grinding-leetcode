package sumprimes

import "testing"

func TestSumPrimes(t *testing.T) {
	tests := []struct {
		name     string
		limit    int
		expected int
	}{
		{
			name:     "limit less than 2",
			limit:    2,
			expected: 2,
		},
		{
			name:     "limit 3",
			limit:    3,
			expected: 5, // primes: 5
		},
		{
			name:     "limit 10",
			limit:    10,
			expected: 17, // primes: 2+3+5+7
		},
		{
			name:     "limit 20",
			limit:    20,
			expected: 77, // primes: 2+3+5+7+11+13+17+19
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumPrimes(tt.limit)
			if result != tt.expected {
				t.Errorf("SumPrimes(%d) = %d, want %d",
					tt.limit, result, tt.expected)
			}
		})
	}
}
