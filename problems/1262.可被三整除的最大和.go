package problems

import "math"

/*
 * @lc app=leetcode.cn id=1262 lang=golang
 *
 * [1262] 可被三整除的最大和
 */

// @lc code=start
func maxSumDivThree(nums []int) int {
	n := len(nums)
	f := make([][3]int, n+1)
	f[0][1] = math.MinInt
	f[0][2] = math.MinInt

	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	for i, x := range nums {
		for j := 0; j < 3; j++ {
			f[i+1][j] = max(f[i][j], f[i][(j+x)%3]+x)
		}
	}
	return f[n][0]
}

// @lc code=end
