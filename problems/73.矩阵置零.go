package problems

/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	c, r := false, false
	for _, v := range matrix {
		if v[0] == 0 {
			c = true
			break
		}
	}
	for _, v := range matrix[0] {
		if v == 0 {
			r = true
			break
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if c {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
	if r {
		for j := range matrix[0] {
			matrix[0][j] = 0
		}
	}
}

// @lc code=end
