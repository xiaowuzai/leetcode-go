package findcontentchildren

import "sort"

// 贪心算法，找当前最优解
func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	i, j := 0, 0
	for i < len(g) && j < len(s) {
		// 满足胃口
		if g[i] <= s[j] {
			i++
		}
		//不满足胃口，找大饼干,满足的时候也要移动
		j++
	}
	return i
}
