package middlenode

import (
	"reflect"
	"testing"

	"github.com/xiaowuzai/leetcode-go/leetcode/types"
)

func TestMiddleNode_save(t *testing.T) {
	tests := []struct {
		vals []int
		want []int
	}{
		{
			vals: []int{1, 2, 3, 4, 5},
			want: []int{3, 4, 5},
		},
		{
			vals: []int{1, 2, 3, 4, 5, 6},
			want: []int{4, 5, 6},
		},
		{
			vals: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		head := types.NewList(tt.vals)

		got := middleNode_save(head)
		gotVals := got.Values()
		if !reflect.DeepEqual(gotVals, tt.want) {
			t.Errorf("NewList(%v) = %v, want %v", tt.vals, gotVals, tt.want)
		}
	}
}

func TestMiddleNode_pointer(t *testing.T) {
	tests := []struct {
		vals []int
		want []int
	}{
		{
			vals: []int{1, 2, 3, 4, 5},
			want: []int{3, 4, 5},
		},
		{
			vals: []int{1, 2, 3, 4, 5, 6},
			want: []int{4, 5, 6},
		},
		{
			vals: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		head := types.NewList(tt.vals)

		got := middleNode_pointer(head)
		gotVals := got.Values()
		if !reflect.DeepEqual(gotVals, tt.want) {
			t.Errorf("NewList(%v) = %v, want %v", tt.vals, gotVals, tt.want)
		}
	}
}

func TestMiddleNode_fast(t *testing.T) {
	tests := []struct {
		vals []int
		want []int
	}{
		{
			vals: []int{1, 2, 3, 4, 5},
			want: []int{3, 4, 5},
		},
		{
			vals: []int{1, 2, 3, 4, 5, 6},
			want: []int{4, 5, 6},
		},
		{
			vals: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		head := types.NewList(tt.vals)

		got := middleNode_fast(head)
		gotVals := got.Values()
		if !reflect.DeepEqual(gotVals, tt.want) {
			t.Errorf("NewList(%v) = %v, want %v", tt.vals, gotVals, tt.want)
		}
	}
}
