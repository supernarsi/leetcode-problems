package problems

/*
 * @lc app=leetcode.cn id=1026 lang=golang
 *
 * [1026] 节点与其祖先之间的最大差值
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max1026(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min1026(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dfs1026(root *TreeNode, mi, ma int) int {
	if root == nil {
		return 0
	}
	diff := max1026(abs(root.Val-mi), abs(root.Val-ma))
	mi, ma = min1026(mi, root.Val), max1026(ma, root.Val)
	diff = max1026(diff, dfs1026(root.Left, mi, ma))
	diff = max1026(diff, dfs1026(root.Right, mi, ma))
	return diff
}

func maxAncestorDiff(root *TreeNode) int {
	return dfs1026(root, root.Val, root.Val)
}

// @lc code=end
