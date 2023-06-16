package problems

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := nums[0]

	maxFunc := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	for i := 1; i < len(nums); i++ {
		dp[i] = maxFunc(dp[i-1]+nums[i], nums[i])
		ans = maxFunc(ans, dp[i])
	}
	return ans
}

// @lc code=end
