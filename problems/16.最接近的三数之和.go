package problems

import "math"

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	t := len(nums)
	ans := 0
	initDiff := math.MaxInt64

	for i := 0; i < t; i++ {
		for j := i + 1; j < t; j++ {
			for k := j + 1; k < t; k++ {
				sum := nums[i] + nums[j] + nums[k]
				diff := sum - target
				if diff < 0 {
					diff = -diff
				}
				if diff < initDiff {
					initDiff = diff
					ans = sum
				}
			}
		}
	}
	return ans
}

// @lc code=end
