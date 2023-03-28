package problems

/*
 * @lc app=leetcode.cn id=1139 lang=golang
 *
 * [1139] 最大的以 1 为边界的正方形
 */

// @lc code=start

func largest1BorderedSquare(grid [][]int) int {
	lenW, lenH := len(grid), len(grid[0])
	maxLen := min(lenW, lenH)
	for curLen := maxLen; curLen > 0; curLen-- {
		// fmt.Println("边长：", curLen)
		// 起点坐标
		for sC := 0; sC < lenH-curLen+1; sC++ {
			for sR := 0; sR < lenW-curLen+1; sR++ {
				// fmt.Println("起点坐标：", sC, sR)
				if isCalBorad(grid, sR, sC, curLen) {
					return curLen * curLen
				}
			}
		}
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isCalBorad(grid [][]int, sR, sC, len int) bool {
	for r := sR; r < sR+len; r++ {
		for c := sC; c < sC+len; c++ {
			if c == sC || r == sR || c == sC+len-1 || r == sR+len-1 {
				if grid[r][c] != 1 {
					return false
				}
			}
		}
	}
	return true
}

// @lc code=end
