package dailytemperatures

/*
方法一： 暴力破解 双层 for 循环。时间复杂度 O(N2),空间复杂度 O(1)

	方法二： 单调栈 时间复杂度 O(N),空间复杂度 O(N)
*/
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := make([]int, 0, n) // 模拟栈，保存元素索引

	for i := 0; i < n; i++ {
		// 如果栈为空，或者当前元素小于等于栈顶元素，则将索引入栈
		if len(stack) == 0 || temperatures[i] <= temperatures[stack[len(stack)-1]] {
			stack = append(stack, i)
			continue
		}

		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			// 出栈，并记录天数
			index := stack[len(stack)-1]
			temperatures[index] = i - index
			stack = stack[:len(stack)-1] // 弹出栈顶元素
		}

		// 将当前元素入栈。这时栈内的元素是单调递减的
		stack = append(stack, i)
	}

	for _, index := range stack {
		temperatures[index] = 0 // 栈内元素对应的天数为0
	}
	return temperatures
}
