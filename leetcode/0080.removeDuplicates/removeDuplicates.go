package removeduplicates

// [111222333]
// [112233]
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	// 记录当前值的个数
	count := 1
	l, r := 0, 1
	for l < r && r < n {
		// 元素相等且个数小于2
		if nums[r] == nums[l] && count == 1 {
			count++
			l++
			// 这里需要将元素拷贝过来，因为 r 可能不是 l+1
			nums[l] = nums[r]
		} else if nums[r] > nums[l] {
			l++
			nums[l] = nums[r]
			count = 1
		}
		r++
	}
	return l + 1
}
