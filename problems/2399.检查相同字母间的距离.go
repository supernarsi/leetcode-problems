package problems

/*
 * @lc app=leetcode.cn id=2399 lang=golang
 *
 * [2399] 检查相同字母间的距离
 */

// @lc code=start
func checkDistances(s string, distance []int) bool {
	sMap := make(map[byte]int)
	lS := len(s)
	if lS%2 != 0 {
		return false
	}
	for i := 0; i < lS; i++ {
		if v, ok := sMap[s[i]]; ok {
			if i-v != distance[s[i]-'a']+1 {
				return false
			}
		} else {
			sMap[s[i]] = i
		}
	}
	return true
}

// @lc code=end
