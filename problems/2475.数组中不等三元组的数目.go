package problems

/*
 * @lc app=leetcode.cn id=2475 lang=golang
 *
 * [2475] 数组中不等三元组的数目
 */

// @lc code=start
func unequalTriplets(nums []int) int {
	var customCN = func(n, x int) int {
		res := 1
		for i := 0; i < x; i++ {
			res *= (n - i)
			res /= (i + 1)
		}
		return res
	}

	numsMap := make(map[int]int)
	for _, v := range nums {
		numsMap[v]++
	}
	total := len(nums)
	if len(numsMap) < 3 {
		return 0
	}

	ans := customCN(total, 3)
	for _, v := range numsMap {
		if v > 1 {
			ans -= (total - v) * customCN(v, 2)
			ans -= customCN(v, 3)
		}
	}
	return ans
}

// @lc code=end
