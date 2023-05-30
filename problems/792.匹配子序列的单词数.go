package problems

import "sort"

/*
 * @lc app=leetcode.cn id=792 lang=golang
 *
 * [792] 匹配子序列的单词数
 */

// @lc code=start
func numMatchingSubseq(s string, words []string) int {
	ans := 0
	dicMap := make(map[byte][]int)
	for i, char := range s {
		dicMap[byte(char)-'a'] = append(dicMap[byte(char)-'a'], i)
	}
	for _, s2 := range words {
		if len(s2) > len(s) {
			continue
		}
		if isSubStr792(dicMap, s2) {
			ans++
		}
	}
	return ans
}

func isSubStr792(dicMap map[byte][]int, subStr string) bool {
	pivot := -1
	for _, c := range subStr {
		if posArr, ok := dicMap[byte(c)-'a']; !ok {
			return false
		} else {
			p := sort.SearchInts(posArr, pivot+1)
			if p >= len(posArr) {
				return false
			}
			pivot = posArr[p]
		}
	}
	return true
}

// @lc code=end
