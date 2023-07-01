package problems

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) (ans []int) {
	need := make(map[int]int)
	for i, v := range nums {
		if res, ok := need[v]; ok {
			return []int{res, i}
		}
		need[target-v] = i
	}
	return
}

// @lc code=end
