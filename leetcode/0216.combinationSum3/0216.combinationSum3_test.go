package combinationsum3

import (
	"reflect"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	tests := []struct {
		k    int
		n    int
		want [][]int
	}{
		// Test cases
		{
			k:    3,
			n:    9,
			want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			k: 4,
			n: 16,
			want: [][]int{
				{1, 2, 4, 9},
				{1, 2, 5, 8},
				{1, 2, 6, 7},
				{1, 3, 4, 8},
				{1, 3, 5, 7},
				{1, 4, 5, 6},
				{2, 3, 4, 7},
				{2, 3, 5, 6},
			},
		},
		{
			k:    2,
			n:    10,
			want: [][]int{{1, 9}, {2, 8}, {3, 7}, {4, 6}},
		},
	}

	for _, tt := range tests {
		got := combinationSum3(tt.k, tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("combinationSum3(%v, %v) = %v; want %v", tt.k, tt.n, got, tt.want)
		}
	}
}
