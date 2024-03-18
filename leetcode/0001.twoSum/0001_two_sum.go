package leetcode

// 暴力破解：时间复杂度 O(N2),空间复杂度 O(1)
func twoSum1(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 借用 map 存储访问过的元素及其索引
func twoSum2(nums []int, target int) []int {
	n := len(nums)
	m := make(map[int]int, n) // 数值及索引
	for i := 0; i < n; i++ {
		complement := target - nums[i] // 判断答案是否在 map 中
		if _, has := m[complement]; has {
			return []int{m[complement], i}
		}
		m[nums[i]] = i // 存储当前值和索引
	}
	return []int{}
}
