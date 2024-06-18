package lcr101canpartition

import (
	"testing"
)

func TestCanPartition(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		// {[]int{1, 5, 11, 5}, true},
		// {[]int{2, 3, 5}, true},
		// {[]int{2, 2, 2, 2}, true},
		{[]int{1, 2, 5}, false},
		// {[]int{1, 2, 3, 4, 5}, false},
	}

	for _, tt := range tests {
		if got := canPartition(tt.nums); got != tt.want {
			t.Errorf("canPartition(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
