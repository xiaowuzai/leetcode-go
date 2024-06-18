package climbstairs

/*
	分析题意：
		当 n = 1 时，只有一种路径. 1
		当 n = 2 时，有两种路径. 11 2
		当 n = 3 时，有三种路径. 111 12 21
		当 n = 4 时，有五种路径. 1111 1112 1121 1211 2111

		是否可以推断出,当 n>2 时 f(n) = f(n-1) + f(n-2)

	解决方案:
	1. 递归
	2. 迭代
	3. 优化
*/

// 递归解法
func climbStairs_dg(n int) int {

	var dg func(n int) int

	dg = func(n int) int {
		//  递归终止条件
		if n == 1 {
			return 1
		}
		if n == 2 {
			return 2
		}

		// 递归
		return dg(n-1) + dg(n-2)
	}
	return dg(n)
}

// 递归优化，由于递归中存在重复计算，所以可以使用数组或者 map 来存储计算结果
func climbStairs_dg_yh(n int) int {

	m := make(map[int]int, 2*n)

	var dg func(n int) int
	dg = func(n int) int {
		//  递归终止条件
		if n == 1 {
			return 1
		}
		if n == 2 {
			return 2
		}

		// 如果已经计算过了，则直接返回
		if v, has := m[n]; has {
			return v
		}

		// 递归
		v := dg(n-1) + dg(n-2)
		m[n] = v
		return v
	}
	return dg(n)
}

// 遍历
func climbStairs_bl(n int) int {
	if n <= 2 {
		return n
	}

	table := make([]int, n+1)
	table[1] = 1
	table[2] = 2

	for i := 3; i <= n; i++ {
		table[i] = table[i-1] + table[i-2]
	}
	return table[n]
}

// 遍历 优化
func climbStairs_bl_yh(n int) int {
	if n <= 2 {
		return n
	}

	pre2, pre1 := 1, 2 // 分别存储前两个值
	for i := 3; i <= n; i++ {
		pre2, pre1 = pre1, pre1+pre2 // 计算之后保存，前进一步
	}
	return pre1
}
