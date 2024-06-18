package lcr101canpartition

/*
分割等和子集

	子集问题一般用动态规划的方法来解决。
	需要确定 dp 数组的定义 dp[i][j] 表示, 从 nums[0...i] 中挑选一些正整数, 是否存在一种组合使得它们的和等于 j。
*/
func canPartition(nums []int) bool {
	n := len(nums) // 计算数量
	sum := 0       // 记录总和
	for _, v := range nums {
		sum += v
	}
	if n == 1 || sum%2 != 0 { // 和为奇数则不可能平分区间
		return false
	}

	sum = sum / 2 // 计算目标和

	// 初始化 dp 数组
	dp := make([][]bool, n+1) // dp[]数组的大小.因为我们要返回是 dp[n][sum] 的值。所以dp 空间要加一
	for i := 0; i <= n; i++ { // 遍历 n+1 次
		dp[i] = make([]bool, sum+1) // 初始化 dp[i] 数组.容量也要+1
		dp[i][0] = true             // 因为背包的容量为 0，则表示已经满了
	}

	// 这里有一个细节。dp[0][j] 这种默认是 false 的。因为一个元素都没有，背包不可能满
	for i := 1; i <= n; i++ { // 外层循环遍历  所有 nums
		for j := 1; j <= sum; j++ { // 内层遍历背包的容量
			cur := nums[i-1] // 当前元素
			if j < cur {     // 如果背包容量小于当前元素，则不能装入当前元素
				dp[i][j] = dp[i-1][j] //不将当前元素放入背包，所以要获取上一个元素的状态
			} else {
				// 将当前元素放入背包，或者不将当前元素放入背包.
				dp[i][j] = dp[i-1][j] || dp[i-1][j-cur]
			}
		}
	}

	return dp[n][sum]
}
