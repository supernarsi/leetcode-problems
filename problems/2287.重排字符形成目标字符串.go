package problems

/*
 * @lc app=leetcode.cn id=2287 lang=golang
 *
 * [2287] 重排字符形成目标字符串
 */

// @lc code=start
func rearrangeCharacters(s string, target string) int {
	dic, sMap := make(map[rune]int), make(map[rune]int)
	for _, word := range target {
		if val, ok := dic[word]; ok {
			dic[word] = val + 1
		} else {
			dic[word] = 1
			sMap[word] = 0
		}
	}

	for _, w := range s {
		if _, ok := dic[w]; ok {
			if v, ok := sMap[w]; ok {
				sMap[w] = v + 1
			} else {
				sMap[w] = 1
			}
		}
	}

	res := -1
	for word, sNum := range sMap {
		nums := dic[word]
		times := sNum / nums
		if times == 0 {
			return 0
		}
		if res == -1 || times < res {
			res = times
		}
	}
	return res
}

// @lc code=end
