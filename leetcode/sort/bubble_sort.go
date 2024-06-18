package sort

func bubbleSort(arr []int) []int {

	n := len(arr)
	for i := 0; i < n; i++ {
		isSorted := true // 标记本次是否已经有序
		for j := 0; j < n-i-1; j++ {
			// 保证内层循环之后，最大值已经确定
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSorted = false // 有元素交换，则表示当前还是无序的
			}
		}
		// 如果已经有序了，则直接跳出循环
		if isSorted {
			break
		}
	}
	return arr
}
