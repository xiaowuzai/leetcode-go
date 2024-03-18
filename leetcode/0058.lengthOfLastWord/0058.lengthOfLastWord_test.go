package lengthoflastword

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"Hello World", 5},
		{"Hello", 5},
		{"   Hello   ", 5},
		{"", 0},
		{"World", 5},
	}

	for _, testCase := range testCases {
		result := lengthOfLastWord(testCase.input)
		if result != testCase.expected {
			t.Errorf("lengthOfLastWord(%s) = %d; want %d", testCase.input, result, testCase.expected)
		}
	}
}
