package findcontentchildren

import (
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	// Test case 1
	g1 := []int{1, 2, 3}
	s1 := []int{1, 1, 2, 3, 3}
	expected1 := 3
	result1 := findContentChildren(g1, s1)
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	}

	// Test case 2
	g2 := []int{2, 3, 4}
	s2 := []int{1, 1, 2, 3, 3}
	expected2 := 2
	result2 := findContentChildren(g2, s2)
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	}

	// Test case 3
	g3 := []int{1, 2, 3}
	s3 := []int{1, 2}
	expected3 := 2
	result3 := findContentChildren(g3, s3)
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	}

	// Test case 4
	g4 := []int{}
	s4 := []int{1, 2, 3}
	expected4 := 0
	result4 := findContentChildren(g4, s4)
	if result4 != expected4 {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result4)
	}

	// Test case 5
	g5 := []int{1, 2, 3}
	s5 := []int{}
	expected5 := 0
	result5 := findContentChildren(g5, s5)
	if result5 != expected5 {
		t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result5)
	}
}
