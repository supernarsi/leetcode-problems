package problems

/*
 * @lc app=leetcode.cn id=1615 lang=golang
 *
 * [1615] 最大网络秩
 */

// @lc code=start
func maximalNetworkRank(n int, roads [][]int) int {
	// 初始化城市的直接连接道路数数组
	ranks := make([]int, n)

	// 用 map 统计每对城市之间的道路数
	roadCount := make(map[int]map[int]int)
	for _, road := range roads {
		i, j := road[0], road[1]
		if _, ok := roadCount[i]; !ok {
			roadCount[i] = make(map[int]int)
		}
		if _, ok := roadCount[j]; !ok {
			roadCount[j] = make(map[int]int)
		}
		roadCount[i][j]++
		roadCount[j][i]++
		ranks[i]++
		ranks[j]++
	}

	// 计算所有城市对的网络秩，找到最大值
	maxRank := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			rank := ranks[i] + ranks[j]
			if cnt, ok := roadCount[i][j]; ok {
				rank -= cnt
			}
			if rank > maxRank {
				maxRank = rank
			}
		}
	}

	return maxRank
}

// @lc code=end
