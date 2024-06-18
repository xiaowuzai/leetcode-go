package canpartition

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	// 奇数则不可能
	if sum%2 != 0 {
		return false
	}

	n := len(nums)
	want := sum / 2
	// 01背包问题
	dp := make([][]bool, n+1) // dp[i][j] 表示前i个物品是否能装满容量为j的背包
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, want+1)
		dp[i][0] = true // 背包为空表示满了
	}

	// 从起始位置开始规划，一直到想要的容量
	for i := 1; i <= n; i++ { // 总共有n个物品
		for j := 1; j <= want; j++ { // 容量也从 1 开始，直到 want
			if j < nums[i-1] { // 当前容量小于当前物品容量
				dp[i][j] = dp[i-1][j] // 不装当前物品
			} else {
				// 如果前一个容量的背包满了
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]] // 装当前物品或者不装当前物品
			}
		}
	}
	return dp[n][want] // 返回 n 对应的 want 的结果
}
