package problems

import "strings"

/*
 * @lc app=leetcode.cn id=1813 lang=golang
 *
 * [1813] 句子相似性 III
 */

// @lc code=start
func areSentencesSimilar(sentence1, sentence2 string) bool {
	words1 := strings.Split(sentence1, " ")
	words2 := strings.Split(sentence2, " ")
	i, n := 0, len(words1)
	j, m := 0, len(words2)
	for i < n && i < m && words1[i] == words2[i] {
		i++
	}
	for j < n-i && j < m-i && words1[n-j-1] == words2[m-j-1] {
		j++
	}
	return i+j == min1813(n, m)
}

func min1813(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
