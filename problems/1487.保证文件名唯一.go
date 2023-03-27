package problems

import "strconv"

/*
 * @lc app=leetcode.cn id=1487 lang=golang
 *
 * [1487] 保证文件名唯一
 */

// @lc code=start
func getFolderNames(names []string) (ans []string) {
	dic := map[string]int{}
	for _, name := range names {
		minN := dic[name]
		if minN == 0 {
			ans = append(ans, name)
			dic[name]++
		} else {
			var tmpName = ""
			for tmpName = name + "(" + strconv.Itoa(minN+1) + ")"; dic[tmpName] != 0; {
				minN += 1
			}
			dic[tmpName] = 1
			dic[name] = minN
			ans = append(ans, tmpName)
		}
	}
	return ans
}

// @lc code=end
