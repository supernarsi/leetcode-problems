package problems

/*
 * @lc app=leetcode.cn id=1606 lang=golang
 *
 * [1606] 找到处理最多请求的服务器
 */

// @lc code=start
func busiestServers(k int, arrival []int, load []int) []int {
	serverStatus := map[int]int{}
	serverTimes := map[int]int{}
	for i, time := range arrival {
		serverNum := i % k
		for n := 0; n < k; n++ {
			trueServerNum := (serverNum + n) % k
			// 如果该次请求到达时，次服务器已经空闲
			if serverStatus[trueServerNum] <= time {
				// 更新该服务器状态
				serverStatus[trueServerNum] = time + load[i]
				// 服务器处理次数 + 1
				serverTimes[trueServerNum]++
				break
			}
		}
	}

	ans := []int{}
	dealTimes := 0
	for i, times := range serverTimes {
		if times > dealTimes {
			dealTimes = times
			ans = []int{i}
		} else if times == dealTimes {
			ans = append(ans, i)
		}
	}
	return ans
}

// @lc code=end
