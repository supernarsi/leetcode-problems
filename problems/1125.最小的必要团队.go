package problems

/*
 * @lc app=leetcode.cn id=1125 lang=golang
 *
 * [1125] 最小的必要团队
 */

// @lc code=start
func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	n, m := len(req_skills), len(people)
	skill_index := make(map[string]int)
	for i, skill := range req_skills {
		skill_index[skill] = i
	}
	dp := make([][]int, 1<<n)
	dp[0] = []int{}
	for i := 0; i < m; i++ {
		cur_skill := 0
		for _, s := range people[i] {
			cur_skill |= 1 << skill_index[s]
		}
		for prev := 0; prev < len(dp); prev++ {
			if dp[prev] == nil {
				continue
			}
			comb := prev | cur_skill
			if dp[comb] == nil || len(dp[prev])+1 < len(dp[comb]) {
				dp[comb] = make([]int, len(dp[prev]))
				copy(dp[comb], dp[prev])
				dp[comb] = append(dp[comb], i)
			}
		}
	}
	return dp[(1<<n)-1]
}

// @lc code=end
