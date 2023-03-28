package problems

/*
 * @lc app=leetcode.cn id=1814 lang=golang
 *
 * [1814] 统计一个数组中好对子的数目
 */

// @lc code=start
func countNicePairs(nums []int) int {
	res := 0
	if len(nums) == 1 {
		return res
	}
	revNumMap := make(map[int]int)
	ans := 0
	for _, iNum := range nums {
		rINum := revNum(iNum)
		ans += revNumMap[iNum-rINum]
		revNumMap[iNum-rINum]++
	}
	return ans % (1e9 + 7)
}

func revNum(num int) int {
	if num < 10 {
		return num
	}
	res := 0
	for num > 0 {
		res = res*10 + num%10
		num /= 10
	}
	return res
}

// @lc code=end
