package types

import (
	"reflect"
	"testing"
)

func TestNewListAddVals(t *testing.T) {
	tests := []struct {
		vals []int
		want []int
		// want *ListNode
	}{
		{
			vals: []int{1, 2, 3},
			// want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			want: []int{1, 2, 3},
		},
		{
			vals: []int{4, 5, 6, 7},
			// want: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7}}}},
			want: []int{4, 5, 6, 7},
		},
		{
			vals: []int{},
			want: []int{},
			// want: &ListNode{},
		},
	}

	for _, tt := range tests {
		got := NewList(tt.vals)
		gotVals := got.Values()

		if !reflect.DeepEqual(gotVals, tt.want) {
			t.Errorf("NewList(%v) = %v, want %v", tt.vals, gotVals, tt.want)
		}

	}
}
