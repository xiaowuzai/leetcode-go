package combinationsum2

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	n := len(candidates)
	if n == 0 {
		return res
	}

	// 保持升序
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	var backtrack func(start, sum int, path []int)
	backtrack = func(start, sum int, path []int) {
		if sum == target {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := start; i < n; i++ {
			if sum+candidates[i] > target {
				break
			}
			path = append(path, candidates[i])      // 加入路径
			backtrack(i+1, sum+candidates[i], path) // 从i+1开始，避免重复
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0, []int{})
	return res
}
