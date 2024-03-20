package mst1615mastermind

/*
使用 数组 对 solution 的颜色进行计数
*/
func masterMind(solution string, guess string) []int {
	// 统计 solution 出现的颜色数
	var arr [26]byte // 26 个字母

	guessOk, guessNot := 0, 0
	for i := 0; i < 4; i++ {
		if solution[i] == guess[i] {
			guessOk++
			continue
		}
		arr[solution[i]-'A']++
	}

	for i := 0; i < 4; i++ {
		if solution[i] == guess[i] {
			continue
		}
		n := arr[guess[i]-'A']
		if n > 0 {
			guessNot++
			arr[guess[i]-'A']--
		}
	}
	return []int{guessOk, guessNot}
}
