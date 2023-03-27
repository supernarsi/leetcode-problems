package problems

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board {
			board[i][j] = '.'
		}
	}
	res := [][]string{}
	var backtack func(row int)
	backtack = func(row int) {
		if row == n {
			outputStr := formatBoard(board)
			res = append(res, outputStr)
			return
		}
		for col := 0; col < n; col++ {
			if !ok(board, row, col, n) {
				continue
			}
			board[row][col] = 'Q'
			backtack(row + 1)
			board[row][col] = '.'
		}
	}
	backtack(0)
	return res
}

func ok(board [][]byte, row, col, n int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func formatBoard(board [][]byte) []string {
	ans := []string{}
	for _, row := range board {
		str := ""
		for _, col := range row {
			str += string(col)
		}
		ans = append(ans, str)
	}
	return ans
}

// @lc code=end
