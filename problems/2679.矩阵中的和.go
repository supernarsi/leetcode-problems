package problems

import "sort"

/*
 * @lc app=leetcode.cn id=2679 lang=golang
 *
 * [2679] 矩阵中的和
 */

// @lc code=start
func matrixSum(nums [][]int) int {
	ans := 0
	colsNum := len(nums[0])
	rowsNum := len(nums)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for _, row := range nums {
		sort.Ints(row)
	}
	for i := colsNum - 1; i >= 0; i-- {
		maxVal := 0
		for j := 0; j < rowsNum; j++ {
			maxVal = max(maxVal, nums[j][i])
		}
		ans += maxVal
	}
	return ans
}

// @lc code=end
