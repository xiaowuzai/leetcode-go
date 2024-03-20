package mst1615mastermind

import (
	"reflect"
	"testing"
)

func TestMasterMind(t *testing.T) {
	tests := []struct {
		solution string
		guess    string
		expected []int
	}{
		{"ABCD", "ABCD", []int{4, 0}},
		{"ABCD", "ABCA", []int{3, 0}},
		{"ABCD", "DCBA", []int{0, 4}},
		{"ABCD", "AABB", []int{1, 1}},
		{"ABCD", "AAAA", []int{1, 0}},
	}

	for _, test := range tests {
		result := masterMind(test.solution, test.guess)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("masterMind(%s, %s) = %v, expected %v", test.solution, test.guess, result, test.expected)
		}
	}
}
