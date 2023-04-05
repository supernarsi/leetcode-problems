package problems

/*
 * @lc app=leetcode.cn id=2427 lang=golang
 *
 * [2427] 公因子的数目
 */

// @lc code=start
func commonFactors(a int, b int) int {
	if a > b {
		a, b = b, a
	}
	res := 0
	for i := 1; i <= a; i++ {
		if a%i == 0 && b%i == 0 {
			res++
		}
	}
	return res
}

// @lc code=end
