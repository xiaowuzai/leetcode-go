package reverselist

import (
	"reflect"
	"testing"

	"github.com/xiaowuzai/leetcode-go/leetcode/types"
)

func TestReverseList(t *testing.T) {
	cases := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		},
		{
			input: []int{1, 2},
			want:  []int{2, 1},
		},
	}

	for _, c := range cases {
		got := reverseList(types.NewList(c.input))
		gotVals := got.Values()
		if !reflect.DeepEqual(gotVals, c.want) {
			t.Errorf("got %v, want %v", gotVals, c.want)
		}
	}
}
