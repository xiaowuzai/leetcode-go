package lcr103coinchange

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	data := []struct {
		code     int
		coins    []int
		amount   int
		expected int
	}{
		{

			code:     1,
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
		{
			code:     2,
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			code:     3,
			coins:    []int{},
			amount:   5,
			expected: -1,
		},
		{
			code:     4,
			coins:    []int{1, 2, 5},
			amount:   0,
			expected: 0,
		},
		{
			code:     5,
			coins:    []int{1, 2, 5},
			amount:   6,
			expected: 2,
		},
	}

	for _, tt := range data {
		result := coinChange1(tt.coins, tt.amount)

		if result != tt.expected {
			t.Errorf("code = %d Expected %v but got %v", tt.code, tt.expected, result)
		}
	}

}
