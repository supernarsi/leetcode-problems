package problems

import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	ans := [][]int{}
	tempArr := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > tempArr[1] {
			ans = append(ans, tempArr)
			tempArr = intervals[i]
		} else {
			if intervals[i][1] > tempArr[1] {
				tempArr[1] = intervals[i][1]
			}
		}

		if i == len(intervals)-1 {
			ans = append(ans, tempArr)
		}
	}
	return ans
}

// @lc code=end
