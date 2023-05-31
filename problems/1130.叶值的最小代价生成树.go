package problems

/*
 * @lc app=leetcode.cn id=1130 lang=golang
 *
 * [1130] 叶值的最小代价生成树
 */

// @lc code=start
func mctFromLeafValues(arr []int) int {
	n := len(arr)
	dp, mval := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		mval[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		mval[j][j] = arr[j]
		for i := j - 1; i >= 0; i-- {
			mval[i][j] = max1130(arr[i], mval[i+1][j])
			dp[i][j] = 0x3f3f3f3f
			for k := i; k < j; k++ {
				dp[i][j] = min1130(dp[i][j], dp[i][k]+dp[k+1][j]+mval[i][k]*mval[k+1][j])
			}
		}
	}
	return dp[0][n-1]
}

func min1130(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max1130(a ...int) int {
	maxVal := a[0]
	for _, num := range a {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}

// @lc code=end
