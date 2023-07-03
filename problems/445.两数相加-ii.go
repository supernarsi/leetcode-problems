package problems

/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	toStack := func(ln *ListNode) (ans []int) {
		for ln != nil {
			ans = append([]int{ln.Val}, ans...)
			ln = ln.Next
		}
		return ans
	}
	s1, s2 := toStack(l1), toStack(l2)
	len1, len2 := len(s1), len(s2)
	totalLen := len1
	if len2 > totalLen {
		totalLen = len2
	}
	var resList []int
	add := false
	for i := 0; i < totalLen; i++ {
		var v1, v2 int
		if i+1 <= len1 {
			v1 = s1[i]
		} else {
			v1 = 0
		}
		if i+1 <= len2 {
			v2 = s2[i]
		} else {
			v2 = 0
		}
		tempSum := v1 + v2
		if add {
			tempSum += 1
		}
		if tempSum > 9 {
			tempSum = tempSum % 10
			add = true
		} else {
			add = false
		}
		resList = append(resList, tempSum)
	}
	if add {
		resList = append(resList, 1)
	}

	ans := &ListNode{}
	tmpAns := ans
	for i := len(resList) - 1; i >= 0; i-- {
		tmpAns.Val = resList[i]
		if i > 0 {
			tmpAns.Next = &ListNode{}
			tmpAns = tmpAns.Next
		}
	}

	return ans
}

// @lc code=end
