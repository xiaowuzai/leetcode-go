package majorityelement

// 1. 使用 map 来计数,时间复杂度 O(N)，空间复杂度 O(N)
// 2. 通过 计数来判断，时间复杂度 O(N)，空间复杂度 O(1)
func majorityElement_map(nums []int) int {
	count := make(map[int]int, len(nums))
	for _, num := range nums {
		count[num]++
		if count[num] > len(nums)/2 {
			return num
		}
	}
	return 0
}

func majorityElement_count(nums []int) int {
	var (
		count = 0
		num   = 0
		n     = len(nums)
	)

	if n == 0 {
		return 0
	}

	num = nums[0]
	count++

	for i := 1; i < n; i++ {
		if nums[i] == num {
			count++
		} else if count == 0 {
			num = nums[i]
			count++
		} else {
			count--
		}
	}

	return num
}
