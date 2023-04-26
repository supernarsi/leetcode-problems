package problems

/*
 * @lc app=leetcode.cn id=1031 lang=golang
 *
 * [1031] 两个非重叠子数组的最大和
 */

// @lc code=start
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	for i := 0; i <= n-firstLen-secondLen; i++ {
		for j := i + firstLen; j <= n-secondLen; j++ {
			res = max(res, sum[i+firstLen]-sum[i]+sum[j+secondLen]-sum[j])
		}
		for j := i + secondLen; j <= n-firstLen; j++ {
			res = max(res, sum[i+secondLen]-sum[i]+sum[j+firstLen]-sum[j])
		}
	}
	return res
}

// @lc code=end
