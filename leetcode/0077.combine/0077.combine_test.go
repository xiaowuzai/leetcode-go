package combine

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want [][]int
	}{
		// Test cases
		{
			n:    4,
			k:    2,
			want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			n:    5,
			k:    3,
			want: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}},
		},
	}

	for _, tt := range tests {
		got := combine(tt.n, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("combine(%v, %v) = %v, want %v", tt.n, tt.k, got, tt.want)
		}
	}
}
