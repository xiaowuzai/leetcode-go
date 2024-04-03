package mergetwolists

import (
	"reflect"
	"testing"

	"github.com/xiaowuzai/leetcode-go/leetcode/types"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{
			[]int{1, 2, 4},
			[]int{1, 3, 4},
			[]int{1, 1, 2, 3, 4, 4},
		},
		{
			[]int{},
			[]int{},
			[]int{},
		},
		{
			[]int{},
			[]int{0},
			[]int{0},
		},
	}

	for _, tt := range tests {
		l1 := types.NewList(tt.l1)
		l2 := types.NewList(tt.l2)

		got := mergeTwoLists(l1, l2)
		gotVals := got.Values()
		if !reflect.DeepEqual(gotVals, tt.want) {
			t.Errorf("mergeTwoLists = %v; want %v", gotVals, tt.want)
		}
	}
}
