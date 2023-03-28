package problems

import "sort"

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) (res [][]string) {
	posMap := make(map[string]int)
	for _, word := range strs {
		sortWord := sortStr(word)
		if pos, ok := posMap[sortWord]; ok {
			res[pos] = append(res[pos], word)
		} else {
			posMap[sortWord] = len(res)
			res = append(res, []string{word})
		}
	}
	return res
}

func sortStr(word string) string {
	b := []byte(word)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}

// @lc code=end
