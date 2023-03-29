package problems

/*
 * @lc app=leetcode.cn id=1641 lang=golang
 *
 * [1641] 统计字典序元音字符串的数目
 */

// @lc code=start
func countVowelStrings(n int) int {
	return (n + 1) * (n + 2) * (n + 3) * (n + 4) / 24
}

// @lc code=end
