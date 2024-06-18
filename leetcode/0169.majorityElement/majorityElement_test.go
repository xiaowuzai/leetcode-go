package majorityelement

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMajorityElement_map(t *testing.T) {
	tdd := []struct {
		nums   []int
		expect int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for _, tc := range tdd {
		result := majorityElement_map(tc.nums)
		require.Equal(t, tc.expect, result)
	}
}

func TestMajorityElement_count(t *testing.T) {
	tdd := []struct {
		nums   []int
		expect int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for _, tc := range tdd {
		result := majorityElement_count(tc.nums)
		require.Equal(t, tc.expect, result)
	}
}
