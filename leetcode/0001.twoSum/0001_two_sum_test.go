package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{2, 3, 4},
			target: 6,
			want:   []int{0, 2},
		},
		{
			nums:   []int{-1, 0},
			target: -1,
			want:   []int{0, 1},
		},
		{
			nums:   []int{},
			target: 5,
			want:   []int{},
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			target: 10,
			want:   []int{},
		},
	}

	for _, test := range tests {
		got := twoSum1(test.nums, test.target)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("twoSum1(%v, %d) = %v; want %v", test.nums, test.target, got, test.want)
		}
	}
}

func TestTwoSum2(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{2, 3, 4},
			target: 6,
			want:   []int{0, 2},
		},
		{
			nums:   []int{-1, 0},
			target: -1,
			want:   []int{0, 1},
		},
		{
			nums:   []int{},
			target: 5,
			want:   []int{},
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			target: 10,
			want:   []int{},
		},
	}

	for _, test := range tests {
		got := twoSum2(test.nums, test.target)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("twoSum2(%v, %d) = %v; want %v", test.nums, test.target, got, test.want)
		}
	}
}
