package problems

/*
 * @lc app=leetcode.cn id=1911 lang=golang
 *
 * [1911] 最大子序列交替和
 */

// @lc code=start
func maxAlternatingSum(nums []int) int64 {
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = nums[0]
	dp[0][1] = 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(nums); i++ {
		dp[i][0] = max(dp[i-1][1]+nums[i], dp[i-1][0])
		dp[i][1] = max(dp[i-1][0]-nums[i], dp[i-1][1])
	}
	return int64(max(dp[len(nums)-1][0], dp[len(nums)-1][1]))
}

// @lc code=end
