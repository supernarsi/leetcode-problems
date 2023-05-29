package problems

/*
 * @lc app=leetcode.cn id=2455 lang=golang
 *
 * [2455] 可被三整除的偶数的平均值
 */

// @lc code=start
func averageValue(nums []int) int {
	n, sum := 0, 0
	for _, v := range nums {
		if v%6 == 0 {
			sum += v
			n++
		}
	}
	if sum == 0 {
		return 0
	}
	return sum / n
}

// @lc code=end
