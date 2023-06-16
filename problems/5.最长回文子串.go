package problems

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	ans := ""
	for l := 1; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			if l == 1 {
				dp[i][j] = true
			} else if l == 2 {
				dp[i][j] = (s[i] == s[j])
			} else {
				dp[i][j] = (dp[i+1][j-1] && (s[i] == s[j]))
			}

			if dp[i][j] && l > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}

// @lc code=end
