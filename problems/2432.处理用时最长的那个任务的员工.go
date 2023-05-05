package problems

/*
 * @lc app=leetcode.cn id=2432 lang=golang
 *
 * [2432] 处理用时最长的那个任务的员工
 */

// @lc code=start
func hardestWorker(n int, logs [][]int) int {
	id := logs[0][0]
	sT := logs[0][1]
	maxT := sT
	for _, log := range logs {
		if log[1]-sT >= maxT {
			if log[1]-sT > maxT {
				id = log[0]
			} else if log[0] < id {
				id = log[0]
			}

			maxT = log[1] - sT
		}
		sT = log[1]
	}
	return id
}

// @lc code=end
