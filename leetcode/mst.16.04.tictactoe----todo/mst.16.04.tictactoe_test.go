package mst1604tictactoe

import "testing"

func TestTictactoe(t *testing.T) {
	cases := []struct {
		input    []string
		expected string
	}{
		{
			input:    []string{"O X", " XO", "X O"},
			expected: "X",
		},
		{
			input:    []string{"OOX", "XXO", "OXO"},
			expected: "Draw",
		},
		{
			input:    []string{"OOX", "XXO", "OX "},
			expected: "Pending",
		},
		{
			input:    []string{" "},
			expected: "Pending",
		},
		{
			input:    []string{"X"},
			expected: "X",
		},
	}

	for _, c := range cases {
		actual := tictactoe(c.input)
		if actual != c.expected {
			t.Errorf("tictactoe(%v) returned %v, expected %v", c.input, actual, c.expected)
		}
	}
}
