package problems

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=393 lang=golang
 *
 * [393] UTF-8 编码验证
 */

// @lc code=start
func validUtf8(data []int) bool {
	binNum := make([]string, len(data))
	for i, num := range data {
		binNum[i] = num2Str(num)
	}
	fmt.Println(binNum)
	skip := 0
	for _, binStr := range binNum {
		if skip == 0 {
			// 分隔位
			if binStr[0] == '0' {
				// 1 字节字符
				continue
			} else {
				// n 字节字符，统计 n 的数量
				n := 0
				for _, char := range binStr {
					if char == '1' {
						n++
					} else {
						break
					}
				}
				skip = n - 1
				if skip == 0 || skip > 4 {
					return false
				}
			}
		} else {
			// 后续位，判断接下来的 skip 个字节是否满足 10 开头
			if binStr[:2] != "10" {
				return false
			}
			skip--
		}
	}
	if skip > 0 {
		return false
	}
	return true
}

func num2Str(num int) string {
	nBin := strconv.FormatInt(int64(num), 2)
	diff := 8 - len(nBin)
	for diff > 0 {
		nBin = "0" + nBin
		diff--
	}
	return nBin
}

// @lc code=end
