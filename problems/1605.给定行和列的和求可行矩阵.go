package problems

/*
 * @lc app=leetcode.cn id=1605 lang=golang
 *
 * [1605] 给定行和列的和求可行矩阵
 */

// @lc code=start
func restoreMatrix(rowSum, colSum []int) [][]int {
	mat := make([][]int, len(rowSum))
	for i, rs := range rowSum {
		mat[i] = make([]int, len(colSum))
		for j, cs := range colSum {
			mat[i][j] = min1605(rs, cs)
			rs -= mat[i][j]
			colSum[j] -= mat[i][j]
		}
	}
	return mat
}

func min1605(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
