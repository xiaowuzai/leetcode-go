package palindromenumber

import "strconv"

/*
方法 1 ： 将数字转化成字符串，然后使用双指针判断是否为回文串

	时间复杂度: O(n)
	空间复杂度: O(n)
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

/*
方法 2 ：将数字存储数组，判断数组中的值是否对称
	时间复杂度: O(n)
	空间复杂度: O(n)
*/

func isPalindrome2(x int) bool {
	// 负数一定不是回文数字
	if x < 0 {
		return false
	}

	arr := make([]int, 0, 10) // 数组存储元素
	for x > 0 {
		// 对 10 取余，也就是把末尾的数字放入数组中
		arr = append(arr, x%10)
		x = x / 10
	}

	i, j := 0, len(arr)-1
	for i < j {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
	return true
}
