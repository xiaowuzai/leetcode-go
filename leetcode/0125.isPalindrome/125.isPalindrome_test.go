package ispalindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"", true},
		{"level", true},
		{"noon", true},
		{"hello", false},
	}

	for _, tc := range testCases {
		got := isPalindrome(tc.input)
		if got != tc.want {
			t.Errorf("isPalindrome(%s) = %v; want %v", tc.input, got, tc.want)
		}
	}
}
