package sort

func selectSort(data []int) []int {
	n := len(data)
	if n <= 1 {
		return data
	}

	for i := 0; i < n; i++ { // i 指向已排序区间的末尾
		min := data[i]
		minIndex := i
		for j := i; j < n; j++ { // 找到最小元素，交换
			if data[j] < min {
				min = data[j]
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
	return data
}
