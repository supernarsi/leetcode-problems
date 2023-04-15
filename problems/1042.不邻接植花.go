package problems

/*
 * @lc app=leetcode.cn id=1042 lang=golang
 *
 * [1042] 不邻接植花
 */

// @lc code=start
func gardenNoAdj(n int, paths [][]int) (ans []int) {
	// 1. 构建邻接表
	adj := make([][]int, n)
	for _, path := range paths {
		adj[path[0]-1] = append(adj[path[0]-1], path[1]-1)
		adj[path[1]-1] = append(adj[path[1]-1], path[0]-1)
	}

	// 2. 染色
	ans = make([]int, n)
	for i := 0; i < n; i++ {
		colors := [5]bool{}
		for _, j := range adj[i] {
			colors[ans[j]] = true
		}
		for c := 1; c <= 4; c++ {
			if !colors[c] {
				ans[i] = c
				break
			}
		}
	}

	return
}

// @lc code=end
