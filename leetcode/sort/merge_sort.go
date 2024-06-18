package sort

func mergeSort(data []int) []int {
	mergeSortHelp(data, 0, len(data)-1)
	return data
}

func mergeSortHelp(data []int, l, r int) {
	if l >= r {
		return
	}

	mid := l + (r-l)/2

	// 注意：这里需要先排序右半部分，再排序左半部分
	mergeSortHelp(data, l, mid)
	mergeSortHelp(data, mid+1, r)
	merge(data, l, mid, r)
}

func merge(data []int, l, mid, r int) {

	temp := make([]int, r-l+1) // 临时数组存储数据
	tempIndex := 0
	p2 := mid + 1
	p1 := l
	for p1 <= mid && p2 <= r {
		if data[p1] <= data[p2] {
			temp[tempIndex] = data[p1]
			tempIndex++
			p1++
		} else {
			temp[tempIndex] = data[p2]
			tempIndex++
			p2++
		}
	}

	for p1 <= mid {
		temp[tempIndex] = data[p1]
		tempIndex++
		p1++
	}

	for p2 <= r {
		temp[tempIndex] = data[p2]
		tempIndex++
		p2++
	}

	copy(data[l:r+1], temp)
}
