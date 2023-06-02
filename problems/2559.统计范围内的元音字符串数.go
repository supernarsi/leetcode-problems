package problems

/*
 * @lc app=leetcode.cn id=2559 lang=golang
 *
 * [2559] 统计范围内的元音字符串数
 */

// @lc code=start
func vowelStrings(words []string, queries [][]int) (ans []int) {
	isEChar := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	preSum := make([]int, len(words)+1)
	for i, word := range words {
		if isEChar(word[0]) && isEChar(word[len(word)-1]) {
			preSum[i+1] = preSum[i] + 1
		} else {
			preSum[i+1] = preSum[i]
		}
	}

	for _, q := range queries {
		ans = append(ans, preSum[q[1]+1]-preSum[q[0]])
	}

	return
}

// @lc code=end
