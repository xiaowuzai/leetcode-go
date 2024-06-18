package removeduplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums1 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	expected1 := 7
	result1 := removeDuplicates(nums1)
	if result1 != expected1 {
		t.Errorf("Expected %d, but got %d", expected1, result1)
	}

	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected2 := 9
	result2 := removeDuplicates(nums2)
	if result2 != expected2 {
		t.Errorf("Expected %d, but got %d", expected2, result2)
	}
}
