package deleteduplicates

import (
	"reflect"
	"testing"

	"github.com/xiaowuzai/leetcode-go/leetcode/types"
)

func TestDeleteDuplicates(t *testing.T) {
	cases := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{1, 1, 2},
			want:  []int{1, 2},
		},
		{
			input: []int{1, 1, 2, 3, 3},
			want:  []int{1, 2, 3},
		},
	}

	for _, c := range cases {
		got := deleteDuplicates(types.NewList(c.input))
		gotVals := got.Values()
		if !reflect.DeepEqual(gotVals, c.want) {
			t.Errorf("got %v, want %v", gotVals, c.want)
		}
	}
}
