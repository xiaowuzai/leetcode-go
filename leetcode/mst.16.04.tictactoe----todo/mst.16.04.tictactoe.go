package mst1604tictactoe

func tictactoe(board []string) string {
	// 如果为 0，直接返回
	n := len(board)
	if n == 0 {
		return ""
	}
	if n == 1 {
		if board[0] == " " {
			return "Pending"
		}
		return string(board[0][0])
	}

	hasPending := false
	// 判断列是否完全相同
	for i := 0; i < n; i++ {
		// 判断行是否完全相同的元素
		r, pending := isEqual(board[i])
		if r != "" {
			return r
		}
		if pending {
			hasPending = true
		}

	}
	for j := 0; j < n; j++ {
		columnEqual := true
		for i := 1; i < n; i++ {
			if board[i-1][j] == ' ' || board[i][j] == ' ' {
				hasPending = true
				break
			}
			if board[i-1][j] != board[i][j] {
				columnEqual = false
			}

			if i == n-1 && columnEqual {
				return string(board[i][j])
			}
		}
	}

	if hasPending {
		return "Pending"
	}
	return "Draw"
}

// 判断字符串的字符是否完全相同。返回相同的元素， 以及是否有空位
func isEqual(str string) (string, bool) {

	result, pending := string(str[0]), false
	if str[0] == ' ' { // 开头为空，就不用进入循环了
		return "", true
	}

	n := len(str)
	for i := 1; i < n; i++ {
		if str[i] == ' ' {
			pending = true // 有空位置
		}
		if str[i] != str[i-1] {
			result = ""
		}
	}

	return result, pending
}
