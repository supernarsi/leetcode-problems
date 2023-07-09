package problems

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉 K 位数字
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	stack := []int{}
	tot := len(num)
	need := tot - k
	if need == 0 {
		return "0"
	}
	for i, v := range num {
		vInt, _ := strconv.Atoi(string(v))
		left := tot - 1 - i
		for len(stack) > 0 && len(stack)+left >= need && stack[len(stack)-1] > vInt {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, vInt)
	}
	stack = stack[:need]
	isZero := true
	ans := ""
	for _, v := range stack {
		if v == 0 && isZero {
			continue
		}
		ans += strconv.Itoa(v)
		isZero = false
	}
	if ans == "" {
		ans = "0"
	}
	return ans
}

// @lc code=end
