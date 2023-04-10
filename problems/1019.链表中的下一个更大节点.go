package problems

/*
 * @lc app=leetcode.cn id=1019 lang=golang
 *
 * [1019] 链表中的下一个更大节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) (ans []int) {
	for node := head; node != nil; node = node.Next {
		ans = append(ans, node.Val)
	}
	stack := []int{}
	for i, v := range ans {
		for len(stack) > 0 && ans[stack[len(stack)-1]] < v {
			ans[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for _, v := range stack {
		ans[v] = 0
	}
	return
}

// @lc code=end
