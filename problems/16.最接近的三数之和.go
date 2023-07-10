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

	abs := func(a int) int {
		if a < 0 {
			a = -a
		}
		return a
	}

	update := func(res int, minDiff int, sum int, target int) (resV int, newMinDiff int) {
		innerNowDiff := abs(sum - target)
		if innerNowDiff < minDiff {
			minDiff = innerNowDiff
			res = sum
		}
		return res, minDiff
	}

	calLeft := func(sortedArr []int, t int) int {
		l, r := 0, len(sortedArr)-1
		innerNowSum, res := 0, 0
		innerMinDiff := math.MaxInt64
		for l < r {
			innerNowSum = sortedArr[l] + sortedArr[r]
			if innerNowSum == t {
				return t
			}

			res, innerMinDiff = update(res, innerMinDiff, innerNowSum, t)
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
		tmpSun := leftSum + v
		ans, minDiff = update(ans, minDiff, tmpSun, target)
	}

	return ans
}

// @lc code=end
