package problems

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=1233 lang=golang
 *
 * [1233] 删除子文件夹
 */

// @lc code=start
func removeSubfolders(folder []string) []string {
	sort.Strings(folder)

	res := []string{folder[0]}

	for _, f := range folder[1:] {
		last := res[len(res)-1]
		if !strings.HasPrefix(f, last) || f[len(last)] != '/' {
			res = append(res, f)
		}
	}
	return res
}

// @lc code=end
