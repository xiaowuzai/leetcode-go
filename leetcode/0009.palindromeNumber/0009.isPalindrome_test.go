package palindromenumber

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{121, true},
		{-121, false},
		{12321, true},
		{12345, false},
		{0, true},
	}

	for _, tc := range testCases {
		result := isPalindrome(tc.input)
		if result != tc.expected {
			t.Errorf("isPalindrome(%d) = %t, expected %t", tc.input, result, tc.expected)
		}
	}
}

func TestIsPalindrome2(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{121, true},
		{-121, false},
		{12321, true},
		{12345, false},
		{0, true},
	}

	for _, tc := range testCases {
		result := isPalindrome2(tc.input)
		if result != tc.expected {
			t.Errorf("isPalindrome2(%d) = %t, expected %t", tc.input, result, tc.expected)
		}
	}
}
