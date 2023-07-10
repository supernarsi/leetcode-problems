package problems

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	// -4, -1, 1, 2

	calLeft := func(sortedArr []int, t int) int {
		l, r := 0, len(sortedArr)-1
		innerNowSum, res := 0, 0
		innerMinDiff := math.MaxInt64
		for l < r {
			innerNowSum = sortedArr[l] + sortedArr[r]
			if innerNowSum == t {
				return t
			}
			innerNowDiff := innerNowSum - t
			if innerNowDiff < 0 {
				innerNowDiff = -innerNowDiff
			}
			if innerNowDiff < innerMinDiff {
				innerMinDiff = innerNowDiff
				res = innerNowSum
			}

			if innerNowSum < t {
				l++
			} else if innerNowSum > t {
				r--
			}
		}
		return res
	}

	ans := 0
	minDiff := math.MaxInt64
	for i := 0; i <= len(nums)-3; i++ {
		v := nums[i]
		leftSum := calLeft(nums[i+1:], target-v)
		if leftSum+v == target {
			return target
		}
		nowDiff := leftSum + v - target
		if nowDiff < 0 {
			nowDiff = -nowDiff
		}
		if nowDiff < minDiff {
			minDiff = nowDiff
			ans = leftSum + v
		}
	}

	return ans
}

// @lc code=end
