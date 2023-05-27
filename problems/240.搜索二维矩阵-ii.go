package problems

/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	pX, pY := 0, len(matrix[0])-1

	for pX < len(matrix) && pY >= 0 {
		if matrix[pX][pY] == target {
			return true
		}
		if matrix[pX][pY] < target {
			pX++
		} else {
			pY--
		}
	}
	return false
}

// @lc code=end
