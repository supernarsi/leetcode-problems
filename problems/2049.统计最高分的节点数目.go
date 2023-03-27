package problems

//func main() {
//	parents := []int{-1, 2, 0, 2, 0}
//	ans := countHighestScoreNodes(parents)
//
//	fmt.Println(ans)
//}

/*
 * @lc app=leetcode.cn id=2049 lang=golang
 *
 * [2049] 统计最高分的节点数目
 */

// @lc code=start
func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	root := make([][]int, n)
	for i := 1; i < n; i++ {
		pRoot := parents[i]
		root[pRoot] = append(root[pRoot], i)
	}
	calRoot := make([]int, n)
	calRoot = dfs(root, 0, calRoot)
	maxV, maxNum := 0, 0
	// fmt.Println("root is: ", root)
	// fmt.Println("calRoot is: ", calRoot)
	for i := 0; i < n; i++ {
		valList := root[i]
		sum := 1
		for _, cVal := range valList {
			sum *= calRoot[cVal]
		}
		if i != 0 {
			sum *= (n - calRoot[i])
		}
		if sum > maxV {
			maxV = sum
			maxNum = 1
		} else if sum == maxV {
			maxNum++
		}
	}
	return maxNum
}

func dfs(tree [][]int, idx int, ans []int) []int {
	if len(tree[idx]) == 0 {
		ans[idx] = 1
		return ans
	}
	size := 1
	for _, child := range tree[idx] {
		ans = dfs(tree, child, ans)
		size += ans[child]
	}
	ans[idx] = size
	return ans
}

// @lc code=end
