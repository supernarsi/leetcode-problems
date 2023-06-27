package problems

/*
 * @lc app=leetcode.cn id=1401 lang=golang
 *
 * [1401] 圆和矩形是否有重叠
 */

// @lc code=start
func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dist := 0
	if xCenter < x1 || xCenter > x2 {
		dist += min((x1-xCenter)*(x1-xCenter), (x2-xCenter)*(x2-xCenter))
	}
	if yCenter < y1 || yCenter > y2 {
		dist += min((y1-yCenter)*(y1-yCenter), (y2-yCenter)*(y2-yCenter))
	}
	return dist <= radius*radius
}

// @lc code=end
