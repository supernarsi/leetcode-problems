package problems

/*
 * @lc app=leetcode.cn id=275 lang=golang
 *
 * [275] H 指数 II
 */

// @lc code=start
func hIndex(citations []int) int {
	n := len(citations)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if citations[mid] < n-mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return n - left
}

// @lc code=end
