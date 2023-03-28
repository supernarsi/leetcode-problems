package problems

/*
 * @lc app=leetcode.cn id=2319 lang=golang
 *
 * [2319] 判断矩阵是否是一个 X 矩阵
 */

// @lc code=start
func checkXMatrix(grid [][]int) bool {
	n := len(grid)
	limit := n / 2
	if n%2 == 1 {
		limit += 1
	}
	for i := 0; i < limit; i++ {
		rI := n - i - 1
		for j := 0; j < limit; j++ {
			rJ := n - j - 1
			if i == j {
				if grid[i][j] == 0 || grid[i][rJ] == 0 || grid[rI][j] == 0 || grid[rI][rJ] == 0 {
					return false
				}
			} else {
				if grid[i][j] != 0 || grid[i][rJ] != 0 || grid[rI][j] != 0 || grid[rI][rJ] != 0 {
					return false
				}
			}
		}
	}
	return true
}

// @lc code=end
