package sort

/*
快速排序：

	查找分区点：让右侧的数据大于分区点，左侧的数据小于分区点，不满足就交换数据。返回分区点索引
	分治：递归调用左右分区

	时间复杂度 O(logN) 空间复杂度 O(1) 。不考虑递归调用堆栈
*/
func quickSort(data []int) []int {
	quickSortHelp(data, 0, len(data)-1)
	return data
}

func quickSortHelp(data []int, left, right int) {
	// 基线条件：当左边的索引大于等于右边的索引时，不需要再排序了
	if left >= right {
		return
	}

	p := partition(data, left, right) // 获取分区的中点
	quickSortHelp(data, left, p-1)    // 对左分区进行排序
	quickSortHelp(data, p+1, right)   // 对右分区进行排序
}

func partition(data []int, left, right int) int {
	pivot := data[left] // 选择第一个元素作为基准值。 这里需要注意。data[left] 是可以被覆盖的
	// 需要满足，左侧 的元素小于基准值，右侧的元素大于基准值
	for left < right {
		// 找右侧小于等于基准值的元素
		for left < right && data[right] > pivot {
			right--
		}

		// 找左侧大于等于基准值的元素
		for left < right && data[left] < pivot {
			left++
		}
		// 交换两个元素
		data[right], data[left] = data[left], data[right]
	}

	// 基准值归位
	data[left] = pivot // 此时 left 已经挪动到中间位置了，所以要归位
	return left
}
