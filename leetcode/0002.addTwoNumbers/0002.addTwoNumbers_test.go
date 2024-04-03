package addtwonumbers

import (
	"reflect"
	"testing"

	"github.com/xiaowuzai/leetcode-go/leetcode/types"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{
			[]int{2, 4, 3},
			[]int{5, 6, 4},
			[]int{7, 0, 8},
		},
		{
			[]int{0},
			[]int{0},
			[]int{0},
		},
		{
			[]int{9, 9, 9, 9, 9, 9, 9},
			[]int{9, 9, 9, 9},
			[]int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		l1 := types.NewList(tt.l1)
		l2 := types.NewList(tt.l2)

		got := addTwoNumbers(l1, l2)
		gotVals := got.Values()
		if !reflect.DeepEqual(gotVals, tt.want) {
			t.Errorf("mergeTwoLists = %v; want %v", gotVals, tt.want)
		}
	}
}

func TestAddTwoNumbers_dg(t *testing.T) {
	tests := []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{
			[]int{2, 4, 3},
			[]int{5, 6, 4},
			[]int{7, 0, 8},
		},
		{
			[]int{0},
			[]int{0},
			[]int{0},
		},
		{
			[]int{9, 9, 9, 9, 9, 9, 9},
			[]int{9, 9, 9, 9},
			[]int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		l1 := types.NewList(tt.l1)
		l2 := types.NewList(tt.l2)

		got := addTwoNumbers_dg(l1, l2)
		gotVals := got.Values()
		if !reflect.DeepEqual(gotVals, tt.want) {
			t.Errorf("mergeTwoLists = %v; want %v", gotVals, tt.want)
		}
	}
}
