package lengthoflastword

/*
方法 1：从后向前遍历字符串，一开始遇到空格则记录， 之后遇到空格则返回

	时间复杂度 O(N),空间复杂度 O(1)
*/
func lengthOfLastWord(s string) int {
	charNum := 0

	n := len(s) - 1
	for i := n; i >= 0; i-- {
		if s[i] == ' ' {
			if charNum != 0 {
				return charNum
			}
			continue
		}
		charNum++
	}
	return charNum
}
