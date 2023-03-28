package problems

/*
 * @lc app=leetcode.cn id=1234 lang=golang
 *
 * [1234] 替换子串得到平衡字符串
 */

// @lc code=start
func balancedString(s string) int {
	cnt := make(map[byte]int)
	for _, c := range s {
		cnt[byte(c)]++
	}
	avg := len(s) / 4

	checkFunc := func() bool {
		if cnt['Q'] > avg || cnt['W'] > avg || cnt['E'] > avg || cnt['R'] > avg {
			return false
		}
		return true
	}

	if checkFunc() {
		return 0
	}

	ans, r := len(s), 0
	for l, c := range s {
		for r < len(s) && !checkFunc() {
			cnt[s[r]]--
			r++
		}
		if !checkFunc() {
			break
		}
		ans = min1234(ans, r-l)
		cnt[byte(c)]++
	}
	return ans
}

func min1234(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
