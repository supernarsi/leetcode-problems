package problems

/*
 * @lc app=leetcode.cn id=2383 lang=golang
 *
 * [2383] 赢得比赛需要的最少训练时长
 */

// @lc code=start
func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) (ans int) {
	totEng := 0
	for _, e := range energy {
		totEng += e
	}
	if initialEnergy <= totEng {
		ans = totEng - initialEnergy + 1
	}
	// ans = 4

	for _, exp := range experience {
		tmpAdd := exp
		if exp >= initialExperience {
			tmpAdd += (exp - initialExperience + 1)
			ans += (exp - initialExperience + 1)
		}
		initialExperience += tmpAdd
	}

	return
}

// @lc code=end
