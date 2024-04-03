package lengthoflis

/*
求最值问题，一般用动态规划来解决
时间复杂度 O(N^2) 空间复杂度 O(N)
*/
func lengthOfLIS(nums []int) int {
	n := len(nums)
	// 初始化一个 dp table
	dp := make([]int, n)
	// 由于最长递增子序列，最少包含自己，所以元素个数为 1
	for i := range dp {
		dp[i] = 1
	}

	// 有一个规律，当求解 dp[i] 应该存放的元素个数时，只需要找到小于 nums[i]的元素, 计数的最大值加一
	// 例如 nums[i] 的元素为 8，我们只需要向前遍历，找到所有小于 8 的元素中 dp数组中的最大值，然后计数+1
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			// 找到小于当前元素的值
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1) // 找到最大值存入
			}
		}
	}

	// 找到最大值
	max := 0
	for _, v := range dp {
		if max < v {
			max = v
		}
	}
	return max
}

/*
优化：

	通过分组的方式来获取最长递增子序列
*/
func lengthOfLIS_2(nums []int) int {

	n := len(nums)
	dp := make([]int, n) // 记录索引位置的最顶层元素
	// 初始组为 0
	group := 0

	for i := 0; i < n; i++ {

		left, right := 0, group // 确定左右边界

		// 二分查找.确定最左边界， 最左边界就是元素要放的组
		for left < right {
			mid := left + (right-left)/2
			if dp[mid] >= nums[i] { // 用组内顶部元素进行比较
				right = mid
			} else {
				left = mid + 1
			}
		}

		// 没找到要分的组
		if left == group {
			group++
		}

		// 将当前元素放入最左边界所在的组
		dp[left] = nums[i]
	}

	return group
}
