package problems

/*
 * @lc app=leetcode.cn id=2357 lang=golang
 *
 * [2357] 使数组中所有元素都等于零
 */

// @lc code=start
func minimumOperations(nums []int) int {
	n := map[int]bool{}
	ans := 0
	for _, v := range nums {
		if v == 0 {
			continue
		}
		if _, ok := n[v]; !ok {
			n[v] = true
			ans++
		}
	}
	return ans
}

// @lc code=end
