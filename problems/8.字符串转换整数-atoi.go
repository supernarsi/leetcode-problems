package problems

import "math"

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	sum, sign, i, n := 0, 1, 0, len(s)

	// 丢弃无用的前导空格
	for i < n && s[i] == ' ' {
		i++
	}
	// 判定正负
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	// 从左到右依次累加
	for i < n && s[i] >= '0' && s[i] <= '9' {
		sum = 10*sum + int(s[i]-'0')
		// 整数超过32位有符号整数范围,特殊处理
		if sign*sum < math.MinInt32 {
			return math.MinInt32
		} else if sign*sum > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * sum
}

// cv

// @lc code=end
