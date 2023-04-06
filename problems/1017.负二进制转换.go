package problems

/*
 * @lc app=leetcode.cn id=1017 lang=golang
 *
 * [1017] 负二进制转换
 */

// @lc code=start
func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	ans := ""
	for n != 0 {
		ans = string(rune('0'+n&1)) + ans
		n = -(n >> 1)
	}
	return ans
}

// @lc code=end
