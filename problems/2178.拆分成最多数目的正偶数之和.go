package problems

/*
 * @lc app=leetcode.cn id=2178 lang=golang
 *
 * [2178] 拆分成最多数目的正偶数之和
 */

// @lc code=start
func maximumEvenSplit(finalSum int64) (ans []int64) {
	if finalSum%2 == 1 || finalSum == 0 {
		return ans
	}
	var n int64 = 1
	var sum int64 = 0
	for sum < finalSum {
		if sum+2*n > finalSum {
			oldV := ans[len(ans)-1]
			ans[len(ans)-1] = oldV + finalSum - sum
			return ans
		}
		ans = append(ans, 2*n)
		sum += 2 * n
		n++
	}

	return
}

// @lc code=end
