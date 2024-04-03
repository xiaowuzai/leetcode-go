package combinationsum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)

	// 升序
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	if n == 0 || candidates[0] > target {
		return [][]int{}
	}
	res := make([][]int, 0)

	// start 表示当前索引
	// sum 表示当前总和
	// path 表示当前路径
	var backtrace func(start, sum int, path []int)
	backtrace = func(start, sum int, path []int) {
		if sum == target {
			res = append(res, append([]int{}, path...))
			return
		}
		// 做选择
		for i := start; i < n; i++ {
			// 剪枝
			tmp := sum + candidates[i]
			if tmp > target {
				break
			}
			path = append(path, candidates[i])
			backtrace(i, tmp, path)
			path = path[:len(path)-1] // 撤销选择
		}
	}

	backtrace(0, 0, []int{})

	return res
}
