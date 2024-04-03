package lengthoflis

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{3, 4, 5, 8, 10}, 5},
		{[]int{1, 3, 5, 4, 7}, 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := lengthOfLIS_2(tt.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
