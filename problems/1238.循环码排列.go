package problems

/*
 * @lc app=leetcode.cn id=1238 lang=golang
 *
 * [1238] 循环码排列
 */

// @lc code=start
func circularPermutation(n int, start int) []int {
	ans := make([]int, 1, 1<<n)
	ans[0] = start
	for i := 1; i <= n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, ((ans[j]^start)|(1<<(i-1)))^start)
		}
	}
	return ans
}

// @lc code=end
