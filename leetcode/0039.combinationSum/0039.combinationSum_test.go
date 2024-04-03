package combinationsum

import (
	"testing"
)

func TestCombinationSum(t *testing.T) {
	candidates := []int{7, 3, 9, 6}
	target := 6
	expected := [][]int{
		{3, 3},
		{6},
	}
	result := combinationSum(candidates, target)

	// 检查结果是否和预期一致
	if len(result) != len(expected) {
		t.Errorf("Expected %v combinations, but got %v", len(expected), len(result))
	}
	for i := 0; i < len(result); i++ {
		if len(result[i]) != len(expected[i]) {
			t.Errorf("Expected combination %v, but got %v", expected[i], result[i])
		}
		for j := 0; j < len(result[i]); j++ {
			if result[i][j] != expected[i][j] {
				t.Errorf("Expected combination %v, but got %v", expected[i], result[i])
				break
			}
		}
	}
}
