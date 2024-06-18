package climbstairs

import (
	"testing"
)

func TestClimbStairs_dg(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		// Add more test cases here
	}
	for _, test := range tests {
		got := climbStairs_dg(test.n)
		if got != test.want {
			t.Errorf("climbStairs_dg(%d) = %d; want %d", test.n, got, test.want)
		}
	}
}

func TestClimbStairs_dg_yh(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		// Add more test cases here
	}
	for _, test := range tests {
		got := climbStairs_dg_yh(test.n)
		if got != test.want {
			t.Errorf("climbStairs_dg(%d) = %d; want %d", test.n, got, test.want)
		}
	}
}
func TestClimbStairs_bl(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		// Add more test cases here
	}
	for _, test := range tests {
		got := climbStairs_bl(test.n)
		if got != test.want {
			t.Errorf("climbStairs_dg(%d) = %d; want %d", test.n, got, test.want)
		}
	}
}
func TestClimbStairs_bl_yh(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		// Add more test cases here
	}
	for _, test := range tests {
		got := climbStairs_bl_yh(test.n)
		if got != test.want {
			t.Errorf("climbStairs_dg(%d) = %d; want %d", test.n, got, test.want)
		}
	}
}
