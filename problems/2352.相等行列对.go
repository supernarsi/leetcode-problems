package problems

import "fmt"

/*
 * @lc app=leetcode.cn id=2352 lang=golang
 *
 * [2352] 相等行列对
 */

// @lc code=start
func equalPairs(grid [][]int) int {
	ans := 0
	m := make(map[string]int)
	for i := range grid {
		key := ""
		for _, v := range grid[i] {
			key += ("-" + fmt.Sprint(v))
		}
		m[key]++
	}
	for i := range grid {
		key := ""
		for j := range grid[i] {
			key += ("-" + fmt.Sprint(grid[j][i]))
		}

		if v, ok := m[key]; ok {
			ans += v
		}
	}
	return ans
}

// @lc code=end
