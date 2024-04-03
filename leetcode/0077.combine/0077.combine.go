package combine

/*
组合、排列、子集的问题一般会用回溯算法来解决
*/
func combine(n int, k int) [][]int {
	// 返回值
	res := make([][]int, 0, n)
	// path := make([]int, 0, k)

	// 递归函数.记录走过的路径
	var backtrack func(start int, path []int)

	backtrack = func(start int, path []int) {
		// 如果当前路径长度等于k，说明找到了一个组合，将当前路径加入结果集
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
			return
		}

		// 同一个 path 里面不允许出现重复元素，所以要访问下一个节点
		for i := start; i <= n; i++ {
			path = append(path, i) // 追加路径
			backtrack(i+1, path)
			path = path[:len(path)-1] // 撤销选择
		}
	}
	backtrack(1, []int{})

	return res
}
