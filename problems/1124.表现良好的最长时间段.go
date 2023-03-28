package problems

/*
 * @lc app=leetcode.cn id=1124 lang=golang
 *
 * [1124] 表现良好的最长时间段
 */

// @lc code=start
func longestWPI(hours []int) int {
	preSum := calPreSum(hours)

	ans := 0
	for l, v := range preSum {
		for r := len(preSum) - 1; r >= l; r-- {
			if preSum[r]-v > 0 {
				ans = max(r-l, ans)
				break
			}
		}
	}
	return ans
}

func calPreSum(hours []int) []int {
	preSum := []int{0}
	for i, h := range hours {
		curVal := 1
		if h <= 8 {
			curVal = -1
		}
		preSum = append(preSum, preSum[i]+curVal)
	}
	return preSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
