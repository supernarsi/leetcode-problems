package problems

/*
 * @lc app=leetcode.cn id=1110 lang=golang
 *
 * [1110] 删点成林
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

func delNodes(root *TreeNode, to_delete []int) (ans []*TreeNode) {
	deleteMap := make(map[int]bool)
	for _, v := range to_delete {
		deleteMap[v] = true
	}

	trees := []*TreeNode{}
	dfs1110(root, true, deleteMap, &trees)
	return trees
}

func dfs1110(root *TreeNode, isRoot bool, deleteMap map[int]bool, trees *[]*TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	isDel := deleteMap[root.Val]
	root.Left = dfs1110(root.Left, isDel, deleteMap, trees)
	root.Right = dfs1110(root.Right, isDel, deleteMap, trees)
	if isDel {
		return nil
	} else {
		if isRoot {
			*trees = append(*trees, root)
		}
		return root
	}
}

// @lc code=end
