package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		result []int
	}{
		{
			name:   "Empty array",
			arr:    []int{},
			result: []int{},
		},
		{
			name:   "Array with one element",
			arr:    []int{5},
			result: []int{5},
		},
		{
			name:   "Sorted array",
			arr:    []int{1, 2, 3, 4, 5},
			result: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "Unsorted array",
			arr:    []int{5, 3, 1, 4, 2},
			result: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bubbleSort(tt.arr); !reflect.DeepEqual(got, tt.result) {
				t.Errorf("bubbleSort() = %v, want %v", got, tt.result)
			}
		})
	}
}
