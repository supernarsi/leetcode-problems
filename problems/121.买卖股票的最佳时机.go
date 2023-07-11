package problems

import "math"

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit1(prices []int) int {
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	minP, maxV := math.MaxInt, 0
	for i := range prices {
		minP = min(prices[i], minP)
		maxV = max(maxV, prices[i]-minP)
	}
	return maxV
}

// @lc code=end
