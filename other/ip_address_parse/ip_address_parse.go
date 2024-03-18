package other

import (
	"strconv"
	"strings"
)

func isIPv4(ip string) bool {

	contentMap := ipContent()
	for i := 0; i < len(ip); i++ {
		if !isContent(contentMap, string(ip[i])) {
			return false
		}
	}

	numbers := strings.Split(ip, ".")
	if len(numbers) != 4 {
		return false
	}
	for _, numStr := range numbers {
		if hasZeroPrefix(numStr) {
			return false
		}

		num, err := strconv.Atoi(trimSpace(numStr))
		if err != nil || !checkIPAddress(num) {
			return false
		}
	}

	return true
}

/*
checkIPAddress 函数用于检查一个整数是否为有效的IP地址数字。

参数:
i int - 待检查的整数。

返回值:
bool - 如果整数在0到255之间（含0和255），返回true，表示是有效的IP地址数字；否则返回false。
*/
func checkIPAddress(i int) bool {
	// 检查整数是否在0到255的范围内
	if i < 0 || i > 255 {
		return false
	}
	return true
}

// ipContent 函数生成一个包含特定字符串键的映射。
// 该映射中的键包括空格、句点以及数字0到9。
//
// 返回值:
// 返回一个 map[string]struct{} 类型的变量，其中键是字符串，值为空结构体。
// 这种映射类型的选择是因为只关心键的存在，而不关心值的具体内容。
func ipContent() map[string]struct{} {
	// 初始化映射 m，包含空格和句点两个键。
	m := map[string]struct{}{
		" ": {},
		".": {},
	}

	// 循环将数字0到9作为键添加到映射 m 中。
	for i := 0; i < 10; i++ {
		m[strconv.Itoa(i)] = struct{}{}
	}

	return m
}

func isContent(m map[string]struct{}, s string) bool {
	_, has := m[s]
	return has
}

func trimSpace(str string) string {
	return strings.TrimSuffix(strings.TrimPrefix(str, " "), " ")
}

func hasZeroPrefix(str string) bool {
	return strings.Contains(str, "0") && len(str) > 1
}
