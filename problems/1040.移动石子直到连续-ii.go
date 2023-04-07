package problems

import "sort"

/*
 * @lc app=leetcode.cn id=1040 lang=golang
 *
 * [1040] 移动石子直到连续 II
 */

// @lc code=start
func numMovesStonesII(stones []int) []int {
	n := len(stones)
	sort.Ints(stones)
	max := max1040(stones[n-1]-stones[1]-n+2, stones[n-2]-stones[0]-n+2)
	min := n
	for i, j := 0, 0; i < n; i++ {
		for j < n && stones[j]-stones[i] < n {
			j++
		}
		if j-i == n-1 && stones[j-1]-stones[i]+1 == n-1 {
			min = min1040(min, 2)
		} else {
			min = min1040(min, n-j+i)
		}
	}
	return []int{min, max}

}

func min1040(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max1040(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
