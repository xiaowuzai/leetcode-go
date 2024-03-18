package leetcode

import (
	"testing"
)

func TestDefangIPaddr2(t *testing.T) {
	tests := []struct {
		name     string
		address  string
		expected string
	}{
		{
			name:     "Test case 1",
			address:  "192.168.0.1",
			expected: "192[.]168[.]0[.]1",
		},
		{
			name:     "Test case 2",
			address:  "10.0.0.1",
			expected: "10[.]0[.]0[.]1",
		},
		{
			name:     "Test case 3",
			address:  "172.16.254.1",
			expected: "172[.]16[.]254[.]1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := defangIPaddr2(test.address)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
