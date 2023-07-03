package problems

/*
 * @lc app=leetcode.cn id=2490 lang=golang
 *
 * [2490] 回环句
 */

// @lc code=start
func isCircularSentence(sentence string) bool {
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}
	for i, v := range sentence {
		if v == ' ' && (sentence[i-1] != sentence[i+1]) {
			return false
		}
	}
	return true
}

// @lc code=end
