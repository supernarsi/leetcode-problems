package problems

import "math"

/*
 * @lc app=leetcode.cn id=1140 lang=golang
 *
 * [1140] 石子游戏 II
 */

// @lc code=start
func stoneGameII(piles []int) int {
	prefixSum := make([]int, len(piles)+1)
	for i, v := range piles {
		prefixSum[i+1] = prefixSum[i] + v
	}
	type pair struct{ i, m int }
	dp := map[pair]int{}
	var f func(i int, m int) int
	f = func(i int, m int) int {
		if i == len(piles) {
			return 0
		}
		if v, ok := dp[pair{i, m}]; ok {
			return v
		}
		maxVal := math.MinInt
		for x := 1; x <= 2*m; x++ {
			if i+x > len(piles) {
				break
			}
			maxVal = max1140(maxVal, prefixSum[i+x]-prefixSum[i]-f(i+x, max1140(m, x)))
		}
		dp[pair{i, m}] = maxVal
		return maxVal
	}
	return (prefixSum[len(piles)] + f(0, 1)) / 2
}

func max1140(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end
