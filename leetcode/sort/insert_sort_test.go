package sort

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	tests := []struct {
		data []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{2, 4, 1, 3, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		data := make([]int, len(tt.data)) // 创建一个与tt.data相同大小的切片
		copy(data, tt.data)               // 将tt.data的值复制到data中

		// 因为是原地排序，所以会修改原切片，所以不能使用tt.data
		got := insertSort(tt.data)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("insertSort(%v) = %v, want %v", data, got, tt.want)
		}
	}
}
