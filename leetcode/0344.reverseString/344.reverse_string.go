package reversestring

// 1. 开辟新的字节数组，从后往前遍历并写入。
// 时间复杂度 O(N),空间复杂度 O(N)
func reverseString(s []byte) {
	n := len(s)
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = s[n-1-i]
	}
	copy(s, bs)
}

// 2.原地翻转
// 时间复杂度 O(N)，空间复杂度 O(1)
func reverseString2(s []byte) {
	i, j := 0, len(s)-1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
