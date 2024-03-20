package removeelements

import (
	"reflect"
	"testing"

	"github.com/xiaowuzai/leetcode-go/leetcode/types"
)

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		vals []int
		val  int
		want []int
	}{
		{
			vals: []int{1, 2, 6, 3, 4, 5, 6},
			val:  6,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			vals: []int{},
			val:  1,
			want: []int{},
		},
		{
			vals: []int{7, 7, 7, 7},
			val:  7,
			want: []int{},
		},
	}

	for _, tt := range tests {
		head := types.NewList(tt.vals)
		got := removeElements(head, tt.val)
		gotVals := got.Values()

		if !reflect.DeepEqual(gotVals, tt.want) {
			t.Errorf("NewList(%v) = %v, want %v", tt.vals, gotVals, tt.want)
		}

	}
}

func TestRemoveElements_2(t *testing.T) {
	tests := []struct {
		vals []int
		val  int
		want []int
	}{
		{
			vals: []int{1, 2, 6, 3, 4, 5, 6},
			val:  6,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			vals: []int{},
			val:  1,
			want: []int{},
		},
		{
			vals: []int{7, 7, 7, 7},
			val:  7,
			want: []int{},
		},
	}

	for _, tt := range tests {
		head := types.NewList(tt.vals)
		got := removeElements_2(head, tt.val)
		gotVals := got.Values()

		if !reflect.DeepEqual(gotVals, tt.want) {
			t.Errorf("NewList(%v) = %v, want %v", tt.vals, gotVals, tt.want)
		}

	}
}
