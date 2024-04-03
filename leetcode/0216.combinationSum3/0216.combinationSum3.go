package combinationsum3

/*
回溯：
*/
func combinationSum3(k int, n int) [][]int {
	/*
		k 个数相加起来等于 n
	*/
	res := make([][]int, 0)

	path := make([]int, 0, k) //记录中间过程
	// start 表示当前起始位置
	// sum 表示当前路径的和
	// path 表示当前路径
	var backtrack func(start, sum int, path []int)
	backtrack = func(start, sum int, path []int) {
		if len(path) == k {
			if sum == n {
				res = append(res, append([]int{}, path...))
			}
			return
		}

		for i := start; i <= 9; i++ {
			if sum+i > n {
				break
			}
			path = append(path, i)
			backtrack(i+1, sum+i, path)
			path = path[:len(path)-1]
		}
	}
	backtrack(1, 0, path)
	return res
}
