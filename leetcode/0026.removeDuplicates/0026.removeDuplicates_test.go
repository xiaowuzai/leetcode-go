package removeduplicates

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
	}

	for _, tt := range tests {
		got := removeDuplicates(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("removeDuplicates(%v) = %v; want %v", tt.nums, got, tt.want)
		}
	}
}
