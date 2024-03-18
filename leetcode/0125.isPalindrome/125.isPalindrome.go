package ispalindrome

import "strings"

func isPalindrome(s string) bool {
	// 转成小写字符串
	s = strings.ToLower(s)
	i, j := 0, len(s)-1 // 从前往后，从后往前

	for i < j {
		// 不是有效字符就跳过
		for i < j && !validChar(s[i]) {
			i++
		}
		for i < j && !validChar(s[j]) {
			j--
		}

		// 有效字符不相等，则不是回文串
		if i < j && s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 有效字符的范围是小写 a ~ z 和 0 ~ 9
func validChar(b byte) bool {
	if (b >= 'a' && b <= 'z') ||
		(b >= '0' && b <= '9') {
		return true
	}
	return false
}
