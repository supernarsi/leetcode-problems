package problems

import "sort"

/*
 * @lc app=leetcode.cn id=1626 lang=golang
 *
 * [1626] 无矛盾的最佳球队
 */

// @lc code=start
func bestTeamScore(scores []int, ages []int) int {
	n := len(scores)
	people := make([][]int, n)
	for i := range scores {
		people[i] = []int{scores[i], ages[i]}
	}
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] < people[j][0] {
			return true
		} else if people[i][0] > people[j][0] {
			return false
		}
		return people[i][1] < people[j][1]
	})
	dp := make([]int, n)
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if people[j][1] <= people[i][1] {
				dp[i] = max1626(dp[i], dp[j])
			}
		}
		dp[i] += people[i][0]
		res = max1626(res, dp[i])
	}
	return res
}

func max1626(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end
