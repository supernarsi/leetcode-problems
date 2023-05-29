package problems

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=475 lang=golang
 *
 * [475] 供暖器
 */

// @lc code=start
func findRadius(houses []int, heaters []int) int {
	// sort.Ints(houses)
	sort.Ints(heaters)

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	radius := 0
	for _, house := range houses {
		minDis := math.MaxInt
		p := sort.SearchInts(heaters, house)
		if p == 0 {
			// 在最左侧
			minDis = heaters[0] - house
		} else if p == len(heaters) {
			// 在最右侧
			minDis = house - heaters[len(heaters)-1]
		} else {
			if heaters[p] == house {
				// 刚好跟 heater 重合
				minDis = 0
			} else {
				// 左右两侧都有 heater
				minDis = min(house-heaters[p-1], heaters[p]-house)
			}
		}

		radius = max(radius, minDis)
	}
	return radius
}

// @lc code=end
