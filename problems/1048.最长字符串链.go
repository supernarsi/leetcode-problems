package problems

import "sort"

/*
 * @lc app=leetcode.cn id=1048 lang=golang
 *
 * [1048] 最长字符串链
 */

// @lc code=start
func longestStrChain(words []string) (ans int) {
	// dp[i] 表示以 words[i] 为结尾的最长字符串链的长度
	dp := make([]int, len(words))
	// 按照字符串长度升序排序
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := range words {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if len(words[i])-len(words[j]) > 1 {
				break
			}
			if len(words[i])-len(words[j]) == 1 {
				if isPredecessor(words[j], words[i]) {
					dp[i] = max(dp[i], dp[j]+1)
				}
			}
		}
		ans = max(ans, dp[i])
	}

	return
}

func isPredecessor(a, b string) bool {
	if len(a) != len(b)-1 {
		return false
	}

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
		}
		j++
	}

	return i == len(a)
}

// @lc code=end
