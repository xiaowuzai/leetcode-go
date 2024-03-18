package removeduplicates

/*
方法一：双指针

	时间复杂度 O(n)
	 空间复杂度 O(1)
*/
func removeDuplicates(nums []int) int {

	// 空数组直接返回 0
	if len(nums) == 0 {
		return 0
	}

	// 让 i 指向第一个不重复的元素， j 用来遍历数组
	i, j := 0, 0
	n := len(nums)
	for j < n {
		// 由于是递增序列，查找到第一个大于等于 nums[i] 的元素
		for j < n && nums[j] <= nums[i] {
			j++
		}
		// 如果 j >= n，说明 nums[j] 之后的元素全部都是重复的，直接返回 i+1 即可
		if j >= n {
			break
		}

		// 移动位置并赋值
		i++
		nums[i] = nums[j]
	}
	// 返回不重复元素的个数
	return i + 1
}
