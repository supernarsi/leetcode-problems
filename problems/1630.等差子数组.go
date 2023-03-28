package problems

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1630 lang=golang
 *
 * [1630] 等差子数组
 */

// @lc code=start
func checkArithmeticSubarrays(nums []int, l []int, r []int) (ans []bool) {
	length := len(l)
	ans = make([]bool, length)
	for i := 0; i < length; i++ {
		s, e := l[i], r[i]

		subNum := make([]int, e-s+1)
		for i, n := range nums[s : e+1] {
			subNum[i] = n
		}
		sort.Ints(subNum)
		ans[i] = numIsEquDiff(subNum)
	}
	return
}

func numIsEquDiff(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	diff := nums[1] - nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != diff {
			return false
		}
	}
	return true
}

// @lc code=end
