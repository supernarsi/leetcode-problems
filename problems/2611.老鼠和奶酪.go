package problems

import "sort"

/*
 * @lc app=leetcode.cn id=2611 lang=golang
 *
 * [2611] 老鼠和奶酪
 */

// @lc code=start
func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	diff := make([]int, len(reward1))
	score := 0
	for i := 0; i < len(reward1); i++ {
		diff[i] = reward1[i] - reward2[i]
		score += reward2[i]
	}
	sort.Ints(diff)
	for i := len(diff) - 1; i >= len(diff)-k; i-- {
		score += diff[i]
	}
	return score
}

// @lc code=end
