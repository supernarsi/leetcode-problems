package problems

/*
 * @lc app=leetcode.cn id=1653 lang=golang
 *
 * [1653] 使字符串平衡的最少删除次数
 */

// @lc code=start
// baabbaabbaabbaaaa
func minimumDeletions(s string) int {
	bTotal := 0
	dp := make([]int, len(s))
	for i, char := range s {
		if i == 0 {
			if char == 'b' {
				bTotal++
			}
			continue
		}
		if char == 'a' {
			dp[i] = min1653(dp[i-1]+1, bTotal)
		} else {
			dp[i] = dp[i-1]
			bTotal++
		}
	}
	return dp[len(s)-1]
}

func min1653(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
