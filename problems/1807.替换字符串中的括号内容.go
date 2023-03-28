package problems

import "strings"

/*
 * @lc app=leetcode.cn id=1807 lang=golang
 *
 * [1807] 替换字符串中的括号内容
 */

// @lc code=start
func evaluate(s string, knowledge [][]string) string {
	dict := map[string]string{}
	for _, kd := range knowledge {
		dict[kd[0]] = kd[1]
	}
	ans := &strings.Builder{}
	start := -1
	for i, c := range s {
		if c == '(' {
			start = i
		} else if c == ')' {
			if t, ok := dict[s[start+1:i]]; ok {
				ans.WriteString(t)
			} else {
				ans.WriteByte('?')
			}
			start = -1
		} else if start < 0 {
			ans.WriteRune(c)
		}
	}
	return ans.String()
}

// @lc code=end
