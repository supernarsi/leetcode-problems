package problems

/*
 * @lc app=leetcode.cn id=1041 lang=golang
 *
 * [1041] 困于环中的机器人
 */

// @lc code=start
func isRobotBounded(instructions string) bool {
	// 0: north, 1: east, 2: south, 3: west
	direction := 0
	x, y := 0, 0
	for _, c := range instructions {
		switch c {
		case 'G':
			switch direction {
			case 0:
				y++
			case 1:
				x++
			case 2:
				y--
			case 3:
				x--
			}
		case 'L':
			direction = (direction + 3) % 4
		case 'R':
			direction = (direction + 1) % 4
		}
	}
	if direction != 0 || (x == 0 && y == 0) {
		return true
	}
	return false
}

// @lc code=end
