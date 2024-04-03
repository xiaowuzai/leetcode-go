package coinchange

/*
思考：
*/
func coinChange(coins []int, amount int) int {

	table := make([]int, amount+1) //  记录从 1 到 amount 的最少硬币数

	max := amount + 2              // 设定一个不可能的最大值
	for i := 1; i <= amount; i++ { // 从 1 到 amount 初始化
		table[i] = max
	}

	//定义递归函数，amount 为当前要凑的金额,返回当前金额的最少硬币数，凑不上则返回 max
	var dp func(amount int) int

	dp = func(amount int) int {
		if amount == 0 {
			return 0
		}

		// 子问题无解
		if amount < 0 {
			return max
		}

		if table[amount] != max {
			return table[amount]
		}

		for _, coin := range coins {
			// 求解子问题
			sub := dp(amount - coin)
			table[amount] = min(table[amount], sub+1) // 子问题+1是最优解
		}

		return table[amount]
	}

	less := dp(amount)
	if less >= max {
		return -1
	}
	return less
}
