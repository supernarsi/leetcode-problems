package problems

/*
 * @lc app=leetcode.cn id=2460 lang=golang
 *
 * [2460] 对数组执行操作
 */

// @lc code=start
func applyOperations(nums []int) (ans []int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = 2 * nums[i]
			nums[i+1] = 0
		}
	}

	zores := 0
	for _, v := range nums {
		if v != 0 {
			ans = append(ans, v)
		} else {
			zores++
		}
	}
	for i := 0; i < zores; i++ {
		ans = append(ans, 0)
	}

	return
}

// @lc code=end
