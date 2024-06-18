package lcr103coinchange

/*
零钱兑换问题
*/
func coinChange(coins []int, amount int) int {
	bad := -1               // 凑不齐
	maxNumber := amount + 1 // 最多个数也就是 amount

	dp := make([]int, maxNumber) // 存储 amount 和 当前凑的钱数需要多少个硬币
	for i := range dp {
		dp[i] = maxNumber // 初始化为最大值
	}

	// 确定状态迁移。
	// 传入 amount 返回 当前凑的钱数需要多少个硬币
	var coinChangeHelp func(amount int) int
	coinChangeHelp = func(amount int) int {
		if amount == 0 {
			return 0 // 不需要凑数
		}

		// 凑不齐
		if amount < 0 {
			return bad
		}

		// 不是第一次计算
		if dp[amount] != maxNumber {
			return dp[amount]
		}

		for _, v := range coins {
			// 求解子问题
			sub := coinChangeHelp(amount - v)

			//子问题无解
			if sub == bad {
				continue
			}

			// 取当前值与子问题的值+1 的最小值
			dp[amount] = min(dp[amount], sub+1)
		}

		// 凑不齐
		if dp[amount] == maxNumber {
			dp[amount] = bad
		}

		return dp[amount]
	}

	return coinChangeHelp(amount)
}

func coinChange1(coins []int, amount int) int {
	max := amount + 1

	dp := make([]int, max) // 存储 amount 和 当前凑的钱数需要多少个硬币
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = max
	}

	// 表示当前金额
	for i := 1; i <= amount; i++ { // 相当于背包容量
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == max {
		return -1
	}

	return dp[amount]
}
