package sort

/*
插入排序：
 1. 将第一个元素看成有序序列，从第二个元素开始，逐个插入到有序序列中。
    时间复杂度 O(N2)
    空间复杂度 O(1)
*/
func insertSort(data []int) []int {
	n := len(data)
	if n <= 1 {
		return data
	}

	for i := 0; i < n; i++ { // i 之前的位置是有序的
		for j := i + 1; j > 0 && j < n; j-- { // 当前元素小于前一个元素，交换
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			} else { // 没有交换元素，就说明当前元素大于等于之前的元素
				break
			}
		}
	}
	return data
}
