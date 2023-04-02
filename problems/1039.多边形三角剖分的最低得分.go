package problems

/*
 * @lc app=leetcode.cn id=1039 lang=golang
 *
 * [1039] 多边形三角剖分的最低得分
 */

// @lc code=start
func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			dp[i][j] = 1<<31 - 1
			for k := i + 1; k < j; k++ {
				dp[i][j] = min1039(dp[i][j], dp[i][k]+dp[k][j]+values[i]*values[k]*values[j])
			}
		}
	}
	return dp[0][n-1]
}

func min1039(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
