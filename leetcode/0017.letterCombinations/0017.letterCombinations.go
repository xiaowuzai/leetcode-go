package lettercombinations

var m = map[byte][]byte{
	'0': {},
	'1': {},
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return []string{}
	}
	// 记录返回值
	res := make([]string, 0, n)

	// index 记录当前索引，path 记录当前路径
	var backtrack func(index int, path []byte)
	backtrack = func(index int, path []byte) {
		// 如果当前索引的值等于 n， 则打到递归结束条件
		if index == n {
			res = append(res, string(path))
			return
		}

		// 开始做选择
		curNum := digits[index] // 当前数字的字节。比如 2,3,4,
		cn := len(m[curNum])    // 当前数字对应字母的数量
		for i := 0; i < cn; i++ {
			// 把当前数字对应的字母加入 path
			path = append(path, m[curNum][i])
			// 进入下一个节点
			backtrack(index+1, path)
			//  撤销选择
			path = path[:len(path)-1]
		}

	}
	// 从索引为 0 开始记录，
	backtrack(0, []byte{})

	return res
}
