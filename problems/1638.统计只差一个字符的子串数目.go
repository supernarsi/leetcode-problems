package problems

/*
 * @lc app=leetcode.cn id=1638 lang=golang
 *
 * [1638] 统计只差一个字符的子串数目
 */

// @lc code=start
func countSubstrings(s, t string) (ans int) {
	m, n := len(s), len(t)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diff := 0
			for k := 0; i+k < m && j+k < n; k++ {
				if s[i+k] != t[j+k] {
					diff++
				}
				if diff > 1 {
					break
				} else if diff == 1 {
					ans++
				}
			}
		}
	}
	return ans
}

// @lc code=end
