package problems

/*
 * @lc app=leetcode.cn id=1023 lang=golang
 *
 * [1023] 驼峰式匹配
 */

// @lc code=start
func camelMatch(queries []string, pattern string) (ans []bool) {
	for _, query := range queries {
		ans = append(ans, match(query, pattern))
	}

	return
}

func match(query, pattern string) bool {
	i, j := 0, 0
	for i < len(query) && j < len(pattern) {
		if query[i] == pattern[j] {
			i++
			j++
		} else if query[i] >= 'A' && query[i] <= 'Z' {
			return false
		} else {
			i++
		}
	}

	for i < len(query) {
		if query[i] >= 'A' && query[i] <= 'Z' {
			return false
		}
		i++
	}

	return j == len(pattern)
}

// @lc code=end
