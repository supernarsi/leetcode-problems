package problems

import "sort"

/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	intervals = append(intervals, newInterval)
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
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
