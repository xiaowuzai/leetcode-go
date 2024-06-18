package canjump

func canJump(nums []int) bool {

	n := len(nums)
	end := n - 1    // 需要到达的最远位置
	result := false // 当前位置是否能到达

	// 从后往前遍历
	for i := n - 1; i >= 0; i-- {
		if i+nums[i] >= end { // 索引加跳数大于等于能到达的最远位置
			result = true // 当前位置可以到达
			end = i       // 修改最远位置
		} else {
			result = false // 不能到达
		}
	}

	// 返回结果
	return result
}
