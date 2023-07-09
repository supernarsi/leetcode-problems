package problems

import (
	"fmt"
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

	exist := map[string]bool{}
	for i, n := range nums {
		res := calLeft(nums[i+1:], -n)
		if len(res) > 0 {
			for key, val := range res {
				tmp := []int{n, key, val}
				sort.Ints(tmp)
				tmpKey := ""
				for _, tInt := range tmp {
					tmpKey += ("," + fmt.Sprintf("%d", tInt))
				}
				if !exist[tmpKey] {
					ans = append(ans, tmp)
					exist[tmpKey] = true
				}
			}
		}
	}
	return
}

// @lc code=end
