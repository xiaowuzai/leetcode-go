package lettercombinations

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		digits string
		want   []string
	}{
		{digits: "", want: []string{}},
		{digits: "23", want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{digits: "2", want: []string{"a", "b", "c"}},
	}

	for _, test := range tests {
		got := letterCombinations(test.digits)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("letterCombinations(%q) = %v, want %v", test.digits, got, test.want)
		}
	}
}
