package problems

/*
 * @lc app=leetcode.cn id=2325 lang=golang
 *
 * [2325] 解密消息
 */

// @lc code=start
func decodeMessage(key string, message string) string {
	dic := map[rune]byte{}
	cur := byte('a')
	for _, word := range key {
		if _, ok := dic[word]; !ok && word != ' ' {
			dic[word] = cur
			cur++
		}
	}

	res := []byte(message)
	for i, w := range message {
		if w != ' ' {
			res[i] = dic[w]
		}
	}
	return string(res)
}

// @lc code=end
