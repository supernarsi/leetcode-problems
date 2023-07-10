package problems

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) (ans [][]int) {
	calLeft := func(arr []int, sum int) map[int]int {
		if len(arr) < 2 {
			return nil
		}
		res := map[int]int{}
		tmpMap := map[int]int{}
		for _, v := range arr {
			if tmpRes, ok := tmpMap[v]; ok {
				res[v] = tmpRes
				continue
			}
			tmpMap[sum-v] = v
		}
		return res
	}

	exist := map[int]bool{}
	sort.Ints(nums)
	for i, n := range nums {
		if exist[n] {
			continue
		}
		res := calLeft(nums[i+1:], -n)
		if len(res) > 0 {
			for key, val := range res {
				tmp := []int{n, key, val}
				ans = append(ans, tmp)
				exist[n] = true
			}
		}
	}
	return
}

// @lc code=end
