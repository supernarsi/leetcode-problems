package problems

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	end, maxP, step := 0, 0, 0

	for i := 0; i < len(nums)-1; i++ {
		maxP = max45(i+nums[i], maxP)
		if i == end {
			end = maxP
			step++
		}
	}
	return step
}

func max45(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
