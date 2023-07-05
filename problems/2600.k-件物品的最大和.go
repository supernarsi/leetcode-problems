package problems

/*
 * @lc app=leetcode.cn id=2600 lang=golang
 *
 * [2600] K 件物品的最大和
 */

// @lc code=start
func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	sum, left := 0, k
	if numOnes > k {
		return k
	} else {
		sum += numOnes
		left = k - numOnes
	}
	if numZeros > left {
		return sum
	} else {
		left = left - numZeros
	}
	return sum - left
}

// @lc code=end
